package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Err
	Data interface{} `json:"data,omitempty"`
}

// Response setting gin.JSON
func WriteResponse(c *gin.Context, httpCode int, resp Response) {
	c.JSON(httpCode, resp)
}

// func ResponseWithPanic(c *gin.Context, err error) {
// 	c.JSON(http.StatusBadGateway, gin.H{
// 		"success": false,
// 	})
// }

func ResponseWithValidation(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		Err: ErrorInvalidParams,
	})
}

func WriteResponseSuccess(c *gin.Context, resp Response) {
	c.JSON(http.StatusOK, resp)
}

// func WriteJSON(c *gin.Context, data interface{}, err error) {
// 	httpcode := http.StatusOK
// 	bcode := Cause(err)

// 	c.JSON(httpcode, Response{
// 		Code: bcode.Code(),
// 		Msg:  bcode.Message(),
// 		Data: data,
// 	})
// }
