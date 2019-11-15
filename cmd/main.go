package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"context"
	"os"
	"os/signal"

	_ "Go-blog-server/internal/models"
	"Go-blog-server/internal/routers"
	"Go-blog-server/pkg/setting"
)

// @title Swagger Astella API
// @version 1.0
// @description This is a sample server celler server.

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <- quit

    log.Println("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
	
}
