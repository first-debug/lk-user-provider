package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLUserStorage struct {
	DB *gorm.DB
}

func NewMySQLUserStorage(db_url string) UserStorage {
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	db.AutoMigrate(&User{})
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
func (mysql_db *MySQLUserStorage) AuthUser(email string) *User {
	user := &User{}
	mysql_db.DB.Where("email = ?", email).First(user)
	return user
}
