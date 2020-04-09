package models

import "time"

type Tag struct {
	Model

	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Icon     string    `json:"icon"`
	Articles []Article `gorm:"many2many:tag_relation;association_jointable_foreignkey:article_id;jointable_foreignkey:tag_id"`
}

type TagRelation struct {
	Model

	TagID     int `json:"tag_id"`
	ArticleID int `json:"article_id"`
}

type QueryTagReq struct {
	Name string `json:"name"`

	Pagination
}

type TagResponse struct {
	ID         int        `json:"id"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc,omitempty"`
	Icon       string     `json:"icon"`
	ArticleNum int        `json:"article_num" default:"0"`
}

type TagsSerializer struct {
	Tags []Tag
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
		ID:   t.ID,
		Name: t.Name,
		Icon: t.Icon,
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

func (*Tag) TableName() string {
	return "tag"
}

//TagRelation db table name of tag relation
func (TagRelation) TableName() string {
	return "tag_relation"
}
