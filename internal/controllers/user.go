package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/models"
	"Go-blog-server/internal/services"
	"Go-blog-server/pkg/e"
	"Go-blog-server/pkg/utils"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{services.NewUserService()}
}

// @Summary Get Auth
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/admin/auth [post]
func (uc *UserController) Auth(c *gin.Context) {
	valid := validation.Validation{}

	var auth models.User
	err := c.BindJSON(&auth)
	fmt.Println(err)
	username := auth.Username
	password := auth.Password

	fmt.Println(auth)
	valid.Required(auth.Username, "username")
	valid.MinSize(auth.Password, 6, "password").Message("密码长度不得小于%d", 6)

	ok, _ := valid.Valid(auth)
	if !ok {
		utils.MarkErrors(valid.Errors)
		common.ResponseWithValidation(c)
		return
	}

	user, isExistError := uc.service.Auth(username, password)
	fmt.Println(isExistError)
	if isExistError != nil {
		common.Response(c, http.StatusBadRequest, e.ERROR_AUTH, nil)
		return
	}

	token, err := utils.GenerateToken(int(user.ID), username, user.Role)
	if err != nil {
		common.Response(c, http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	common.ResponseSuccess(c, user.ResponseWithToken(token))
}
