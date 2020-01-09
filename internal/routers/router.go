package routers

import (
	"github.com/gin-gonic/gin"

	_ "Go-blog-server/docs"
	"Go-blog-server/internal/routers/api"
	_ "Go-blog-server/internal/routers/api/admin"
	"Go-blog-server/pkg/setting"

	"Go-blog-server/internal/middleware"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())

	gin.SetMode(setting.RunMode)

	

	apiAdmin := r.Group("/api/admin")
	{
		apiAdmin.POST("/auth", api.GetAuth)

		// apiAdmin.GET("/tags", admin.GetTags)
		// apiAdmin.POST("/tags", admin.AddTag)
		// apiAdmin.PUT("/tags/:id", admin.EditTag)
		// apiAdmin.DELETE("/tags/:id", admin.DeleteTag)
	}

	// apiv1 := r.Group("/api/admin")
	// {
	// 	apiv1.GET("/tags", v1.GetTags)
	// 	apiv1.POST("/tags", v1.AddTag)
	// 	apiv1.PUT("/tags/:id", v1.EditTag)
	// 	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	// 	apiv1.GET("/category", v1.GetCategorys)
	// 	apiv1.POST("/category", v1.AddCategory)
	// 	apiv1.PUT("/category/:id", v1.EditCategory)
	// 	apiv1.DELETE("/category/:id", v1.DeleteCategory)

	// 	apiv1.GET("/articles", v1.GetArticles)
	// 	apiv1.GET("/articles/:id", v1.GetArticle)
	// 	apiv1.POST("/articles", v1.AddArticle)
	// 	apiv1.PUT("/articles/:id", v1.EditArticle)
	// 	apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	// }

	if setting.RunMode == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
