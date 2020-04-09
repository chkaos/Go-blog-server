package validators

import "Go-blog-server/internal/models"

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

func (a *AddArticleForm) Transform() (article models.Article) {
	return models.Article{
		Desc: a.Desc,
	}
}

func (e *EditArticleForm) Transform() (article models.Article) {
	article = models.Article{
		Desc: e.Desc,
	}
	article.ID = e.ID
	return
}
