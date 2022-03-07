package service

import "mall-admin-server-go/model"

type PermissionSrv struct {
	adminName string
	password  string
	role      string
}

func GetAdminList() ([]model.Permission, error) {
	list, err := model.GetAdminList()
	if err != nil {
		return list, err
	}
	return list, err
}
