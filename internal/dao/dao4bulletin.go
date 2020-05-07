package dao

import (
	_ "fmt"

	"github.com/jinzhu/gorm"

	"Go-blog-server/internal/models"
)

type BulletinDAO struct {
	*Dao
}

// NewBulletinDAO creates a new BulletinDAO
func NewBulletinDAO() *BulletinDAO {
	return &BulletinDAO{}
}

// AddBulletin  add new Bulletin
func (b *BulletinDAO) AddBulletin(bulletin models.Bulletin) error {
	return b.db().Create(&bulletin).Error
}

func (b *BulletinDAO) UpdateBulletin(bulletin models.Bulletin) error {
	return b.db().Model(&bulletin).Update(&bulletin).Error
}

// QueryBulletin query Bulletin by Bulletin id
func (b *BulletinDAO) QueryBulletinByID(id int) (bulletin models.Bulletin, err error) {
	if err = b.db().Where("id = ?", id).First(&bulletin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// DeleteBulletin delete Bulletin
func (b *BulletinDAO) DeleteBulletin(id int) error {
	return b.db().Where("id = ?", id).Delete(&models.Bulletin{}).Error
}

func (b *BulletinDAO) QueryBulletins(req *models.QueryBulletinReq) (total int, bulletins []models.Bulletin, err error) {

	db := b.db().Model(&models.Bulletin{}).Order("created_at desc").Order("top desc")

	if req.Top >= 0 {
		db = db.Where("top = ?", req.Top)
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = db.Find(&bulletins).Error

	return
}
