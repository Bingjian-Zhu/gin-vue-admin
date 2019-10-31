package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bingjian-zhu/gin-vue-admin/common/setting"
	"github.com/bingjian-zhu/gin-vue-admin/routers"
)

func main() {
	router := routers.InitRouter()
	conf := setting.Config.Server
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", conf.Port),
		Handler:        router,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
