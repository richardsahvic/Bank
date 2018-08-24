package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.User) (bool, error)
	Login(username string, password string) (string, error)
	CheckBalance(token string) (string, error)
	ChangePassword(token string, password string, newPassword string) (bool, error)
	DeleteAccount(token string, password string) (bool, error)
}

var User UserService

func NewService(userRepo repo.BankRepository) {
	User = NewUserService(userRepo)
}
