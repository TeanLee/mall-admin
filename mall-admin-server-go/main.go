package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/config"
	"mall-admin-server-go/middleware"
	"mall-admin-server-go/router"
)

func main() {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysession", config.InitSessionStore()))
	r2 := router.Router{}
	r2.RegisterAPI(r)
	r.Run(":8081")
}
