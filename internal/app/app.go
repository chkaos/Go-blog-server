package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"Go-blog-server/internal/models"
	"Go-blog-server/internal/routers"
	"Go-blog-server/pkg/redis"
	"Go-blog-server/pkg/setting"
)

var server http.Server

//启动服务器
func Launch() {
	models.InitModels()
	redis.InitRedis()
	r := routers.InitRouter()
	startServer(r)
}

//关闭操作
func Destory() {
	models.CloseDB()
	redis.CloseRedis()
}

func startServer(router *gin.Engine) {
	gin.SetMode(setting.RunMode)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
