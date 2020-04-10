package dao

import "Go-blog-server/internal/models"

type FileDAO struct {
	*Dao
}

// NewFileDAO creates a new FileDAO
func NewFileDAO() *FileDAO {
	return &FileDAO{}
}

func (f *FileDAO) GetOSSConf() (conf models.SysConf, err error) {
	conf = models.SysConf{}
	err = f.db().Where("id = ?", 1).First(&conf).Error
	return
}

func (f *FileDAO) AddFile(file models.File) error {
	return f.db().Create(file).Error
}

func (d *FileDAO) QueryFiles(req *models.QueryFileReq) (total int, Files []models.File, err error) {
	db := d.db().Model(&models.File{}).Order("created_at desc")

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = db.Find(&Files).Error

	return
}
