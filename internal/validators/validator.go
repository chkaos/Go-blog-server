package validators

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/pkg/logging"
)

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

// Bind Query ID and check
func BindID(c *gin.Context) (httpcode int, Err common.Err, id int) {
	id, err := strconv.Atoi(c.Param("id"))

	fmt.Println(id, err)

	httpcode = http.StatusOK
	Err = common.SUCCESS

	if err != nil || id <= 0 {
		httpcode = http.StatusBadRequest
		Err = common.ERROR_INVALID_PAMAMS
	}

	return
}

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		fmt.Println(err)
		logging.Info(err.Key, err.Message)
	}

	return
}
