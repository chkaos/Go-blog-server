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

type TagController struct {
	service *services.TagService
}

func NewTagController() *TagController {
	return &TagController{services.NewTagService()}
}

// @Summary 分页获取标签
// @Produce  json
// @Param pageSize query int true "PageSize"
// @Param pageNum query int true "PageNum"
// @Param name query string "Name"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/tag [get]
func (tc *TagController) GetTags(c *gin.Context) {

	var req models.QueryTagReq
	err := c.Bind(&req)

	resp, err := tc.service.QueryTagsReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 获取全部标签
// @Produce  json
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/tags [get]
func (tc *TagController) GetAllTags(c *gin.Context) {

	if resp, err := tc.service.QueryAllTags(); err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
	} else {
		common.WriteResponseSuccess(c, resp)
	}

}

// @Summary 添加标签
// @Produce  json
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Param icon body string "Icon"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/tag [post]
func (tc *TagController) AddTag(c *gin.Context) {
	var form validators.AddTagForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag := form.Transform()

	resp, err := tc.service.AddTag(tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

// @Summary 更新标签
// @Produce  json
// @Param id body int true "ID"
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Param icon body string "Icon"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/tag [put]
func (tc *TagController) UpdateTag(c *gin.Context) {

	var form validators.EditTagForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag := form.Transform()

	resp, err := tc.service.UpdateTag(tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/tag/{id} [delete]
func (tc *TagController) DeleteTag(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := tc.service.DeleteTag(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorDeleteTagFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
