package api

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "Go-blog-server/pkg/e"
    "Go-blog-server/pkg/utils"
    "Go-blog-server/internal/models"
    "Go-blog-server/pkg/logging"
)

type user struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username body string true "userName"
// @Param password body string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/auth [post]
func GetAuth(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    valid := validation.Validation{}
    a := user{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS
    if ok {
        user, isExist := models.CheckAuth(username, password)
        if isExist {
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
            logging.Info(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}