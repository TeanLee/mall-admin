package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/api"
	"net/http"
)

type Router struct {
	LoginAPI api.LoginAPI
}

func (a *Router) RegisterAPI(app *gin.Engine) {
	app.POST("/login", a.LoginAPI.Login)
	app.GET("/logout", a.LoginAPI.Logout)

	private := app.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", a.LoginAPI.Me)
		private.GET("/status", a.LoginAPI.Status)
	}
}

const (
	userkey = "user"
)

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
