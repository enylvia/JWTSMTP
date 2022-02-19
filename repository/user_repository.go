package repository

import (
	"gorm.io/gorm"
	"jwtsmtp/entity"
	"jwtsmtp/helper"
)

type Repository interface {
	Create(user entity.User) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	CreateLog (logs entity.LogMail) (entity.LogMail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	helper.ErrorIfNotNil(err)

	return user,nil
}

func (r *repository) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	helper.ErrorIfNotNil(err)

	return user, nil
}

func (r *repository)CreateLog (logs entity.LogMail) (entity.LogMail, error){
	err := r.db.Create(&logs).Error

	helper.ErrorIfNotNil(err)
	return logs,nil
}

