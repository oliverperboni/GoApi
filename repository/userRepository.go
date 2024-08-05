package repository

import (
	"github.com/oliverperboni/GoTomekeeper/schemas"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func CreateUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user schemas.User) error {
	return u.DB.Save(user).Error
}

func (u *UserRepository) GetUser(username string, password string) error {
	return u.DB.Find("username = ? AND passwordhash = ?", username, password).Error
}
