package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"service"

	"datasource"
	"repo"
	"request"
)

var db = datasource.InitConnection()
var userService = service.NewUserService(repo.NewRepository(db))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var regRequest request.RegisterRequest
	json.Unmarshal(body, &regRequest)

	userRegister := repo.User{
		Email:    regRequest.Email,
		Phone:    regRequest.Phone,
		Username: regRequest.Username,
		Password: regRequest.Password,
		Balance:  regRequest.Balance,
	}

	registerResult, err := userService.Register(userRegister)
	if err != nil {
		log.Println("failed to register,    ", err)
	}

	var regResponse request.Response

	if !registerResult {
		regResponse.Message = "Register failed"
	} else {
		regResponse.Message = "Register success"
	}

	json.NewEncoder(w).Encode(regResponse)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var loginReq request.LoginRequest
	json.Unmarshal(body, &loginReq)

	loginToken, err := userService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var loginResp request.Response

	if len(loginToken) == 0 {
		loginResp.Message = "Login failed"
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.Header().Set("token", loginToken)
	}
}

func CheckBalanceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	balance, err := userService.CheckBalance(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balance = "balance: " + balance

	json.NewEncoder(w).Encode(balance)
}
