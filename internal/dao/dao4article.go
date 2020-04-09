package dao

import (
	_ "fmt"

	"github.com/jinzhu/gorm"

	"Go-blog-server/internal/models"
)

type ArticleDAO struct {
	*Dao
}

// NewArticleDAO creates a new ArticleDAO
func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}

// AddArticle  add new Article
func (a *ArticleDAO) AddArticle(Article *models.Article) error {
	return a.db().Create(Article).Error
}

func (a *ArticleDAO) UptadeArticle(Article *models.Article) error {
	return a.db().Model(&models.Article{}).Update(Article).Error
}

// QueryArticles  query all Articles
func (a *ArticleDAO) QueryAllArticles() (Articles []*models.Article, err error) {
	err = a.db().Find(&Articles).Error
	return
}

// QueryArticle query Article by Article name
func (a *ArticleDAO) QueryArticleByTitle(title string) (Article *models.Article, err error) {
	Article = &models.Article{}
	if err = a.db().Where("title = ?", title).First(&Article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// QueryArticle query Article by Article id
func (a *ArticleDAO) QueryArticleByID(id int) (Article *models.Article, err error) {
	Article = &models.Article{}
	if err = a.db().Where("id = ?", id).First(&Article).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
	}
	return
}

// DeleteArticle delete Article
func (a *ArticleDAO) DeleteArticle(id int) error {
	return a.db().Where("id = ?", id).Delete(&models.Article{}).Error
}

func (a *ArticleDAO) QueryArticles(req *models.QueryArticleReq) (total int, Articles []*models.Article, err error) {
	Db := a.db().Preload("Tags").Model(&models.Article{})

	if err = Db.Count(&total).Error; err != nil {
		return
	}

	if req.PageNum > 0 && req.PageSize > 0 {
		Db = Db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	err = Db.Find(&Articles).Error

	return
}
