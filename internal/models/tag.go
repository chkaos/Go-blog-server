package models

import "time"

type Tag struct {
	Model

	Name     string    `json:"name"`
	Desc     string    `json:"desc"`
	Slug     string    `json:"slug"`
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
	Slug       string     `json:"slug,omitempty"`
	ArticleNum int        `json:"article_num,omitempty"`
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
		Slug:       t.Slug,
		ArticleNum: len(t.Articles),
	}
	return tag
}

func (t *Tag) PreviewResponse() TagResponse {
	tag := TagResponse{
		ID:   t.ID,
		Name: t.Name,
		Slug: t.Slug,
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
