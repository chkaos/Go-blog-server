package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"Go-blog-server/internal/common"
		"Go-blog-server/internal/validators"
	"Go-blog-server/internal/models"
	"Go-blog-server/internal/services"
	"Go-blog-server/pkg/e"
)

type TagController struct {
	service *services.TagService
}

func NewTagController() *TagController {
	return &TagController{services.NewTagService()}
}

// @Summary Get multiple article tags
// @Produce  json
// @Param PageNum query string false "PageNum"
// @Param PageSize query int false "PageSize"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/admin/tags [get]
func (tc *TagController)  GetTags(c *gin.Context) {

	// pageNum, pageSize := utils.GetPaginationParams(c)
	// fmt.Println(pageNum, pageSize)
	var req models.QueryTagReq
	err := c.Bind(&req)
	fmt.Println(req)

	res, err := tc.service.QueryTagsReq(&req)
	if err != nil {
		common.Response(c, http.StatusBadRequest, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	common.ResponseSuccess(c, res)
}

func (tc *TagController) GetAllTags(c *gin.Context) {

	tags, err := tc.service.QueryAllTags()

	if err != nil {
		common.Response(c, http.StatusBadRequest, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	common.ResponseSuccess(c, tags)
}



func (tc *TagController) AddTag(c *gin.Context) {
		var (
			form validators.AddTagForm
			tag models.Tag
		)

		httpCode, errCode := validators.BindAndValid(c, &form)

		if errCode != e.SUCCESS {
			common.Response(c, httpCode, errCode, nil)
			return
		}

		tag = form.Transform()

		res, err := tc.service.AddTag(&tag)
		if err != nil {
			common.Response(c, http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
			return
		}

		common.ResponseSuccess(c, res)

}

// // func UpdateTag(c *gin.Context) {

// // 	pageNum, pageSize := utils.GetPaginationParams(c)
// // 	fmt.Println(pageNum, pageSize)

// // 	tags, err := s.FindMany(map[string]interface{}{}, pageNum, pageSize)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		common.Response(http.StatusBadRequest, e.ERROR_AUTH, nil)
// // 		return
// // 	}
// // 	count, err := s.CountAll()
// // 	fmt.Println(err)

// // 	tagsSerializer := serializers.TagsSerializer{tags}
// // 	common.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
// // 		"list":  tagsSerializer.Response(),
// // 		"total": count,
// // 	})
// // }

// // func DeleteTag(c *gin.Context) {
// // 	common := helper.Gin{C: c}

// // 	pageNum, pageSize := utils.GetPaginationParams(c)
// // 	fmt.Println(pageNum, pageSize)

// // 	tags, err := s.FindMany(map[string]interface{}{}, pageNum, pageSize)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		common.Response(http.StatusBadRequest, e.ERROR_AUTH, nil)
// // 		return
// // 	}
// // 	count, err := s.CountAll()
// // 	fmt.Println(err)

// // 	tagsSerializer := serializers.TagsSerializer{tags}
// // 	common.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
// // 		"list":  tagsSerializer.Response(),
// // 		"total": count,
// // 	})
// // }
