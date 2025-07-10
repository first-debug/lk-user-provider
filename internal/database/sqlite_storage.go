package database

import (
	"log/slog"
	sl "main/libs/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteUserStorage struct {
	DB  *gorm.DB
	log *slog.Logger
}

func NewSQLiteUserStorage(dbUrl string, log *slog.Logger) (UserStorage, error) {
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect database", sl.Err(err))
		return nil, err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Error("failed to migrate database", sl.Err(err))
		return nil, err
	}
	log.Info("connected to database and successfully migrated")
	return &SQLiteUserStorage{
		DB:  db,
		log: log,
	}, nil
}

func (s *SQLiteUserStorage) GetUser(email string) *User {
	user := &User{}
	s.DB.Where("email = ?", email).First(user)
	return user
}

func (s *SQLiteUserStorage) AuthUser(email string) *User {
	user := &User{}
	s.DB.Where("email = ?", email).First(user)
	return user
}
