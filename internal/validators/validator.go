package validators

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/pkg/logging"
)

type IDForm struct {
	ID int `form:"id" valid:"Required;Min(1)" json:"id"`
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (httpcode int, Err common.Err) {
	err := c.Bind(form)

	if err != nil {
		return http.StatusBadRequest, common.ERROR_INVALID_PAMAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		return http.StatusInternalServerError, common.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, common.ERROR_INVALID_PAMAMS
	}

	return http.StatusOK, common.SUCCESS
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Println(err)
		logging.Info(err.Key, err.Message)
	}

	return
}
