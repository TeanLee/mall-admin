package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
)

type AdminAPI struct {
	PermissionSrv service.PermissionSrv
}

func (AdminAPI) GetAdminList(c *gin.Context) {
	list, err := service.GetAdminList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (AdminAPI) GetRoles(c *gin.Context) {
	roles, err := service.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func (AdminAPI) UpdateRole(c *gin.Context) {
	// 解析 query 的参数
	adminId := c.Query("admin_id")
	if adminId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The lack of admin id"})
		return
	}

	id, err := strconv.Atoi(adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "请说明要修改的 Admin Id"})
		return
	}

	var params model.Admin
	err = c.Bind(&params)

	// 修改 admin.username、admin.password
	err = service.UpdateAdminUsernameAndPassword(id, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}

	// 修改权限
	var paramRole model.RoleAdmin
	err = c.Bind(&paramRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}

	// 前端传递参数时，admin id 在 query 中，role id 在 body 中
	err = service.UpdateAdminRole(id, paramRole.RoleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "success updated"})
}
