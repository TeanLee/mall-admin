package service

import (
	"encoding/json"
	"mall-admin-server-go/model"
)

func ExistByID(id int) (bool, error) {
	return model.ExistAdminByID(id)
}

func GetAdmins() ([]model.Admin, error) {
	return model.GetAdmins()
}

func GetAdminByUsername(username string) (*model.Admin, error) {
	return model.GetAdminByUsername(username)
}

func GetRoles() ([]model.Role, error) {
	roles, err := model.GetRoles()
	if err != nil {
		return roles, err
	}
	return roles, nil
}

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

func UpdateAdminUsernameAndPassword(adminId int, data interface{}) error {
	marshalAdmin, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var admin model.Admin
	err = json.Unmarshal(marshalAdmin, &admin)
	if err != nil {
		return err
	}
	admin.AdminId = adminId
	return model.UpdateAdmin(adminId, admin)
}

func UpdateAdminRole(adminId int, roleId int) error {
	return model.UpdateAdminRole(adminId, roleId)
}

func DeleteAdmin(adminId int) error {
	if err := model.DeleteAdminInAdmin(adminId); err != nil {
		return err
	}
	if err := model.DeleteAdminInRoleAdmin(adminId); err != nil {
		return err
	}
	return nil
}
