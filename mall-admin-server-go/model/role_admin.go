package model

import (
	"fmt"
	"gorm.io/gorm"
)

type RoleAdmin struct {
	Id      int `json:"id"`
	RoleId  int `json:"role_id"`
	AdminId int `json:"admin_id"`
}

func (RoleAdmin) TableName() string {
	return "role_admin"
}

type Permission struct {
	AdminId  int    `json:"admin_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleName string `json:"role_name"`
	RoleId   int    `json:"role_id"`
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

func GetAdminList() ([]Permission, error) {
	var adminList []Permission
	if err := db.Model(&RoleAdmin{}).Select("admin.admin_id, admin.username, admin.password, role.role_name, role.role_id").Joins("inner join mall.role on mall.role_admin.role_id = mall.role.role_id inner join mall.admin on mall.role_admin.admin_id = mall.admin.admin_id").Scan(&adminList).Error; err != nil {
		fmt.Println("err：", err)
		return adminList, err
	}
	//db.Debug().Raw("SELECT admin.username, admin.password, role.role_name FROM mall.role_admin inner join mall.role on mall.role_admin.role_id = mall.role.role_id inner join mall.admin on mall.role_admin.admin_id = mall.admin.admin_id").Find(&adminList)
	return adminList, nil
}

func UpdateAdminRole(adminId int, roleId int) error {
	if err := db.Model(&RoleAdmin{}).Where("admin_id = ?", adminId).Update("role_id", roleId).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAdminInRoleAdmin(id int) error {
	if err := db.Where("admin_id = ?", id).Delete(Admin{}).Error; err != nil {
		return err
	}

	return nil
}
