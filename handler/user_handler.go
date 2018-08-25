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
		loginResp.Message = "Logged in"
		w.Header().Set("token", loginToken)
	}
	json.NewEncoder(w).Encode(loginResp)
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

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var changePasswordReq request.ChangePasswordRequest
	json.Unmarshal(body, &changePasswordReq)

	success, err := userService.ChangePassword(token, changePasswordReq.Password, changePasswordReq.NewPassword)
	if err != nil {
		log.Println("Failed to register: ", err)
	}

	var changePwResp request.Response

	if !success {
		changePwResp.Message = "Failed to change password"
	} else {
		changePwResp.Message = "Password changed"
	}
	json.NewEncoder(w).Encode(changePwResp)
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var deleteReq request.DeleteRequest
	json.Unmarshal(body, &deleteReq)

	success, err := userService.DeleteAccount(token, deleteReq.Password)
	if err != nil {
		log.Println("Failed to delete account: ", err)
	}

	var deleteResp request.Response

	if !success {
		deleteResp.Message = "Failed to delete account"
	} else {
		deleteResp.Message = "Account Deleted"
	}

	json.NewEncoder(w).Encode(deleteResp)
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var depositReq request.DepositandWithdrawRequest
	json.Unmarshal(body, &depositReq)

	success, err := userService.Deposit(token, depositReq.Amount)
	if err != nil {
		log.Println("Failed to deposit: ", err)
	}

	var depositResp request.Response

	if !success {
		depositResp.Message = "Failed to deposit money"
	} else {
		depositResp.Message = "Deposit success"
	}

	json.NewEncoder(w).Encode(depositResp)
}

func WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var withdrawReq request.DepositandWithdrawRequest
	json.Unmarshal(body, &withdrawReq)

	success, err := userService.Withdrawal(token, withdrawReq.Amount)
	if err != nil {
		log.Println("Failed to deposit: ", err)
	}

	var depositResp request.Response

	if !success {
		depositResp.Message = "Failed to withdraw money"
	} else {
		depositResp.Message = "Withdraw success"
	}

	json.NewEncoder(w).Encode(depositResp)
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("token")

	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 5000))

	var transferReq request.TransferRequest
	json.Unmarshal(body, &transferReq)

	success, err := userService.Transfer(token, transferReq.DestPhone, transferReq.Amount)
	if err != nil {
		log.Println("Failed to transfer: ", err)
	}

	var transferResp request.Response

	if !success {
		transferResp.Message = "Transaction failed"
	} else {
		transferResp.Message = "Transaction success"
	}

	json.NewEncoder(w).Encode(transferResp)
}
