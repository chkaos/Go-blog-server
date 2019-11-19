package models

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

func CheckAuth(username, password string) (User, bool) {
	var user User
	var isExist bool
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.ID > 0 {
		isExist = true
	}

	return user, isExist
}
