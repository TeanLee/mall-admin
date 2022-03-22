package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"mall-admin-server-go/config"
	"mall-admin-server-go/middleware"
	"mall-admin-server-go/router"
	"os"
)

func main() {
	// 根据线程传进来的函数，可以设置日志的输出级别
	env := os.Getenv("env")
	if env == "prod" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if env == "stagging" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	r := gin.Default()

	r.Use(sessions.Sessions("mysession", config.InitSessionStore()))
	r.Use(middleware.Cors())
	r2 := router.Router{}
	r2.RegisterAPI(r)
	r.Run(":8081")
}
