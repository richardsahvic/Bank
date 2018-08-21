package repo

type BankRepository interface {
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	FindByPhone(phone string) (User, error)
	InsertNewUser(user User) (bool, error)
}
