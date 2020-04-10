package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
	"Go-blog-server/internal/services"
	"Go-blog-server/internal/validators"
	"Go-blog-server/pkg/e"
	"Go-blog-server/pkg/utils"
)

type UserController struct {
	service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{services.NewUserService()}
}

// @Summary 登录
// @Produce  json
// @Param username body string true "username"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/admin/auth [post]
func (uc *UserController) Auth(c *gin.Context) {

	var form validators.AuthForm

	if httpCode, Err := validators.BindAndValid(c, &form); httpCode != e.SUCCESS {
		common.WriteResponse(c, httpCode, common.Response{Err: Err})
		return
	}

	user, isExistError := uc.service.Auth(form.Username, form.Password)

	if isExistError != nil {
		common.WriteResponse(c, http.StatusBadRequest, common.Response{
			Err: common.ErrAuth,
		})
		return
	}

	token, err := utils.GenerateToken(int(user.ID), form.Username, user.Role)
	if err != nil {
		common.WriteResponse(c, http.StatusInternalServerError, common.Response{Err: common.ErrorAuthToken})
		return
	}

	common.WriteResponseSuccess(c, common.Response{Err: common.SUCCESS, Data: user.ResponseWithToken(token)})
}
