package service

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"repo"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repo.BankRepository
}

// Token is a struct made to generate token
type Token struct {
	jwt.StandardClaims
}

var mySigningKey []byte
var minimalTransaction = 10000

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

func (s *userService) Register(userRegister repo.User) (success bool, err error) {
	success = false

	reEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	emailValid := reEmail.MatchString(userRegister.Email)
	if !emailValid {
		log.Println("Email's format is not valid.")
		return
	}

	checkEmail, err := s.userRepo.FindByEmail(userRegister.Email)
	newEmail := checkEmail.Email
	if len(newEmail) != 0 {
		success = false
		log.Printf("Email: %v is already exist", newEmail)
		return
	}

	checkUsername, err := s.userRepo.FindByUsername(userRegister.Username)
	newUsername := checkUsername.Username
	if len(newUsername) != 0 {
		success = false
		log.Printf("Username: %v is already exist", newUsername)
		return
	}

	checkPhone, err := s.userRepo.FindByPhone(userRegister.Phone)
	newPhone := checkPhone.Phone
	if len(newPhone) != 0 {
		success = false
		log.Printf("Phone: %v is already exist", newPhone)
		return
	}

	userRegister.Password, err = HashPassword(userRegister.Password)
	if err != nil {
		log.Println("Failed encrypting password,  ", err)
		return
	}

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println("Failed generating snowflake id,    ", err)
		return
	}
	id := node.Generate().String()

	userRegister.ID = id

	success, err = s.userRepo.InsertNewUser(userRegister)
	if err != nil {
		fmt.Println("Error at user_service.go, ", err)
		return
	}
	return
}

func (s *userService) Login(username string, password string) (token string, err error) {
	mySigningKey = []byte("TheSignatureofTheBank")

	userData, err := s.userRepo.FindByUsername(username)
	if err != nil {
		fmt.Println("Error at user service, getting user data: ", err)
		return
	}

	match := CheckPasswordHash(password, userData.Password)
	if !match {
		log.Println("Wrong password")
		return
	}

	claims := Token{
		jwt.StandardClaims{
			Subject:   userData.ID,
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	signing := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = signing.SignedString(mySigningKey)
	if len(token) == 0 {
		log.Println("Failed to generate token")
		return
	}
	return
}

func (s *userService) CheckBalance(token string) (balance string, err error) {
	var id string

	at(time.Unix(0, 0), func() {
		tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(tokenClaims *jwt.Token) (interface{}, error) {
			return []byte("IDKWhatThisIs"), nil
		})

		if claims, _ := tokenClaims.Claims.(*Token); claims.ExpiresAt > time.Now().Unix() {
			id = claims.StandardClaims.Subject
			log.Println(claims.Subject)
		} else {
			fmt.Println("token Invalid,    ", err)
		}
	})

	userData, err := s.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error at user service, getting balance: ", err)
		return
	}

	balance = strconv.Itoa(userData.Balance) + " IDR"

	return

}

func (s *userService) ChangePassword(token string, password string, newPassword string) (success bool, err error) {
	success = false

	var id string

	at(time.Unix(0, 0), func() {
		tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(tokenClaims *jwt.Token) (interface{}, error) {
			return []byte("IDKWhatThisIs"), nil
		})

		if claims, _ := tokenClaims.Claims.(*Token); claims.ExpiresAt > time.Now().Unix() {
			id = claims.StandardClaims.Subject
			log.Println(claims.Subject)
		} else {
			fmt.Println("token Invalid,    ", err)
		}
	})

	userData, err := s.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error at user service, getting balance: ", err)
		return
	}

	match := CheckPasswordHash(password, userData.Password)
	if !match {
		log.Println("Wrong password")
		return
	}

	hashedNewPass, err := HashPassword(newPassword)
	if err != nil {
		log.Println("Failed encrypting password,  ", err)
		return
	}

	success, err = s.userRepo.UpdatePassword(id, hashedNewPass)
	if err != nil {
		log.Println("Error at user service, updating password: ", err)
		return
	}
	return

}

func (s *userService) DeleteAccount(token string, password string) (success bool, err error) {
	success = false

	var id string

	at(time.Unix(0, 0), func() {
		tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(tokenClaims *jwt.Token) (interface{}, error) {
			return []byte("IDKWhatThisIs"), nil
		})

		if claims, _ := tokenClaims.Claims.(*Token); claims.ExpiresAt > time.Now().Unix() {
			id = claims.StandardClaims.Subject
			log.Println(claims.Subject)
		} else {
			fmt.Println("token Invalid,    ", err)
		}
	})

	userData, err := s.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error at user service, getting balance: ", err)
		return
	}

	match := CheckPasswordHash(password, userData.Password)
	if !match {
		log.Println("Wrong password")
		return
	}

	success, err = s.userRepo.DeleteAccount(id)
	if err != nil {
		log.Println("Error at user service, updating password: ", err)
		return
	}
	return
}

func (s *userService) Deposit(token string, amount int) (success bool, err error) {
	success = false

	if amount < minimalTransaction {
		log.Println("The amount of the transaction is too low")
		return
	}

	var id string

	at(time.Unix(0, 0), func() {
		tokenClaims, err := jwt.ParseWithClaims(token, &Token{}, func(tokenClaims *jwt.Token) (interface{}, error) {
			return []byte("IDKWhatThisIs"), nil
		})

		if claims, _ := tokenClaims.Claims.(*Token); claims.ExpiresAt > time.Now().Unix() {
			id = claims.StandardClaims.Subject
			log.Println(claims.Subject)
		} else {
			fmt.Println("token Invalid,    ", err)
		}
	})

	userData, err := s.userRepo.FindByID(id)
	if err != nil {
		fmt.Println("Error at user service, getting balance: ", err)
		return
	}

	balance := userData.Balance + amount

	success, err = s.userRepo.UpdateBalance(id, balance)
	if err != nil {
		log.Println("Error at user service, updating balance: ", err)
		return
	}
	return
}
