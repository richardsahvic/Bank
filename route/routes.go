package route

import (
	"handler"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	route.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	route.HandleFunc("/checkbalance", handler.CheckBalanceHandler).Methods("GET")
	route.HandleFunc("/transactionbyreciever", handler.CheckTransactionRecieverHandler).Methods("GET")
	route.HandleFunc("/changepassword", handler.ChangePasswordHandler).Methods("POST")
	route.HandleFunc("/deleteaccount", handler.DeleteAccountHandler).Methods("POST")
	route.HandleFunc("/deposit", handler.DepositHandler).Methods("POST")
	route.HandleFunc("/withdraw", handler.WithdrawalHandler).Methods("POST")
	route.HandleFunc("/transfer", handler.TransferHandler).Methods("POST")
	return route
}
