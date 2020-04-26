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

type BulletinController struct {
	service *services.BulletinService
}

func NewBulletinController() *BulletinController {
	return &BulletinController{services.NewBulletinService()}
}

// @Summary 分页获取公告
// @Produce  json
// @Param pageSize query int true "PageSize"
// @Param pageNum query int true "PageNum"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/bulletin [get]
func (tc *BulletinController) GetBulletins(c *gin.Context) {

	var req models.QueryBulletinReq
	err := c.Bind(&req)

	resp, err := tc.service.QueryBulletinsReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 添加公告
// @Produce  json
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/bulletin [post]
func (tc *BulletinController) AddBulletin(c *gin.Context) {
	var form validators.AddBulletinForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	bulletin := form.Transform()

	resp, err := tc.service.AddBulletin(bulletin)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAddBulletinFail})
		return
	}

	common.WriteResponseSuccess(c, resp)

}

// @Summary 更新公告
// @Produce  json
// @Param id body int true "ID"
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/bulletin [put]
func (tc *BulletinController) UpdateBulletin(c *gin.Context) {

	var form validators.EditBulletinForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	bulletin := form.Transform()

	resp, err := tc.service.UpdateBulletin(bulletin)

	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorUpdateBulletinFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}

// @Summary 删除公告
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/bulletin/{id} [delete]
func (tc *BulletinController) DeleteBulletin(c *gin.Context) {

	httpCode, Err, id := validators.BindID(c)

	if httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	resp, err := tc.service.DeleteBulletin(id)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorDeleteBulletinFail})
		return
	}

	common.WriteResponseSuccess(c, resp)
}
