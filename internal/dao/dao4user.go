package dao

import (
	"Go-blog-server/internal/models"
)

type UserDAO struct {
	*Dao
}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// AddUser add user by user object
func (d *UserDAO) Add(user models.User) error {
	return d.DB.Create(user).Error
}

// func (d *UserDAO) AuthUser(username string, password string) (user *models.User, err error) {
// 	user = &models.User{}
// 	fmt.Println(user)
// 	err = d.db().Where("username=? AND password=?", username, password).First(user).Error
// 	return
// }

func (d *UserDAO) QueryUser(u *models.User) (user models.User, err error) {
	user = models.User{}
	err = d.db().Where(u).First(user).Error
	return
}

// QueryUserByUserName query user info by userName
func (d *UserDAO) QueryUserByUserName(userName string) (rows *models.User, err error) {
	rows = &models.User{}
	if err = d.DB.Where("name = ?", userName).First(rows).Error; err != nil {

	}
	return
}

// UpdateUser update user info
func (d *UserDAO) Update(user *models.User) error {
	return d.DB.Model(&models.User{}).Update(user).Where("ID=?", user.ID).Error
}
