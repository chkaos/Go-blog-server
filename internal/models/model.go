package models

import "time"

type Model struct {
	ID         int        `json:"id" form:"id" gorm:"AUTO_INCREMENT;primary_key;"`
	CreatedAt  *time.Time `json:"created_at" gorm:"default:current_time"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"default:current_time"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// Pagination Pagination.
type Pagination struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type PaginationRep struct {
	PageSize int         `json:"page_size"`
	PageNum  int         `json:"page_num"`
	Total    int         `json:"total"`
	List     interface{} `json:"list,omitempty"`
}
