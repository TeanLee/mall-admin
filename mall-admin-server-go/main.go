package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/middleware"
	"mall-admin-server-go/router"
)

func main() {
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.Use(middleware.Cors())
	r2 := router.Router{}
	r2.RegisterAPI(r)
	r.Run()
}
