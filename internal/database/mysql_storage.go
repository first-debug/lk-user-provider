package database

import (
	"main/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLUserStorage struct {
	DB *gorm.DB
}

func NewMySQLUserStorage(cfg *config.Config) UserStorage {
	db, err := gorm.Open(mysql.Open(cfg.DB_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &MySQLUserStorage{DB: db}
}

func (mysql_db *MySQLUserStorage) GetUser(email string) *User {
	user := &User{}
	mysql_db.DB.Where("email = ?", email).First(user)
	return user
}
