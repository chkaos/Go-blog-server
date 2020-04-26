package routers

import (
	"github.com/gin-gonic/gin"

	_ "Go-blog-server/docs"
	"Go-blog-server/internal/controllers"
	"Go-blog-server/pkg/setting"

	"Go-blog-server/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(middleware.CORS())

	gin.SetMode(setting.RunMode)

	uc := controllers.NewUserController()
	tc := controllers.NewTagController()
	cc := controllers.NewCategoryController()
	bc := controllers.NewBulletinController()
	fc := controllers.NewFileController()
	ac := controllers.NewArticleController()

	apiAdmin := r.Group("/api/admin")
	{
		apiAdmin.POST("/auth", uc.Auth)

		apiAdmin.GET("/file", fc.GetFiles)
		apiAdmin.POST("/file/upload", fc.Upload)

		apiAdmin.GET("/tag", tc.GetTags)
		apiAdmin.GET("/tags", tc.GetAllTags)
		apiAdmin.POST("/tag", tc.AddTag)
		apiAdmin.PUT("/tag", tc.UpdateTag)
		apiAdmin.DELETE("/tag/:id", tc.DeleteTag)

		apiAdmin.GET("/category", cc.GetCategorys)
		apiAdmin.GET("/categorys", cc.GetAllCategorys)
		apiAdmin.POST("/category", cc.AddCategory)
		apiAdmin.PUT("/category", cc.UpdateCategory)
		apiAdmin.DELETE("/category/:id", cc.DeleteCategory)

		apiAdmin.GET("/bulletin", bc.GetBulletins)
		apiAdmin.POST("/bulletin", bc.AddBulletin)
		apiAdmin.PUT("/bulletin", bc.UpdateBulletin)
		apiAdmin.DELETE("/bulletin/:id", bc.DeleteBulletin)

		apiAdmin.GET("/article/:id", ac.GetArticle)
		apiAdmin.GET("/article", ac.GetArticles)
		apiAdmin.POST("/article", ac.AddArticle)
		apiAdmin.PUT("/article", ac.UpdateArticle)
		apiAdmin.PUT("/article/state", ac.UpdateState)
		apiAdmin.DELETE("/article/:id", ac.DeleteArticle)
	}

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", tc.GetAllTags)

		apiv1.GET("/article/:id", ac.GetArticle)
		apiv1.GET("/article", ac.GetArticles)

		apiv1.GET("/bulletin", bc.GetBulletins)
	}

	if setting.RunMode == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
