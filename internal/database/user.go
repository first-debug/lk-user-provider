package database

type User struct {
	ID      int `gorm:"primaryKey"`
	Email   string
	Role    string
	Version int
}
