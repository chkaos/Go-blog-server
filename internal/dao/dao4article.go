package dao

import (
	_ "fmt"

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
func (a *ArticleDAO) AddArticle(article models.Article) error {

	article.HandlePublishedAt()
	return a.db().Create(&article).Error
}

// UpdateArticle update Article
func (a *ArticleDAO) UpdateArticle(article models.Article) error {

	article.HandlePublishedAt()

	db := a.db()
	tx := db.Begin()
	tx.Model(&article).Update(&article)
	tx.Model(&article).Association("Tag").Replace(article.Tags)
	err := tx.Commit().Error

	return err
}

// QueryArticle query Article by Article title
func (a *ArticleDAO) QueryArticleByTitle(title string) (article models.Article, err error) {
	article, err = a.FindOneArticle("title = ?", title)
	return
}

// QueryArticle query an Article by Article id
func (a *ArticleDAO) QueryArticleByID(id int) (article models.Article, err error) {
	article, err = a.FindOneArticle("id = ?", id)
	return
}

// FindOneArticle query an Article by condition
func (a *ArticleDAO) FindOneArticle(condition ...interface{}) (model models.Article, err error) {
	db := a.db()
	tx := db.Begin()
	tx.Where(condition).First(&model)
	tx.Model(&model).Related(&model.Category, "CategoryID")
	tx.Model(&model).Related(&model.Tags, "Tags")
	err = tx.Commit().Error
	return
}

// DeleteArticle delete an Article by id
func (a *ArticleDAO) DeleteArticle(id int) error {
	var model models.Article

	db := a.db()
	tx := db.Begin()
	tx.Where("id = ?", id).Delete(&model)
	tx.Model(&model).Association("Tag").Clear()
	err := tx.Commit().Error
	return err
}

// QueryArticles query Articles by condition
func (a *ArticleDAO) QueryArticles(req *models.QueryArticleReq) (articles []models.Article, total int, err error) {
	db := a.db()
	db.Model(&articles).Count(&total)
	if req.PageNum > 0 && req.PageSize > 0 {
		db = db.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize)
	}

	tx := db.Begin()
	if req.Tag > 0 {
		var tagModel models.Tag
		tx.Where("id = ?", req.Tag).First(&tagModel)
		if tagModel.ID != 0 {
			tx.Model(&tagModel).Related(&articles, "Articles")
			total = tx.Model(&tagModel).Association("Articles").Count()
		}
	}

	if req.Category > 0 {
		var categoryModel models.Category
		tx.Where("id = ?", req.Category).First(&categoryModel)
		if categoryModel.ID != 0 {
			tx.Model(&articles).Where("category_id = ?", categoryModel.ID)
			total = tx.Model(&categoryModel).Association("Articles").Count()
		}
	}

	for i, _ := range articles {
		tx.Model(&articles[i]).Related(&articles[i].Tags, "Tags")
	}
	err = tx.Commit().Error
	return articles, total, err
}
