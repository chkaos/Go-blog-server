package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	pageNum := com.StrTo(c.Query("pageNum")).MustInt()
	pageSize := com.StrTo(c.Query("pageSize")).MustInt()
	pageNum = (pageNum - 1) * pageSize
	fmt.Println(pageNum, pageSize)
	if pageNum == 0 && pageSize == 0 {
		pageSize = -1
	}
	return pageNum, pageSize
}
