package services

import (
	"Go-blog-server/internal/common"
	"Go-blog-server/internal/dao"
	"Go-blog-server/internal/models"
)

type ArticleService struct {
	dao *dao.ArticleDAO
}

func NewArticleService() *ArticleService {
	return &ArticleService{dao: new(dao.ArticleDAO)}
}

// QueryArticles
func (as *ArticleService) QueryArticlesReq(req *models.QueryArticleReq) (resp common.Response, err error) {
	var (
		total    int
		Articles []*models.Article
	)

	if total, Articles, err = as.dao.QueryArticles(req); err != nil {
		resp.Err = common.ErrorGetArticleFail
		return
	}

	ArticlesSerializer := models.ArticlesSerializer{Articles}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     ArticlesSerializer.Response(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

func (as *ArticleService) AddArticle(Article *models.Article) (resp common.Response, err error) {

	var (
		ArticleModel *models.Article
	)
	ArticleModel, err = as.dao.QueryArticleByTitle(Article.Title)

	if ArticleModel.ID > 0 {

		resp.Err = common.ErrorArticleExist
		return
	}

	if err = as.dao.AddArticle(Article); err != nil {
		resp.Err = common.ErrorAddArticleFail
	} else {
		resp.Err = common.SUCCESS
		resp.Data = Article.Response()
	}

	return
}

func (as *ArticleService) UpdateArticle(Article *models.Article) (resp common.Response, err error) {

	var (
		ArticleModel *models.Article
	)
	ArticleModel, err = as.dao.QueryArticleByID(Article.ID)

	if ArticleModel.ID == 0 {

		resp.Err = common.ErrorArticleNotExist
		return
	}

	if err = as.dao.UptadeArticle(Article); err != nil {
		resp.Err = common.ErrorUpdateArticleFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

func (as *ArticleService) GetArticle(id int) (resp common.Response, err error) {

	var (
		ArticleModel *models.Article
	)
	ArticleModel, err = as.dao.QueryArticleByID(id)

	if ArticleModel.ID == 0 {

		resp.Err = common.ErrorArticleNotExist
		return
	}

	resp.Data = ArticleModel.EditResponse()
	resp.Err = common.SUCCESS

	return
}

func (as *ArticleService) DeleteArticle(id int) (resp common.Response, err error) {

	var (
		ArticleModel *models.Article
	)
	ArticleModel, err = as.dao.QueryArticleByID(id)

	if ArticleModel.ID == 0 {

		resp.Err = common.ErrorArticleNotExist
		return
	}

	if err = as.dao.DeleteArticle(id); err != nil {
		resp.Err = common.ErrorDeleteArticleFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}
