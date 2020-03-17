package controllers

import (
	"net/http"

	_ "github.com/astaxie/beego/validation"
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
