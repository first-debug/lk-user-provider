package models

type Student struct {
	ID           int `gorm: primaryKey`
	Email        string
	HashPassword string
	Role         string
}
