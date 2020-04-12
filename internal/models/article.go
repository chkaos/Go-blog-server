package models

import (
	"time"
)

type Article struct {
	Model

	CategoryID      int       `json:"category_id"`
	Category        Category  `json:"category" gorm:"foreignkey:CategoryID"`
	Title           string    `json:"title"`
	Desc            string    `json:"desc"`
	Keywords        string    `json:"keywords"`
	Content         string    `json:"content"`
	RenderedContent string    `json:"rendered_content"`
	Tags            []Tag     `gorm:"many2many:tag_relation;association_jointable_foreignkey:tag_id;jointable_foreignkey:article_id"`
	PublishedAt     time.Time `json:"published_at"`
	Source          int       `json:"source"`
	ReproduceURL    string    `json:"reproduce_url"`
	Thumb           string    `json:"thumb"`
	LikesNum        int       `json:"like_num" gorm:"-"`
	PvsNm           int       `json:"pvs_num" gorm:"-"`
	CommentNum      int       `json:"comments_num" gorm:"-"`
	State           int       `json:"state"`
}

type QueryArticleReq struct {
	Tag      int `json:"tag"`
	Category int `json:"category"`
	State    int `json:"state"`
	Source   int `json:"source"`
	Pagination
}

type ArticleResponse struct {
	ID              int              `json:"id"`
	CreatedAt       *time.Time       `json:"created_at,omitempty"`
	ModifiedAt      *time.Time       `json:"modified_at,omitempty"`
	CategoryID      int              `json:"category_id,omitempty"`
	Category        CategoryResponse `json:"category,omitempty"`
	Title           string           `json:"title"`
	Desc            string           `json:"desc"`
	Keywords        string           `json:"keywords"`
	Content         string           `json:"content"`
	RenderedContent string           `json:"rendered_content"`
	Tags            []TagResponse    `json:"tags"`
	PublishedAt     time.Time        `json:"published_at"`
	Source          int              `json:"source"`
	ReproduceURL    string           `json:"reproduce_url"`
	Thumb           string           `json:"thumb"`
	LikesNum        int              `json:"like_num,omitempty"`
	PvsNm           int              `json:"pvs_num,omitempty"`
	CommentNum      int              `json:"comments_num,omitempty"`
	State           int              `json:"state"`
}

type ArticlesSerializer struct {
	Articles []Article
}

func (a *Article) HandlePublishedAt() {
	if a.State == 0 && a.PublishedAt.IsZero() {
		a.PublishedAt = time.Now()
	}
}

func (a *Article) Response() ArticleResponse {
	article := ArticleResponse{
		ID:              a.ID,
		Title:           a.Title,
		Desc:            a.Desc,
		Keywords:        a.Keywords,
		Content:         a.Content,
		RenderedContent: a.RenderedContent,
		CreatedAt:       a.CreatedAt,
		ModifiedAt:      a.ModifiedAt,
		PublishedAt:     a.PublishedAt,
		Source:          a.Source,
		ReproduceURL:    a.ReproduceURL,
		Thumb:           a.Thumb,
		LikesNum:        a.LikesNum,
		PvsNm:           a.PvsNm,
		CommentNum:      a.CommentNum,
		State:           a.State,
	}
	serializer := TagsSerializer{Tags: a.Tags}
	article.Tags = serializer.PreviewResponse()
	article.Category = a.Category.PreviewResponse()

	return article
}

func (a *Article) PreviewResponse() ArticleResponse {
	article := ArticleResponse{
		ID:         a.ID,
		Title:      a.Title,
		Desc:       a.Desc,
		Keywords:   a.Keywords,
		CreatedAt:  a.CreatedAt,
		Source:     a.Source,
		Thumb:      a.Thumb,
		LikesNum:   a.LikesNum,
		PvsNm:      a.PvsNm,
		CommentNum: a.CommentNum,
	}
	serializer := TagsSerializer{Tags: a.Tags}
	article.Tags = serializer.PreviewResponse()
	article.Category = a.Category.PreviewResponse()

	return article
}

func (a *Article) EditResponse() ArticleResponse {
	article := ArticleResponse{
		ID:              a.ID,
		Title:           a.Title,
		CategoryID:      a.CategoryID,
		Desc:            a.Desc,
		Content:         a.Content,
		RenderedContent: a.RenderedContent,
		Keywords:        a.Keywords,
		CreatedAt:       a.CreatedAt,
		ModifiedAt:      a.ModifiedAt,
		PublishedAt:     a.PublishedAt,
		ReproduceURL:    a.ReproduceURL,
		Source:          a.Source,
		Thumb:           a.Thumb,
	}
	serializer := TagsSerializer{Tags: a.Tags}
	article.Tags = serializer.PreviewResponse()

	return article
}

func (s *ArticlesSerializer) Response() []ArticleResponse {
	var articles []ArticleResponse
	for _, article := range s.Articles {
		articles = append(articles, article.Response())
	}
	return articles
}

func (s *ArticlesSerializer) PreviewResponse() []ArticleResponse {
	var articles []ArticleResponse
	for _, article := range s.Articles {
		articles = append(articles, article.PreviewResponse())
	}
	return articles
}
