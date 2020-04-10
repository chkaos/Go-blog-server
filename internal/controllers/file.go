package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/models"
	"Go-blog-server/internal/services"
)

type FileController struct {
	service *services.FileService
}

func NewFileController() *FileController {
	return &FileController{services.NewFileService()}
}

// @Summary 上传文件
// @Produce  json
// @Param file formData file true "File"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/file/upload [post]
func (fc *FileController) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		common.ResponseWithValidation(c)
		return
	}

	resp, err := fc.service.UploadImg(header, file)

	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)

}

// @Summary 分页获取文件资源
// @Produce  json
// @Param pageSize query int true "PageSize"
// @Param pageNum query int true "PageNum"
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /api/admin/file [get]
func (tc *FileController) GetFiles(c *gin.Context) {

	var req models.QueryFileReq
	err := c.Bind(&req)

	resp, err := tc.service.QueryFilesReq(&req)
	if err != nil {
		common.WriteResponse(c, http.StatusBadRequest, resp)
		return
	}

	common.WriteResponseSuccess(c, resp)
}
