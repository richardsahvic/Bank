package service

import (
	"repo"
)

type UserService interface {
	Register(userRegister repo.User, deposit int) (bool, error)
}

var User UserService

func NewService(userRepo repo.BankRepository) {
	User = NewUserService(userRepo)
}
