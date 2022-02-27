package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
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

	// 验证用户名和密码是否正确
	verify, admin := a.LoginSrv.VerifyByUsername(username, password)

	if verify == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	session := sessions.Default(c)

	// Save the username in the session
	session.Set(userkey, admin.AdminId) // In real world usage you'd set this to the users ID

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session", "5r": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func (a LoginAPI) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (a LoginAPI) Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	var role model.Role
	role, err := service.GetCurrentRole(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "role": role.RoleName})
}

func (a LoginAPI) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func IsSuperAdmin(c *gin.Context) bool {
	var role model.Role
	role, err := service.GetCurrentRole(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	fmt.Println("is super admin:", role.RoleId)
	if role.RoleId == -1 {
		return true
	} else {
		return false
	}
}
