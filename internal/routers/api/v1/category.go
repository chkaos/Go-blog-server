package v1

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"Go-blog-server/internal/models"
	"Go-blog-server/pkg/e"
)

// @Summary Get categorys
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/category [get]

func GetCategorys(c *gin.Context) {
	fmt.Println(c)
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetCategorys(maps)
	data["total"] = models.GetCategorysTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// @Summary Add category
// @Produce  json
// @Param name body string true "Name"
// @Param desc body string true "Desc"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/category [post]

func AddCategory(c *gin.Context) {
	name := c.Query("name")
	desc := c.Query("desc")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistCategoryByName(name) {
			code = e.SUCCESS
			models.AddCategory(name, desc)
		} else {
			code = e.ERROR_NOT_EXIST_CATEGORY
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary Edit category
// @Produce  json
// @Param id path int true "ID"
// @Param name body string false "Name"
// @Param desc body string false "Desc"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/category [put]
func EditCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")

	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCategoryByID(id) {
			data := make(map[string]interface{})
			if name != "" {
				data["name"] = name
			}

			models.EditCategory(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_CATEGORY
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary Delete category
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/category [delete]
func DeleteCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	code := e.INVALID_PARAMS

	if models.ExistCategoryByID(id) {
		code = e.SUCCESS
		models.DeleteCategory(id)
	} else {
		code = e.ERROR_NOT_EXIST_CATEGORY
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
