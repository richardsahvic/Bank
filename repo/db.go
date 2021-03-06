package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	conn                      *sqlx.DB
	findIDStmt                *sqlx.Stmt
	findEmailStmt             *sqlx.Stmt
	findPhoneStmt             *sqlx.Stmt
	findUsernameStmt          *sqlx.Stmt
	updatePasswordStmt        *sqlx.Stmt
	updateBalanceStmt         *sqlx.Stmt
	deleteAccountStmt         *sqlx.Stmt
	transactionByRecieverStmt *sqlx.Stmt
	insertNewUserStmt         *sqlx.NamedStmt
	newTransactionStmt        *sqlx.NamedStmt
}

func (db *userRepository) MustPrepareStmt(query string) *sqlx.Stmt {
	stmt, err := db.conn.Preparex(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %s\n", err)
		os.Exit(1)
	}
	return stmt
}

func (db *userRepository) MustPrepareNamedStmt(query string) *sqlx.NamedStmt {
	stmt, err := db.conn.PrepareNamed(query)
	if err != nil {
		fmt.Printf("Error preparing statement: %s\n", err)
		os.Exit(1)
	}
	return stmt
}

func NewRepository(db *sqlx.DB) BankRepository {
	r := userRepository{conn: db}
	r.findIDStmt = r.MustPrepareStmt("SELECT * FROM mybank.user_detail WHERE id=?")
	r.findEmailStmt = r.MustPrepareStmt("SELECT * FROM mybank.user_detail WHERE email=?")
	r.findPhoneStmt = r.MustPrepareStmt("SELECT * FROM mybank.user_detail WHERE phone=?")
	r.findUsernameStmt = r.MustPrepareStmt("SELECT * FROM mybank.user_detail WHERE username=?")
	r.updatePasswordStmt = r.MustPrepareStmt("UPDATE mybank.user_detail SET password=? WHERE id=?")
	r.updateBalanceStmt = r.MustPrepareStmt("UPDATE mybank.user_detail SET Balance=? WHERE id=?")
	r.deleteAccountStmt = r.MustPrepareStmt("DELETE FROM mybank.user_detail WHERE id=?")
	r.transactionByRecieverStmt = r.MustPrepareStmt("SELECT * FROM mybank.transactions WHERE sender_phone=? AND reciever_phone=?")
	r.insertNewUserStmt = r.MustPrepareNamedStmt("INSERT INTO mybank.user_detail (id, phone, email, username, password, balance) VALUES (:id, :phone, :email, :username, :password, :balance)")
	r.newTransactionStmt = r.MustPrepareNamedStmt("INSERT INTO mybank.transactions (id, sender_phone, reciever_phone, total) VALUES (:id, :sender_phone, :reciever_phone, :total)")
	return &r
}

func (db *userRepository) FindByID(id string) (usr User, err error) {
	err = db.findIDStmt.Get(&usr, id)
	if err != nil {
		log.Printf("ID: %v , doesn't exist", id)
		log.Println("Error at finding id:  ", err)
	}
	return
}

func (db *userRepository) FindByPhone(phone string) (usr User, err error) {
	err = db.findPhoneStmt.Get(&usr, phone)
	if err != nil {
		log.Printf("Phone: %v, doesn't exist", phone)
		log.Println("Error at finding phone:  ", err)
	}
	return
}

func (db *userRepository) FindByEmail(email string) (usr User, err error) {
	err = db.findEmailStmt.Get(&usr, email)
	if err != nil {
		log.Printf("Email: %v, doesn't exist", email)
		log.Println("Error at finding email:  ", err)
	}
	return
}

func (db *userRepository) FindByUsername(username string) (usr User, err error) {
	err = db.findUsernameStmt.Get(&usr, username)
	if err != nil {
		log.Printf("Username: %v doesn't exist", username)
		log.Println("Error at finding username:  ", err)
	}
	return
}

func (db *userRepository) UpdatePassword(id string, newPassword string) (success bool, err error) {
	_, err = db.updatePasswordStmt.Exec(newPassword, id)
	if err != nil {
		log.Println("Failed to update password: ", err)
		success = false
	}
	success = true
	return
}

func (db *userRepository) DeleteAccount(id string) (success bool, err error) {
	_, err = db.deleteAccountStmt.Exec(id)
	if err != nil {
		log.Println("Error deleting account: ", err)
		success = false
		return
	}
	success = true
	return
}

func (db *userRepository) UpdateBalance(id string, balance int) (success bool, err error) {
	_, err = db.updateBalanceStmt.Exec(balance, id)
	if err != nil {
		log.Println("Error updating balance: ", err)
		success = false
		return
	}
	success = true
	return
}

func (db *userRepository) InsertNewUser(newUser User) (success bool, err error) {
	_, err = db.insertNewUserStmt.Exec(newUser)
	if err != nil {
		log.Println("Error inserting new user:  ", err)
		success = false
		return
	}
	success = true
	return
}

func (db *userRepository) NewTransaction(transaction Transaction) (success bool, err error) {
	_, err = db.newTransactionStmt.Exec(transaction)
	if err != nil {
		log.Println("Error inserting new transaction to db:  ", err)
		success = false
		return
	}
	success = true
	return
}

func (db *userRepository) CheckTransaction(sender_phone string, reciever_phone string) (transactions []Transaction, err error) {
	err = db.transactionByRecieverStmt.Select(&transactions, sender_phone, reciever_phone)
	if err != nil {
		log.Println("Error getting transactions:  ", err)
	}
	return
}
