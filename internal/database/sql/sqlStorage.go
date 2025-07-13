package sql

import (
	"main/internal/database"

	"gorm.io/gorm"
)

type SQLStorage struct {
	DB *gorm.DB
}

func NewSQLStorage(dialector gorm.Dialector, config *gorm.Config) (*SQLStorage, error) {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	return &SQLStorage{DB: db}, nil
}

func (s *SQLStorage) AuthUser(email string) (*database.User, error) {
	user := &database.User{}
	if err := s.DB.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *SQLStorage) Migrate() {
	s.DB.AutoMigrate(&database.User{})
}
