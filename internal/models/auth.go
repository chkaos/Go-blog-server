package models

import (
	"fmt"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Avatar   string `json:"avatar"`
	Role     int    `json:"role"`
	Source   int    `json:"source"`
	IsMuted  int    `json:"is_muted"`
}

func CheckAuth(username, password string) (User, error) {
	var user User
	if err := db.Where(map[string]interface{}{"username": username, "password": password}).First(&user).Error; err != nil {
		fmt.Println(user, err)
		return user, err
	}

	if user.ID > 0 {
		return user, nil
	}

	return user, nil
}

// func CheckAuth(username, password string) (bool, User,) {
// 	var auth User
// 	err := db.Select("id").Where(User{Username: username, Password: password}).First(&auth).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}

// 	if auth.ID > 0 {
// 		return true, auth, nil
// 	}

// 	return false, auth, nil
// }
