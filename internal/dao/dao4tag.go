package dao

import (
	_ "fmt"

	"github.com/jinzhu/gorm"

	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/utils"
)

type TagDAO struct {
	*Dao
}

// NewTagDAO creates a new TagDAO
func NewTagDAO() *TagDAO {
	return &TagDAO{}
}

// AddTag  add new tag
func (d *TagDAO) AddTag(tag *models.Tag) error {
	return d.db().Create(tag).Error
}

// UpdateTag  update tag
func (d *TagDAO) UpdateTag(tag *models.Tag) error {
	return d.db().Model(&models.Tag{}).Update(tag).Error
}

// Querytags  query all tags
func (d *TagDAO) QueryAllTags() (tags []*models.Tag, err error) {
	err = d.db().Find(&tags).Error
	return
}

// QueryTag query tag by tag name
func (d *TagDAO) QueryTagByName(name string) (tag *models.Tag, err error) {
	tag = &models.Tag{}
	if err = d.db().Where("name = ?", name).First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// QueryTag query tag by tag id
func (d *TagDAO) QueryTagByID(id int) (tag *models.Tag, err error) {
	tag = &models.Tag{}
	if err = d.db().Where("id = ?", id).First(&tag).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// DeleteTag delete tag
func (d *TagDAO) DeleteTag(id int) error {
	return d.db().Where("id = ?", id).Delete(&models.Tag{}).Error
}

// QueryTags query tags by condition
func (d *TagDAO) QueryTags(req *models.QueryTagReq) (total int, tags []*models.Tag, err error) {
	db := d.db().Preload("Articles").Model(&models.Tag{})

	if req.Name != "" {
		name := utils.FuzzyInquiry(req.Name)
		db = db.Where("name LIKE ?", name)
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = db.Find(&tags).Error

	return
}

// SetTags set article-tag-relation
func (d *TagDAO) SetTags(tagIds []int) (tags []models.Tag, err error) {
	err = db.Where(tagIds).Find(&tags).Error

	return
}
