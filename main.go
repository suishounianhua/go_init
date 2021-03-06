package main

/*
 * @Script: main.go
 * @Author: pangxiaobo
 * @Email: 10846295@qq.com
 * @Create At: 2018-11-06 14:49:41
 * @Last Modified By: pangxiaobo
 * @Last Modified At: 2018-11-08 20:19:13
 * @Description: This is description.
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go_init/libs"
	"github.com/go_init/models"
	"github.com/go_init/routers"
	"net/http"
	"runtime"
)

func main() {

	serverConfig := libs.LoadServerConfig()
	models.InitDB(serverConfig)
	defer models.DB.Close()

	gin.SetMode(serverConfig.RunMode)
	//gin.DisableConsoleColor()

	//set the number of CPU processor will be used
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := routers.SetupRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConfig.HTTPPort),
		Handler:        router,
		ReadTimeout:    serverConfig.ReadTimeout,
		WriteTimeout:   serverConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
