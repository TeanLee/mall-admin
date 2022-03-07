package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/service"
	"net/http"
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
