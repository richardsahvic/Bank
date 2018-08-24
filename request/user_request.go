package request

type Response struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

type DeleteRequest struct {
	Password string `json:"password"`
}

type DepositandWithdrawRequest struct {
	Amount int `json:"amount"`
}
