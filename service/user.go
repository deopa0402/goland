package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main/domain"
	"main/repository"
	"time"
)

var SecretKey = []byte("secret")

func GetUser(id int) (domain.User, error) {
	user := repository.FindUserByID(id)
	return user, nil
}

func CreateUser(User domain.User) domain.User {
	return repository.SaveUser(User)
}

func NewJWT(id int) (domain.LoginRes, error) {
	claims := &domain.JWTClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	res := domain.LoginRes{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SecretKey)
	if err != nil {
		return res, err
	}
	res.JWT = t

	return res, nil
}

func Login(login domain.Login) (bool, error) {
	user := repository.FindUserByID(login.ID)
	if user == (domain.User{}) {
		return false, fmt.Errorf("user not found")
	}
	if user.Password != login.Password {
		return false, fmt.Errorf("password is not correct")
	}
	return true, nil
}
