package dao

import (
	"Go-blog-server/internal/common"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	DB *gorm.DB
}

func (d *Dao) db() *gorm.DB {
	return common.GetDB()
}
