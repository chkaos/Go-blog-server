package controllers

import (
	"net/http"

	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/models"
	"Go-blog-server/internal/services"
	"Go-blog-server/internal/validators"
	"Go-blog-server/pkg/e"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{services.NewCategoryService()}
}

func (tc *CategoryController) GetCategorys(c *gin.Context) {

	var req models.QueryCategoryReq
	err := c.Bind(&req)

	resp, err := tc.service.QueryCategorysReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (tc *CategoryController) GetAllCategorys(c *gin.Context) {

	if resp, err := tc.service.QueryAllCategorys(); err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
	} else {
		common.WriteResponseSuccess(c, resp)
	}

}

func (tc *CategoryController) AddCategory(c *gin.Context) {
	var (
		form     validators.AddCategoryForm
		category models.Category
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	category = form.Transform()

	resp, err := tc.service.AddCategory(&category)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

func (tc *CategoryController) UpdateCategory(c *gin.Context) {

	var (
		form     validators.EditCategoryForm
		category models.Category
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	category = form.Transform()

	resp, err := tc.service.UpdateCategory(&category)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (tc *CategoryController) DeleteCategory(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := tc.service.DeleteCategory(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorDeleteTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
