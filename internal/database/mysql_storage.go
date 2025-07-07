package database

import (
	"log/slog"
	sl "main/libs/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLUserStorage struct {
	DB  *gorm.DB
	log *slog.Logger
}

func NewMySQLUserStorage(db_url string, log *slog.Logger) (UserStorage, error) {
	db, err := gorm.Open(mysql.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database", sl.Err(err))
		return nil, err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Error("failed to migrate database", sl.Err(err))
		return nil, err
	}
	log.Info("connected to database")
	return &MySQLUserStorage{
		DB:  db,
		log: log,
	}, nil
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
