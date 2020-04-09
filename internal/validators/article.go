package validators

import (
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
)

type AddArticleForm struct {
	CategoryID   int    `form:"category_id" valid:"Required;Min(1)"`
	Title        string `form:"title" valid:"Required;MaxSize(100)"`
	Desc         string `form:"desc" valid:"Required;MaxSize(100)"`
	Keywords     string `form:"keywords" valid:"MaxSize(100)"`
	Content      string `form:"content" valid:"Required;MinSize(1)"`
	Source       int    `form:"source" valid:"Required;Range(0,3)"`
	ReproduceURL string `form:"reproduce_url" valid:"MaxSize(100)"`
	Thumb        string `form:"thumb" valid:"Required;MinSize(1)"`
	State        int    `form:"state" valid:"Required;Range(0,1)"`
	Tags         []int  `form:"tags" valid:"Required"`
}

type EditArticleForm struct {
	ID           int    `form:"id" valid:"Required;Min(1)"`
	CategoryID   int    `form:"category_id" valid:"Required;Min(1)"`
	Title        string `form:"title" valid:"Required;MaxSize(100)"`
	Desc         string `form:"desc" valid:"Required;MaxSize(100)"`
	Keywords     string `form:"keywords" valid:"MaxSize(100)"`
	Content      string `form:"content" valid:"Required;MinSize(1)"`
	Source       int    `form:"source" valid:"Required;Range(0,3)"`
	ReproduceURL string `form:"reproduce_url" valid:"MaxSize(100)""`
	Thumb        string `form:"thumb" valid:"Required;MinSize(1)"`
	Tags         []int  `form:"tags" valid:"Required"`
	State        int    `form:"state" valid:"Required;Range(0,1)"`
}

func (a *AddArticleForm) Transform() (articleModel models.Article, err error ) {
	var (
		categoryDao dao.CategoryDao
		tagDao dao.TagDao
		categoryModel models.Category
		tagModels []models.Tag
	)

	articleModel = models.Article{
		Title:        a.Title,
		Desc:         a.Desc,
		Keywords:     a.Keywords,
		Content:      a.Content,
		Source:       a.Source,
		ReproduceURL: a.ReproduceURL,
		Thumb:        a.Thumb,
		State:        a.State,
	}

	if categoryModel, err = categoryDao.QueryCategoryByID(a.CategoryID); err != nil {
		return
	}

	articleModel.Category = categoryModel

	if tagModels, err = tagDao.SetTags(a.Tags); err != nil {
		return
	}

	articleModel.Tags = tagModels

	return
}

func (e *EditArticleForm) Transform() (articleModel models.Article err error) {
	var (
		categoryDao dao.CategoryDao
		tagDao dao.TagDao
		categoryModel models.Category
		tagModels []models.Tag
	)

	articleModel = models.Article{
		ID:           e.ID,
		Title:        e.Title,
		Desc:         e.Desc,
		Keywords:     e.Keywords,
		Content:      e.Content,
		Source:       e.Source,
		ReproduceURL: e.ReproduceURL,
		Thumb:        e.Thumb,
		State:        e.State,
	}

	if categoryModel, err = categoryDao.QueryCategoryByID(e.CategoryID); err != nil {
		return
	}

	articleModel.Category = categoryModel

	if tagModels, err = tagDao.SetTags(e.Tags); err != nil {
		return
	}

	articleModel.Tags = tagModels

	return
}
