package authSrv

import (
	"fmt"

	"Go-blog-server/internal/models"
	"Go-blog-server/internal/dao/daoUser"
)

func CheckAuth(username, password string) (models.UserModel, error) {
	userModel, err := daoUser.FindOne(map[string]interface{}{"username": username, "password": password})
	fmt.Println(userModel)
	if userModel.ID > 0 {
		return userModel, err
	}

	return userModel, err
}

