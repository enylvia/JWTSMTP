package service

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/mail.v2"
	"jwtsmtp/entity"
	"jwtsmtp/helper"
	"jwtsmtp/repository"
	"time"
)

type Service interface {
	Register(input entity.RegisterInput) (entity.User, error)
	Login(input entity.LoginInput) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	SendMail(input entity.MailInput) (entity.LogMail, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Register(input entity.RegisterInput) (entity.User, error) {
	user := entity.User{}

	checkuser, err := s.repository.FindByUsername(input.Username)
	if checkuser.Username != "" {
		return checkuser, fmt.Errorf("username %s is already taken", input.Username)
	}
	user.Username = input.Username
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	helper.ErrorIfNotNil(err)
	user.Password = string(hashPassword)

	newUser, err := s.repository.Create(user)
	helper.ErrorIfNotNil(err)

	return newUser, nil
}

func (s *service) Login(input entity.LoginInput) (entity.User, error) {
	user := entity.User{}

	user, err := s.repository.FindByUsername(input.Username)
	if user.Username == "" {
		return user, fmt.Errorf("username or password is incorrect")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, fmt.Errorf("username or password is incorrect")
	}

	return user, nil
}

func (s *service) GetUserByUsername(username string) (entity.User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *service) SendMail(input entity.MailInput) (entity.LogMail, error) {
	logmails := entity.LogMail{}

	mail := gomail.NewMessage()
	mail.SetHeader("From", "")
	mail.SetHeader("To", input.To)
	mail.SetHeader("Subject", "JWT SMTP TEST")
	mail.SetBody("text/plain", input.Message)

	delivery := gomail.NewDialer("smtp.office365.com", 587, "", "")
	delivery.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := delivery.DialAndSend(mail); err != nil {
		logmails.Status = "failed"
		logmails.CreatedAt = time.Now()
		return logmails, fmt.Errorf("failed to send mail %s", err)
	}
	logmails.Status = "success"
	logmails.CreatedAt = time.Now()
	return logmails, nil
}
