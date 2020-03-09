package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"Go-blog-server/pkg/e"
)

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response setting gin.JSON
func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, BaseResponse{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}

func ResponseWithPanic(c *gin.Context, err error) {
	c.JSON(http.StatusBadGateway, gin.H{
		"success": false,
	})
}

func ResponseWithValidation(c *gin.Context) {
	errCode := e.INVALID_PARAMS
	c.JSON(http.StatusBadRequest, BaseResponse{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	errCode := e.SUCCESS
	c.JSON(http.StatusOK, BaseResponse{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}