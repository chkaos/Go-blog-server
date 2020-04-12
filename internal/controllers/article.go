package controllers

import (
	"net/http"

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

// @Summary 获取文章详情
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/article/{id} [get]
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

// @Summary 分页获取文章
// @Produce  json
// @Param pageSize query int true "PageSize"
// @Param pageNum query int true "PageNum"
// @Param tag query string "Tags"
// @Param category query int "Category"
// @Param state query int "Category"
// @Param source query int "Category"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/articles [get]
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

// @Summary 添加文章
// @Produce  json
// @Param tagIds body int true "TagIDs"
// @Param category body int true "Category"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param source body int true "Source"
// @Param reproduceUrl body string "ReproduceURL"
// @Param thumb body string true "Thumb"
// @Param state body int true "State"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/article [post]
func (ac *ArticleController) AddArticle(c *gin.Context) {
	var form validators.AddArticleForm

	if httpCode, err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: err})
		return
	}

	article, err := form.Transform()
	if err != nil {
		common.ResponseWithValidation(c)
		return
	}

	resp, err := ac.service.AddArticle(article)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

// @Summary 更新文章
// @Produce  json
// @Param id body int true "ID"
// @Param tagIds body int true "TagIDs"
// @Param category body int true "Category"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param source body int true "Source"
// @Param reproduceUrl body string "ReproduceURL"
// @Param thumb body string true "Thumb"
// @Param state body int true "State"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/article [put]
func (ac *ArticleController) UpdateArticle(c *gin.Context) {

	var form validators.EditArticleForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	article, err := form.Transform()
	if err != nil {
		common.ResponseWithValidation(c)
		return
	}

	resp, err := ac.service.UpdateArticle(article)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateArticleFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/article/{id} [delete]
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
