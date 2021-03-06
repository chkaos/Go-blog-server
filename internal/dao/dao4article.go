package dao

import (
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
	if tx.Model(&article).Update(&article); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if tx.Model(&article).Association("Tag").Replace(article.Tags); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	err := tx.Commit().Error

	return err
}

// UpdateArticleState update Article State Only
func (a *ArticleDAO) UpdateArticleState(id, state int) error {
	return a.db().Model(&models.Article{}).Where("id = ?", id).Update("state", state).Error
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

	if tx.Where("id = ?", id).Delete(&model); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	if tx.Exec("DELETE FROM tag_relation WHERE article_id = ?", id); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	err := tx.Commit().Error

	return err
}

// QueryArticles query Articles by condition
func (a *ArticleDAO) QueryArticles(req *models.QueryArticleReq) (articles []models.Article, total int, err error) {
	db := a.db()

	if req.Category > 0 {
		db = db.Where("category_id = ?", req.Category)
	}

	if req.State >= 0 {
		db = db.Where("state = ?", req.State)
	}

	if req.Source >= 0 {
		db = db.Where("source = ?", req.Source)
	}

	db.Model(&articles).Count(&total)

	tx := db.Order("created_at desc").Begin()
	if req.Tag > 0 {
		var tagModel models.Tag
		tx.Where("id = ?", req.Tag).First(&tagModel)
		if tagModel.ID != 0 {
			tx.Model(&tagModel).Offset((req.PageNum-1)*req.PageSize).Limit(req.PageSize).Related(&articles, "Articles")
			total = tx.Model(&tagModel).Association("Articles").Count()
		}
	} else {
		tx.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&articles)
	}

	for i, _ := range articles {
		tx.Model(&articles[i]).Related(&articles[i].Category, "Category")
		tx.Model(&articles[i]).Related(&articles[i].Tags, "Tags")
	}
	err = tx.Commit().Error

	return articles, total, err
}
