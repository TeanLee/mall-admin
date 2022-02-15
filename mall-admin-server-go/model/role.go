package model

import "gorm.io/gorm"

type Role struct {
	Id       int    `json:"-"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func (Role) TableName() string {
	return "role"
}

func GetRoles() ([]Role, error) {
	var roles []Role
	err := db.Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
}

func GetRoleByRoleId(id int) (*Role, error) {
	var role Role
	err := db.Where("role_id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &role, nil
}

func AddRole(data map[string]interface{}) error {
	role := Role{
		RoleId:   data["role_id"].(int),
		RoleName: data["role_name"].(string),
	}
	if err := db.Create(&role).Error; err != nil {
		return err
	}

	return nil
}

func UpdateRole(role Role) error {
	if err := db.Model(&Admin{}).Where("role_id = ?", role.RoleId).Updates(role).Error; err != nil {
		return err
	}
	return nil
}
