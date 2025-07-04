package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/internal/database/models"
)

type DB struct {
	DB *gorm.DB
}

type Database interface {
	AuthUser(email string) *models.Student
}

func GetDB(Db_url string) Database {
	db, err := gorm.Open(mysql.Open(Db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{DB: db}
}

func (db *DB) AuthUser(email string) *models.Student {
	student := &models.Student{}
	db.DB.Where("email = ?", email).First(student)
	return student
}
