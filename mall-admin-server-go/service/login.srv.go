package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/config"
	"mall-admin-server-go/model"
	"mall-admin-server-go/utils"
)

var rdb = config.InitRedis()

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

func GetCurrentRole(c *gin.Context) (model.Role, error) {
	session := sessions.Default(c)

	// 获取当前 user 信息
	// 先从 redis 缓存中获取 role，如果缓存已经失效了，则查库（减少查库次数）
	// 在初次查询时，将 role 的信息存储在 redis 中，设置 5 秒过期时间

	var role model.Role

	user := session.Get("user")
	if user == nil {
		return role, nil
	}
	userId := user.(int)

	// 从 redis 中读取用户角色信息
	getRoleFromRedis, err := utils.GetUserRoleFromRedisByUserId(userId)
	if err == nil {
		return getRoleFromRedis, nil
	}

	// 如果从 redis 中没有读取到用户角色信息，则从 mysql 数据库中读取
	roleId, err := model.GetRoleIdByAdminId(userId)
	if err != nil {
		panic(err)
		return role, err
	}
	if role, err = model.GetRoleByRoleId(roleId); err != nil {
		return role, err
	}

	// redis 缓存中没有当前用户信息时，将用户信息写入。此处不设置缓存的过期时间，在更新用户权限时候需要更新当前用户权限
	if err := utils.SetRedisRoleByUserId(userId, role); err != nil {
		return role, err
	}
	return role, nil
}
