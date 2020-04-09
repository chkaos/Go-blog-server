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

// AddArticle add new Article
func (a *ArticleDAO) AddArticle(Article *models.Article) error {
	return a.db().Create(Article).Error
}

// UpdateArticle update Article
func (a *ArticleDAO) UpdateArticle(Article *models.Article) error {
	db := a.db()
	tx := db.Begin()
	tx.Model(&Article).Update(Article)
	tx.Model(&Article).Association("Tag").Replace(Article.Tags)
	err := tx.Commit().Error
	return err
}

// QueryArticle query Article by Article title
func (a *ArticleDAO) QueryArticleByTitle(title string) (Article *models.Article, err error) {
	Article, err = FindOneArticle(struct {
		title string
	}{
		title: title,
	})
	return
}

// QueryArticle query an Article by Article id
func (a *ArticleDAO) QueryArticleByID(id int) (Article *models.Article, err error) {
	Article, err = FindOneArticle(struct {
		id int
	}{
		id: id,
	})
	return
}

// FindOneArticle query an Article by condition
func (a *ArticleDAO) FindOneArticle(condition interface{}) (Article *models.Article, err error) {
	db := a.db()
	var model models.Article
	tx := db.Begin()
	tx.Where(condition).First(&model)
	tx.Model(&model).Related(&model.Category, "Category")
	tx.Model(&model).Related(&model.Tags, "Tags")
	err := tx.Commit().Error
	return model, err
}

// DeleteArticle delete an Article by id
func (a *ArticleDAO) DeleteArticle(id int) error {
	var model models.Article

	db := a.db()
	tx := db.Begin()
	tx.Where("id = ?", id).Delete(&model)
	tx.Model(&models).Association("Tag").Clear()
	err := tx.Commit().Error
	return err
}

// QueryArticles query Articles by condition
func (a *ArticleDAO) QueryArticles(req *models.QueryArticleReq) (models []*models.Article, total int, err error) {
	db := a.db()
	db.Model(&models).Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	tx := db.Begin()
	if req.Tag > 0 {
		var tagModel models.Tag
		tx.Where("id = ?", req.Tag).First(&tagModel)
		if tagModel.ID != 0 {
			tx.Model(&tagModel).Related(&models, "Articles")
			total = tx.Model(&tagModel).Association("Articles").Count()
		}
	}

	if req.Category > 0 {
		var categoryModel models.Category
		tx.Where("id = ?", req.Category).First(&categoryModel)
		if categoryModel.ID != 0 {
			tx.Model(&models).Where("category_id = ?", categoryModel.ID)
			total = tx.Where("category_id = ?", categoryModel.ID).Count()
		}
	}

	for i, _ := range models {
		tx.Model(&models[i]).Related(&models[i].Tags, "Tags")
	}
	err = tx.Commit().Error
	return models, total, err
}
