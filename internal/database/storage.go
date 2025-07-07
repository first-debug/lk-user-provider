package database

type UserStorage interface {
	GetUser(email string) *User
	AuthUser(email string) *User
}
