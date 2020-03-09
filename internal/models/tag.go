package models

import "time"

type Tag struct {
	Model

	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	Icon       string    `json:"icon"`
	Articles   []*Article `gorm:"many2many:tag_relation;association_jointable_foreignkey:article_id;jointable_foreignkey:tag_id"`
	ArticleNum int       `json:"article_num" sql:"-"`
}

type TagRelation struct {
	Model

	TagID     uint `json:"tag_id"`
	ArticleID uint `json:"article_id"`
}

type QueryTagReq struct {
	Name       string    `json:"name"`

	Pagination
}

type TagResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc,omitempty"`
	Icon       string    `json:"icon"`
	ArticleNum int       `json:"article_num,omitempty"`
}

type TagsSerializer struct {
	Tags []*Tag
}

func (t *Tag) Response() TagResponse {
	tag := TagResponse{
		ID:         t.ID,
		CreatedAt:  t.CreatedAt,
		ModifiedAt: t.ModifiedAt,
		Name:       t.Name,
		Desc:       t.Desc,
		Icon:       t.Icon,
		ArticleNum: len(t.Articles),
	}
	return tag
}

func (t *Tag) PreviewResponse() TagResponse {
	tag := TagResponse{
		ID:         t.ID,
		Name:       t.Name,
		Icon:       t.Icon,
	}
	return tag
}

func (s *TagsSerializer) Response() []TagResponse {
	var tags []TagResponse
	for _, tag := range s.Tags {
		tags = append(tags, tag.Response())
	}
	return tags
}

func (s *TagsSerializer) PreviewResponse() []TagResponse {
	var tags []TagResponse
	for _, tag := range s.Tags {
		tags = append(tags, tag.PreviewResponse())
	}
	return tags
}

//TagRelation db table name of tag relation
func (t *TagRelation) TableName() string {
	return "tag_relation"
}
