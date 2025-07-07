package database

import (
	"log"
	"main/internal/database/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

type Database interface {
	AuthUser(email string) *models.Student
}

func GetDB() Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Db_url := os.Getenv("DB_URL")
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
