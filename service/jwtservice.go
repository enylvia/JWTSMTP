package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTService interface {
	GenerateJWT(username string, id int) (string, error)
	ValidateJWT(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewJWTService() *jwtService{
	return &jwtService{}
}
var mySigningKey = []byte("")
func (s *jwtService) GenerateJWT(username string, id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something wrong when generate token")
		return "", err
	}
	return tokenString, nil
}

func (s *jwtService) ValidateJWT(encodedToken string) (*jwt.Token,error){
	token ,err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_,ok:= token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil , errors.New("invalid token")
		}
		return []byte(mySigningKey),nil
	})

	if err != nil {
		return token,err
	}
	return token,nil
}
