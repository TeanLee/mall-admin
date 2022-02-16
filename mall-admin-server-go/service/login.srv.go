package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
)

//var (
//	session =
//)

type LoginSrv struct {
	userName string
	password string
	role     string
}

type LoginBody struct {
	userName string `json:"username"`
	password string `json:"password"`
}

func (a LoginSrv) VerifyByUsername(username string, password string) (bool, *model.Admin) {
	userInfo, _ := model.GetAdminByUsername(username)

	// Check for username and password match, usually from a database
	if password != userInfo.Password {
		return false, nil
	}

	return true, userInfo
}

func GetCurrentRole(c *gin.Context) *model.Role {
	// 获取当前 user 信息
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		return nil
	}
	userId := user.(int)
	roleId, err := model.GetRoleIdByAdminId(userId)
	if err != nil {
		panic(err)
		return nil
	}
	role, err := model.GetRoleByRoleId(roleId)
	if err != nil {
		panic(err)
		return nil
	}
	return role
}
