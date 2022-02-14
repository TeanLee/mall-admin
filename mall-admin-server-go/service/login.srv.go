package service

import "context"

type LoginSrv struct {
	userName string
	password string
	role     string
}

func (a *LoginSrv) Verify(ctx context.Context, userName, password string) bool {
	if userName == "admin" && password == "admin" {
		return true
	}
	return false
}
