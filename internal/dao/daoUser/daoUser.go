package daoUser

import (
	"Go-blog-server/internal/common"
	"Go-blog-server/internal/models"
)

const userTableName = "user"

func FindOne(condition interface{}) (models.UserModel, error) {
	db := common.GetDB()
	var model models.UserModel
	err := db.Table(userTableName).Where(condition).First(&model).Error
	return model, err
}