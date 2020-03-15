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

func (d *TagDAO) UptadeTag(tag *models.Tag) error {
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

// AddTagRelation add tag relation of target
func (d *TagDAO) AddTagRelation(relation *models.TagRelation) error {
	return d.db().Create(relation).Error
}

func (d *TagDAO) QueryTags(req *models.QueryTagReq) (total int, tags []*models.Tag, err error) {
	Db := d.db().Preload("Articles").Model(&models.Tag{})

	if req.Name != "" {
		name := utils.FuzzyInquiry(req.Name)
		Db = Db.Where("name LIKE ?", name)
	}

	if err = Db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		Db = Db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = Db.Find(&tags).Error

	return
}

// QueryTagRelation query tag relation
// func (d *TagDAO) QueryTagRelation(lre *models.TagRelation) (lr []*models.TagRelation, err error) {
// 	err = d.db().
// 		Where("active = ?", 1).
// 		Where(models.TagRelation{Type: lre.Type, TagID: lre.TagID, TargetID: lre.TargetID}).Find(&lr).Error
// 	return
// }

// QueryTagRelationByIDs  Query tag relation by ids
// func (d *TagDAO) QueryTagRelationByIDs(ids []int) (lr []*models.TagRelation, err error) {
// 	err = d.db().Where("id in (?)", ids).Find(lr).Error
// 	return
// }

// CheckTagRelationExist check tag relation exist
// func (d *TagDAO) CheckTagRelationExist(id int) (result bool, err error) {
// 	result = false
// 	lr := &models.TagRelation{}
// 	err = d.db().Where("active = ?", 1).Where(" id = ?", id).First(lr).Error
// 	if lr.ID > 0 {
// 		result = true
// 	}
// 	return
// }

// QueryTagExist check tag exist
// func (d *TagDAO) QueryTagExist(lre *models.TagRelation) (lr *models.TagRelation, err error) {
// 	lr = &models.TagRelation{}
// 	err = d.db().
// 		Where("active = ?", 1).
// 		Where(models.TagRelation{Type: lre.Type, TagID: lre.TagID, TargetID: lre.TargetID}).First(lr).Error
// 	return
// }

// DeleteTagRelation delete relation of tag
func (d *TagDAO) DeleteTagRelation(id int) (err error) {
	return d.db().Where(" id = ?", id).Update("active", 0).Error
}
