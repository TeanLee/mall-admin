package api

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/config"
	"mall-admin-server-go/model/admin"
	"mall-admin-server-go/service"
	"net/http"
	"strings"
)

type LoginAPI struct {
	LoginSrv service.LoginSrv
}

type LoginParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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

func (a LoginAPI) Login(c *gin.Context) {
	// 获取 body 中的所有数据
	var loginParam LoginParam

	err := c.BindJSON(&loginParam)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	username := loginParam.UserName
	password := loginParam.Password

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "hello" || password != "itsme" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	session := sessions.Default(c)

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	db := config.InitMysql()

	var (
		admin []admin.Admin
	)

	res := db.Find(&admin)

	fmt.Println(res.RowsAffected)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}
