package dao

import (
	_ "fmt"

	"github.com/jinzhu/gorm"

	"Go-blog-server/internal/models"
)

type CategoryDAO struct {
	*Dao
}

// NewCategoryDAO creates a new CategoryDAO
func NewCategoryDAO() *CategoryDAO {
	return &CategoryDAO{}
}

// AddCategory  add new Category
func (d *CategoryDAO) AddCategory(category models.Category) error {
	return d.db().Create(&category).Error
}

func (d *CategoryDAO) UpdateCategory(category models.Category) error {
	return d.db().Model(&category).Update(&category).Error
}

// QueryCategorys  query all Categorys
func (d *CategoryDAO) QueryAllCategorys() (categorys []models.Category, err error) {
	err = d.db().Find(&categorys).Error
	return
}

// QueryCategory query Category by Category name
func (d *CategoryDAO) QueryCategoryByName(name string) (category models.Category, err error) {
	if err = d.db().Where("name = ?", name).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// QueryCategory query Category by Category id
func (d *CategoryDAO) QueryCategoryByID(id int) (category models.Category, err error) {
	if err = d.db().Where("id = ?", id).First(&category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// DeleteCategory delete Category
func (d *CategoryDAO) DeleteCategory(id int) error {
	return d.db().Where("id = ?", id).Delete(&models.Category{}).Error
}

func (d *CategoryDAO) QueryCategorys(req *models.QueryCategoryReq) (total int, categorys []models.Category, err error) {
	db := d.db().Preload("Articles").Model(&models.Category{})

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = db.Find(&categorys).Error

	return
}
