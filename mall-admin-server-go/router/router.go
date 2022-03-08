package router

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/api"
	"mall-admin-server-go/middleware"
)

type Router struct {
	LoginAPI    api.LoginAPI
	CategoryApi api.CategoryAPI
	ProductApi  api.ProductAPI
	UserApi     api.UserAPI
	AdminApi    api.AdminAPI
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
	category.Use(middleware.AuthRequired())
	{
		category.GET("", a.CategoryApi.GetCategories)
		category.POST("/:id", a.CategoryApi.UpdateCategory)
	}
	//category.Use(middleware.AuthRequired())
	//{
	//	category.GET("/", a.CategoryApi.GetCategories)
	//}

	product := app.Group("/product")
	product.Use(middleware.AuthRequired())
	{
		product.GET("", a.ProductApi.GetProducts)
		product.POST("/:id", a.ProductApi.UpdateProduct)
		product.POST("/add", a.ProductApi.AddProduct)
		product.GET("/most-add", a.ProductApi.MostAdd)
	}

	user := app.Group("/user")
	user.Use(middleware.AuthRequired())
	{
		user.GET("", a.UserApi.GetUsers)
		user.POST("add", a.UserApi.AddUser)
	}

	admin := app.Group("/permission")
	admin.Use(middleware.AuthRequired())
	{
		admin.GET("", a.AdminApi.GetAdminList)
		admin.GET("/roles", a.AdminApi.GetRoles)
		admin.POST("/update-role/:admin_id", a.AdminApi.UpdateRole)
	}
}
