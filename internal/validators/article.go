package validators

import (
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
)

type AddArticleForm struct {
	CategoryID   int    `json:"category_id" form:"category_id" valid:"Required;Min(1)"`
	Title        string `form:"title" valid:"Required;MaxSize(100)"`
	Desc         string `form:"desc" valid:"Required;MaxSize(100)"`
	Keywords     string `form:"keywords" valid:"MaxSize(100)"`
	Content      string `form:"content" valid:"Required"`
	Source       int    `form:"source" valid:"Range(0,3)"`
	ReproduceURL string `json:"reproduce_url" form:"reproduce_url" valid:"MaxSize(100)"`
	Thumb        string `form:"thumb" valid:"Required"`
	State        int    `form:"state" valid:"Range(0,1)"`
	Tags         []int  `form:"tags" valid:"Required"`
}

type EditArticleForm struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	CategoryID   int    `json:"category_id" form:"category_id" valid:"Required;Min(1)"`
	Title        string `form:"title" valid:"Required;MaxSize(100)"`
	Desc         string `form:"desc" valid:"Required;MaxSize(100)"`
	Keywords     string `form:"keywords" valid:"MaxSize(100)"`
	Content      string `form:"content" valid:"Required"`
	Source       int    `form:"source" valid:"Range(0,3)"`
	ReproduceURL string `json:"reproduce_url" form:"reproduce_url" valid:"MaxSize(100)"`
	Thumb        string `form:"thumb" valid:"Required"`
	State        int    `form:"state" valid:"Range(0,1)"`
	Tags         []int  `form:"tags" valid:"Required"`
}

func (a *AddArticleForm) Transform() (articleModel models.Article, err error) {
	var (
		tagModels []models.Tag
	)

	tagDao := dao.NewTagDAO()

	articleModel = models.Article{
		CategoryID:   a.CategoryID,
		Title:        a.Title,
		Desc:         a.Desc,
		Keywords:     a.Keywords,
		Content:      a.Content,
		Source:       a.Source,
		ReproduceURL: a.ReproduceURL,
		Thumb:        a.Thumb,
		State:        a.State,
	}

	if tagModels, err = tagDao.SetTags(a.Tags); err != nil {
		return
	}

	articleModel.Tags = tagModels

	return
}

func (e *EditArticleForm) Transform() (articleModel models.Article, err error) {
	var (
		tagModels []models.Tag
	)

	tagDao := dao.NewTagDAO()

	articleModel = models.Article{
		CategoryID:   e.CategoryID,
		Title:        e.Title,
		Desc:         e.Desc,
		Keywords:     e.Keywords,
		Content:      e.Content,
		Source:       e.Source,
		ReproduceURL: e.ReproduceURL,
		Thumb:        e.Thumb,
		State:        e.State,
	}
	articleModel.ID = e.ID

	if tagModels, err = tagDao.SetTags(e.Tags); err != nil {
		return
	}

	articleModel.Tags = tagModels

	return
}
