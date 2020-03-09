package services

import (
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
)

type UserService struct {
	dao *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{dao: new(dao.UserDAO)}
}

// Auth  auth login by username and password
func (s *UserService) Auth(username, password string) (user *models.User, err error) {
	user, err = s.dao.QueryUser(&models.User{Username: username, Password: password})
	return
}
