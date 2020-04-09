package models

import "time"

type Category struct {
	Model

	Name     string     `json:"name"`
	Desc     string     `json:"desc"`
	Articles []*Article `gorm:"foreignkey:category_id"`
}

type QueryCategoryReq struct {
	Pagination
}

type CategoryResponse struct {
	ID         int        `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc,omitempty"`
	ArticleNum int        `json:"article_num,omitempty" default:"0"`
}

type CategorysSerializer struct {
	Categorys []*Category
}

func (t *Category) Response() CategoryResponse {
	Category := CategoryResponse{
		ID:         t.ID,
		CreatedAt:  t.CreatedAt,
		ModifiedAt: t.ModifiedAt,
		Name:       t.Name,
		Desc:       t.Desc,
		ArticleNum: len(t.Articles),
	}
	return Category
}

func (t *Category) PreviewResponse() CategoryResponse {
	Category := CategoryResponse{
		ID:   t.ID,
		Name: t.Name,
	}
	return Category
}

func (s *CategorysSerializer) Response() []CategoryResponse {
	var Categorys []CategoryResponse
	for _, Category := range s.Categorys {
		Categorys = append(Categorys, Category.Response())
	}
	return Categorys
}

func (s *CategorysSerializer) PreviewResponse() []CategoryResponse {
	var Categorys []CategoryResponse
	for _, Category := range s.Categorys {
		Categorys = append(Categorys, Category.PreviewResponse())
	}
	return Categorys
}
