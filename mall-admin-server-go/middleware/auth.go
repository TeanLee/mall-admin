package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/api"
	"net/http"
)

const (
	userkey = "user"
)

// AuthRequired is a simple middleware to check the session
// 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

// SuperAdminRequired 需要超管权限
func SuperAdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		isSuperAdmin := api.IsSuperAdmin(c)
		if isSuperAdmin == false {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "need super-admin permissions"})
			return
		}
		c.Next()
	}
}
