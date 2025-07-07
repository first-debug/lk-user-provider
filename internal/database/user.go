package database

type User struct {
	ID           int32 `gorm:"primaryKey"`
	Email        string
	HashPassword string
	Role         string
	Version      int32
}
