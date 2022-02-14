package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/service"
)

type LoginAPI struct {
	LoginSrv *service.LoginSrv
}

func (a *LoginAPI) Login(c *gin.Context, userName string, password string) {
	a.LoginSrv.Verify(c, userName, password)
}
