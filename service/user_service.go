package service

import (
	"log"
	"regexp"
	"time"

	"repo"
)

type userService struct {
	userRepo repo.BankRepository
}

// Token is a struct made to generate token
type Token struct {
	jwt.StandardClaims
}

var mySigningKey []byte

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func NewUserService(userRepo repo.BankRepository) UserService {
	s := userService{userRepo: userRepo}
	return &s
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *userService) Register(userRegister repo.User, deposit int) (bool, error) {
	success = false

	reEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailValid := reEmail.MatchString(userRegister.Email)
	if !emailValid {
		log.Println("Email format is not valid.")
		return
	}

	checkEmail, err := s.userRepo.FindByEmail(userRegister.Email)
	if len(checkEmail.Email) != 0 {
		return
	}

	checkUsername, err := s.userRepo.F
}
