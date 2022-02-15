package service

import (
	"mall-admin-server-go/model"
)

type LoginSrv struct {
	userName string
	password string
	role     string
}

type LoginBody struct {
	userName string `json:"username"`
	password string `json:"password"`
}

func (a LoginSrv) VerifyByUsername(username string, password string) bool {
	userInfo, _ := model.GetAdminByUsername(username)

	// Check for username and password match, usually from a database
	if password != userInfo.Password {
		return false
	}

	return true
}
