package service

import "mall-admin-server-go/model"

func ExistByID(id int) (bool, error) {
	return model.ExistAdminByID(id)
}

func GetAdmins() ([]model.Admin, error) {
	return model.GetAdmins()
}

func GetAdminByUsername(username string) (*model.Admin, error) {
	return model.GetAdminByUsername(username)
}
