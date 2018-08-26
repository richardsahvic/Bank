package repo

import (
	"time"
)

// User is table user_detail's model
type User struct {
	ID        string    `json:"id" db:"id"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Balance   int       `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Transaction is table transaction's model
type Transaction struct {
	ID            string    `json:"trans_id" db:"id"`
	SenderPhone   string    `json:"sender_phone" db:"sender_phone"`
	RecieverPhone string    `json:"reciever_phone" db:"reciever_phone"`
	Total         int       `json:"total" db:"total"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
