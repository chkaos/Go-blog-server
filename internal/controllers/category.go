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

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{services.NewCategoryService()}
}

// @Summary 分页获取分类
// @Produce  json
// @Param pageSize query int true "PageSize"
// @Param pageNum query int true "PageNum"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/category [get]
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

// @Summary 获取全部分类
// @Produce  json
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/categorys [get]
func (tc *CategoryController) GetAllCategorys(c *gin.Context) {

	if resp, err := tc.service.QueryAllCategorys(); err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
	} else {
		common.WriteResponseSuccess(c, resp)
	}

}

// @Summary 添加分类
// @Produce  json
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/category [post]
func (tc *CategoryController) AddCategory(c *gin.Context) {
	var form validators.AddCategoryForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	category := form.Transform()

	resp, err := tc.service.AddCategory(category)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

// @Summary 更新分类
// @Produce  json
// @Param id body int true "ID"
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/category [put]
func (tc *CategoryController) UpdateCategory(c *gin.Context) {

	var form validators.EditCategoryForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	category := form.Transform()

	resp, err := tc.service.UpdateCategory(category)

	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 删除分类
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/category/{id} [delete]
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
