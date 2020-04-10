package validators

import "Go-blog-server/internal/models"

type AuthForm struct {
	Username string `form:"username" valid:"Required;MaxSize(20)"`
	Password string `form:"password" valid:"Required;MinSize(6);MaxSize(20)"`
}

func (a *AuthForm) Transform() models.User {
	return models.User{
		Username: a.Username,
		Password: a.Password,
	}
}
