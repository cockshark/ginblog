package main

import (
	"fmt"
	"ginblog/pkg/setting"
	"ginblog/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
