package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/e"
	"Go-blog-server/pkg/logging"
	"Go-blog-server/pkg/utils"
)

type User struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/auth [post]
func GetAuth(c *gin.Context) {
	valid := validation.Validation{}
	var auth User
	c.BindJSON(&auth)

	ok, _ := valid.Valid(auth)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	username := auth.Username
	password := auth.Password

	if ok {
		user, isExistError := models.CheckAuth(username, password)
		fmt.Println(user, isExistError)
		if isExistError == nil {
			token, err := utils.GenerateToken(username, password, user.Role)
			fmt.Println(err)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
