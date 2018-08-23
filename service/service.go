package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.User) (bool, error)
	Login(username string, password string) (string, error)
	CheckBalance(token string) (string, error)
}

var User UserService

func NewService(userRepo repo.BankRepository) {
	User = NewUserService(userRepo)
}
