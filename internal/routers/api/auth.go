package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/models"
	"Go-blog-server/internal/service/authSrv"
	"Go-blog-server/pkg/e"
	"Go-blog-server/pkg/helper"
	"Go-blog-server/pkg/utils"
)

// @Summary Get Auth
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/auth [post]
func GetAuth(c *gin.Context) {
	appG := helper.Gin{C: c}
	valid := validation.Validation{}
	
	var auth models.UserModel
	c.BindJSON(&auth)
	code := e.INVALID_PARAMS
	username := auth.Username
	password := auth.Password

	fmt.Println(auth)
	valid.Required(auth.Username, "username")
	valid.MinSize(auth.Password, 6, "password").Message("密码长度不得小于%d", 6)

	ok, _ := valid.Valid(auth)
	if !ok {
		utils.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	user, isExistError := authSrv.CheckAuth(username, password)
	if(isExistError != nil){
		appG.Response(http.StatusBadRequest, e.ERROR_AUTH, nil)
		return 
	}
	
	token, err := utils.GenerateToken(int(user.ID), username, user.Role)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	} else {
		code = e.SUCCESS
	}

	appG.Response(http.StatusOK, code, user.UserResponseWithToken(token))
}
