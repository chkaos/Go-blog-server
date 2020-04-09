package models

import "time"

type SysConf struct {
	ID                 int    `json:"id" form:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	OssAccessKeyID     string `json:"oss_access_key_id"`
	OssAccessKeySercet string `json:"oss_access_key_sercet"`
}

type File struct {
	Model

	FileName string `json:"file_name"`
	URL      string `json:"url"`
	Type     string `json:"type"`
	Size     int    `json:"size"`
}

type QueryFileReq struct {
	Pagination
}

type FileResponse struct {
	ID         int        `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	FileName   string     `json:"file_name"`
	URL        string     `json:"url"`
	Type       string     `json:"type"`
	Size       int        `json:"size"`
}

type FilesSerializer struct {
	Files []*File
}

func (f *File) Response() FileResponse {
	File := FileResponse{
		ID:         f.ID,
		CreatedAt:  f.CreatedAt,
		ModifiedAt: f.ModifiedAt,
		FileName:   f.FileName,
		URL:        f.URL,
		Type:       f.Type,
		Size:       f.Size,
	}
	return File
}

func (f *FilesSerializer) Response() []FileResponse {
	var Files []FileResponse
	for _, File := range f.Files {
		Files = append(Files, File.Response())
	}
	return Files
}

func (SysConf) TableName() string {
	return "system_config"
}
