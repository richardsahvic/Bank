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
	Deposit(token string, amount int) (bool, error)
	Withdrawal(token string, amount int) (bool, error)
	Transfer(token string, destPhone string, amount int) (bool, error)
	TransactionByReciever(token string, recieverPhone string) ([]repo.Transaction, error)
}

var User UserService

func NewService(userRepo repo.BankRepository) {
	User = NewUserService(userRepo)
}
