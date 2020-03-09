package models

import (
	"time"
)

type Article struct {
	Model

	CategoryID      uint      `json:"category_id"`
	Category        Category  `json:"category";association_foreignkey:category_id`
	Title           string    `json:"title"`
	Desc            string    `json:"desc"`
	Keywords        string    `json:keywords`
	Content         string    `json:"content"`
	RenderedContent string    `json:"rendered_content"`
	Tags            []*Tag     `gorm:"many2many:mapping-article-tag;association_jointable_foreignkey:tag_id;jointable_foreignkey:article_id"`
	Published_At    *time.Time `json:"published_at"`
	Source          int       `json:"source"`
	ReproduceURL    string    `json:"reproduce_url"`
	Thumb           string    `json:"thumb"`
	LikesNum        int       `json:"like_num", gorm:"-"`
	PvsNm           int       `json:"pvs_num", gorm:"-"`
	CommentNum      int       `json:"comments_num", gorm:"-"`
	State           int       `json:"state"`
}

type ArticleResponse struct {
	ID              uint          `json:"id"`
	CreatedAt       *time.Time     `json:"created_at,omitempty"`
	ModifiedAt      *time.Time     `json:"modified_at,omitempty"`
	CategoryID      uint          `json:"category_id"`
	Category        Category  `json:"category"`
	Title           string        `json:"title"`
	Desc            string        `json:"desc"`
	Keywords        string        `json:keywords`
	Content         string        `json:"content"`
	RenderedContent string        `json:"rendered_content"`
	Tags            []TagResponse `json:"tags"`
	Published_At    *time.Time     `json:"published_at"`
	Source          int           `json:"source"`
	ReproduceURL    string        `json:"reproduce_url"`
	Thumb           string        `json:"thumb"`
	LikesNum        int           `json:"like_num"`
	PvsNm           int           `json:"pvs_num"`
	CommentNum      int           `json:"comments_num"`
	State           int           `json:"state"`
}

type ArticlesSerializer struct {
	Articles []*Article
}

func (a *Article) Response() ArticleResponse {
	article := ArticleResponse{
		ID:              a.ID,
		CategoryID:      a.CategoryID,
		Title:           a.Title,
		Desc:            a.Desc,
		Keywords:        a.Keywords,
		Content:         a.Content,
		RenderedContent: a.RenderedContent,
		CreatedAt:       a.CreatedAt,
		ModifiedAt:      a.ModifiedAt,
		Published_At:    a.Published_At,
		Source:          a.Source,
		ReproduceURL:    a.ReproduceURL,
		Thumb:           a.Thumb,
		LikesNum:        a.LikesNum,
		PvsNm:           a.PvsNm,
		CommentNum:      a.CommentNum,
		State:           a.State,
	}
	serializer := TagsSerializer{Tags: a.Tags}
	article.Tags = serializer.Response()

	return article
}

func (s *ArticlesSerializer) Response() []ArticleResponse {
	var articles []ArticleResponse
	for _, article := range s.Articles {
		articles = append(articles, article.Response())
	}
	return articles
}
