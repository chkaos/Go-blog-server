package models

type Category struct {
	Model

	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Articles []Article `gorm:"foreignkey:category_id"`
}
