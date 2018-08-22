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
