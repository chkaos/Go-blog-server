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
func (d *CategoryDAO) AddCategory(Category *models.Category) error {
	return d.db().Create(Category).Error
}

func (d *CategoryDAO) UpdateCategory(Category *models.Category) error {
	return d.db().Model(&models.Category{}).Update(Category).Error
}

// QueryCategorys  query all Categorys
func (d *CategoryDAO) QueryAllCategorys() (Categorys []*models.Category, err error) {
	err = d.db().Find(&Categorys).Error
	return
}

// QueryCategory query Category by Category name
func (d *CategoryDAO) QueryCategoryByName(name string) (Category *models.Category, err error) {
	Category = &models.Category{}
	if err = d.db().Where("name = ?", name).First(&Category).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// QueryCategory query Category by Category id
func (d *CategoryDAO) QueryCategoryByID(id int) (Category *models.Category, err error) {
	Category = &models.Category{}
	if err = d.db().Where("id = ?", id).First(&Category).Error; err != nil {
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

func (d *CategoryDAO) QueryCategorys(req *models.QueryCategoryReq) (total int, Categorys []*models.Category, err error) {
	db := d.db().Preload("Articles").Model(&models.Category{})

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = db.Find(&Categorys).Error

	return
}
