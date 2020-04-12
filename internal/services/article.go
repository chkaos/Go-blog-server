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
	return &ArticleService{dao: dao.NewArticleDAO()}
}

// QueryArticles
func (as *ArticleService) QueryArticlesReq(req *models.QueryArticleReq) (resp common.Response, err error) {
	var (
		total    int
		articles []models.Article
	)

	if articles, total, err = as.dao.QueryArticles(req); err != nil {
		resp.Err = common.ErrorGetArticleFail
		return
	}

	ArticlesSerializer := models.ArticlesSerializer{articles}

	rep := &models.PaginationRep{
		Total:    total,
		PageSize: req.PageSize,
		PageNum:  req.PageNum,
		List:     ArticlesSerializer.PreviewResponse(),
	}

	resp = common.Response{Err: common.SUCCESS, Data: rep}

	return
}

// AddArticle
func (as *ArticleService) AddArticle(Article models.Article) (resp common.Response, err error) {

	var (
		articleModel models.Article
	)
	articleModel, err = as.dao.QueryArticleByTitle(Article.Title)

	if articleModel.ID > 0 {
		resp.Err = common.ErrorArticleExist
		return
	}

	if err = as.dao.AddArticle(Article); err != nil {
		resp.Err = common.ErrorAddArticleFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// UpdateArticle
func (as *ArticleService) UpdateArticle(Article models.Article) (resp common.Response, err error) {

	var (
		articleModel models.Article
	)
	articleModel, err = as.dao.QueryArticleByID(Article.ID)

	if articleModel.ID == 0 {
		resp.Err = common.ErrorArticleNotExist
		return
	}

	if err = as.dao.UpdateArticle(Article); err != nil {
		resp.Err = common.ErrorUpdateArticleFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// UpdateArticle
func (as *ArticleService) UpdateArticleState(id, state int) (resp common.Response, err error) {

	if err = as.dao.UpdateArticleState(id, state); err != nil {
		resp.Err = common.ErrorUpdateArticleFail
	} else {
		resp.Err = common.SUCCESS
	}

	return
}

// GetArticle
func (as *ArticleService) GetArticle(id int) (resp common.Response, err error) {

	var (
		articleModel models.Article
	)
	articleModel, err = as.dao.QueryArticleByID(id)

	if articleModel.ID == 0 {
		resp.Err = common.ErrorArticleNotExist
		return
	}

	resp.Data = articleModel.EditResponse()
	resp.Err = common.SUCCESS

	return
}

// DeleteArticle
func (as *ArticleService) DeleteArticle(id int) (resp common.Response, err error) {

	var (
		articleModel models.Article
	)
	articleModel, err = as.dao.QueryArticleByID(id)

	if articleModel.ID == 0 {
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
