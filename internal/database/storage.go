package database

type UserStorage interface {
	AuthUser(email string) (*User, error)
	Migrate()
}
