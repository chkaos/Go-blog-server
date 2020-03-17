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

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{services.NewCategoryService()}
}

// @Summary Get multiple article tags
// @Produce  json
// @Param PageNum query string false "PageNum"
// @Param PageSize query int false "PageSize"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/admin/tags [get]
func (tc *CategoryController) GetCategorys(c *gin.Context) {

	// pageNum, pageSize := utils.GetPaginationParams(c)
	// fmt.Println(pageNum, pageSize)
	var req models.QueryCategoryReq
	err := c.Bind(&req)
	fmt.Println(req)

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
		form validators.AddCategoryForm
		tag  models.Category
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag = form.Transform()

	resp, err := tc.service.AddCategory(&tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_ADD_TAG_FAIL})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

func (tc *CategoryController) UpdateCategory(c *gin.Context) {

	var (
		form validators.EditCategoryForm
		tag  models.Category
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	fmt.Println(form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag = form.Transform()

	resp, err := tc.service.UpdateCategory(&tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_UPDATE_TAG_FAIL})
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
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_DETELE_TAG_FAIL})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
