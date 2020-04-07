package main

import (
	"fmt"
	"net/http"

	"github.com/zzs/go-gin-example/pkg/setting"
	"github.com/zzs/go-gin-example/routers"
)

func main() {
	routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}