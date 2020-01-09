package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"Go-blog-server/internal/models"
)

type UserDao struct {
	*BaseDaoInitializer
}

func (ud *UserDao) FindOneUser(condition interface{}) (models.UserModel, error) {
	var user models.UserModel
	found := ud.FindOne(func(db *gorm.DB) *gorm.DB {
		return db.Where(condition)
	}, &user)
	fmt.Println(found)

	return user, found
}