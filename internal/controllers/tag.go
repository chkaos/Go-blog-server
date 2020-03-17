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

type TagController struct {
	service *services.TagService
}

func NewTagController() *TagController {
	return &TagController{services.NewTagService()}
}

// @Summary Get multiple article tags
// @Produce  json
// @Param PageNum query string false "PageNum"
// @Param PageSize query int false "PageSize"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/admin/tags [get]
func (tc *TagController) GetTags(c *gin.Context) {

	var req models.QueryTagReq
	err := c.Bind(&req)
	fmt.Println(req)

	resp, err := tc.service.QueryTagsReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (tc *TagController) GetAllTags(c *gin.Context) {

	if resp, err := tc.service.QueryAllTags(); err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
	} else {
		common.WriteResponseSuccess(c, resp)
	}

}

func (tc *TagController) AddTag(c *gin.Context) {
	var (
		form validators.AddTagForm
		tag  models.Tag
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag = form.Transform()

	resp, err := tc.service.AddTag(&tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_ADD_TAG_FAIL})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

func (tc *TagController) UpdateTag(c *gin.Context) {

	var (
		form validators.EditTagForm
		tag  models.Tag
	)

	httpCode, Err := validators.BindAndValid(c, &form)

	fmt.Println(form)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	tag = form.Transform()

	resp, err := tc.service.UpdateTag(&tag)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_UPDATE_TAG_FAIL})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

func (tc *TagController) DeleteTag(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := tc.service.DeleteTag(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ERROR_DETELE_TAG_FAIL})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
