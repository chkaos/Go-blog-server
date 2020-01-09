package services

import (
	"fmt"

	"Go-blog-server/internal/models"
	"Go-blog-server/internal/dao"
)

type UserService struct {
	dao *dao.UserDao
}

func NewUserService() *UserService {
	return &UserService{dao: new(dao.UserDao)}
}

func (us *UserService)CheckAuth(username, password string) (models.UserModel, error) {
	userModel, err := us.dao.FindOneUser(map[string]interface{}{"username": username, "password": password})
	fmt.Println(userModel)
	if userModel.ID > 0 {
		return userModel, err
	}

	return userModel, err
}

