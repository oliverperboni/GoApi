package services

import (
	"github.com/oliverperboni/GoTomekeeper/repository"
	"github.com/oliverperboni/GoTomekeeper/schemas"
)

type UserService struct {
	repo repository.UserRepository
}

func CreateUserService(r repository.UserRepository) UserService {
	return UserService{repo: r}
}

func (u *UserService) CreateUser(user schemas.User) error {
	return u.repo.CreateUser(user)
}

func (u *UserService) GetUser(username string, password string) error {
	return u.repo.GetUser(username, password)
}
