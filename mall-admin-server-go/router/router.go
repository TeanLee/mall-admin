package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/api"
	"mall-admin-server-go/middleware"
)

type Router struct {
	LoginAPI api.LoginAPI
}

func (a *Router) RegisterAPI(app *gin.Engine) {
	app.POST("/login", a.LoginAPI.Login)
	app.GET("/logout", a.LoginAPI.Logout)

	private := app.Group("/private")
	private.Use(middleware.AuthRequired()).Use(middleware.SuperAdminRequired())
	{
		private.GET("/me", a.LoginAPI.Me)
		private.GET("/status", a.LoginAPI.Status)
	}
}
