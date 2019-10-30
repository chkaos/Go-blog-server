package main

import (
	"fmt"
	"log"
	"net/http"

	_ "Go-blog-server/internal/models"
	"Go-blog-server/internal/routers"
	"Go-blog-server/pkg/setting"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
