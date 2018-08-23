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
	return route
}
