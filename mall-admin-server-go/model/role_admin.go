package model

import "gorm.io/gorm"

type RoleAdmin struct {
	Id      int `json:"id"`
	RoleId  int `json:"role_id"`
	AdminId int `json:"admin_id"`
}

func (RoleAdmin) TableName() string {
	return "role_admin"
}

func GetAdminIdByRoleId(roleId int) (int, error) {
	var roleAdmin RoleAdmin
	err := db.Where("role_id = ?", roleId).First(&roleAdmin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}

	return roleAdmin.AdminId, nil
}

func GetRoleIdByAdminId(adminId int) (int, error) {
	var roleAdmin RoleAdmin
	err := db.Where("admin_id = ?", adminId).First(&roleAdmin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return -1, err
	}

	return roleAdmin.RoleId, nil
}

func UpdateAdminIdByRoleId(roleId int, adminId int) {
	db.Model(&RoleAdmin{}).Where("role_id = ?", roleId).Update("admin_id", adminId)
}

func UpdateRoleIdByAdminId(adminId int, roleId int) {
	db.Model(&RoleAdmin{}).Where("admin_id = ?", adminId).Update("role_id", roleId)
}
