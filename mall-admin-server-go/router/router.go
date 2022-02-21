package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/api"
	"mall-admin-server-go/middleware"
)

type Router struct {
	LoginAPI    api.LoginAPI
	CategoryApi api.CategoryAPI
}

func (a *Router) RegisterAPI(app *gin.Engine) {
	app.POST("/login", a.LoginAPI.Login)
	app.POST("/logout", a.LoginAPI.Logout)

	private := app.Group("/private")
	private.Use(middleware.AuthRequired()).Use(middleware.SuperAdminRequired())
	{
		private.GET("/me", a.LoginAPI.Me)
		private.GET("/status", a.LoginAPI.Status)
	}

	category := app.Group("/category")
	category.GET("", a.CategoryApi.GetCategories)
	category.POST("/:id", a.CategoryApi.UpdateCategory)
	//category.Use(middleware.AuthRequired())
	//{
	//	category.GET("/", a.CategoryApi.GetCategories)
	//}
}
