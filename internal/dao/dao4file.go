package dao

import "Go-blog-server/internal/models"

type FileDAO struct {
	*Dao
}

// NewFileDAO creates a new FileDAO
func NewFileDAO() *FileDAO {
	return &FileDAO{}
}

func (f *FileDAO) GetOSSConf() (conf *models.SysConf, err error) {
	conf = &models.SysConf{}
	err = f.db().Where("id = ?", 1).First(&conf).Error
	return
}

func (f *FileDAO) AddFile(file *models.File) error {
	return f.db().Create(file).Error
}
