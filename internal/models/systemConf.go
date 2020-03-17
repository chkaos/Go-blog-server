package models

type SysConf struct {
	ID                 int    `json:"id" form:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	OssAccessKeyID     string `json:"oss_access_key_id"`
	OssAccessKeySercet string `json:"oss_access_key_sercet"`
}

func (SysConf) TableName() string {
	return "system_config"
}
