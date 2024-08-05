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

func (u *UserRepository) CreateUser(user *schemas.User) error {
	return u.DB.Save(user).Error
}

func (u *UserRepository) GetUser(username string, password string) (*schemas.User, error) {
	var user schemas.User
	err := u.DB.Where("username = ? AND password_hash = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
