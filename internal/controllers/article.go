package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/models"
	"Go-blog-server/internal/services"
	"Go-blog-server/internal/validators"
	"Go-blog-server/pkg/e"
)

type ArticleController struct {
	service *services.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{services.NewArticleService()}
}

func (ac *ArticleController) GetArticle(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := ac.service.GetArticle(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorGetArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (ac *ArticleController) GetArticles(c *gin.Context) {

	var req models.QueryArticleReq
	err := c.Bind(&req)

	resp, err := ac.service.QueryArticlesReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (ac *ArticleController) AddArticle(c *gin.Context) {
	var form validators.AddArticleForm

	if httpCode, err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: err})
		return
	}

	if article, err := form.Transform(); err != nil {
		common.ResponseWithValidation()
	}

	if resp, err := ac.service.AddArticle(&article); err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

func (ac *ArticleController) UpdateArticle(c *gin.Context) {

	var form  validators.EditArticleForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	if article, err := form.Transform(); err != nil {
		common.ResponseWithValidation()
		return
	}

	if resp, err := ac.service.UpdateArticle(&article); err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := ac.service.DeleteArticle(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorDeleteArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
