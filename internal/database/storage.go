package database

type UserStorage interface {
	GetUser(email string)
}
