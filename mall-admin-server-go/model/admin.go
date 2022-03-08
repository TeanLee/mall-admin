package model

import (
	"gorm.io/gorm"
	"mall-admin-server-go/config"
)

type Admin struct {
	AdminId  int    `json:"admin_id"`
	Username string `json:"username, omitempty"`
	Password string `json:"password"`
}

func (Admin) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为testgorms（结构体+s）
	return "admin"
}

var db = config.InitMysql()

func ExistAdminByID(id int) (bool, error) {
	var admin Admin
	err := db.Where("admin_id = ?", id, 0).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if admin.AdminId > 0 {
		return true, nil
	}

	return false, nil
}

func ExistAdminByUsername(username string) (bool, error) {
	var admin Admin
	err := db.Where("username = ?", username, 0).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if admin.Username != "" {
		return true, nil
	}

	return false, nil
}

func GetAdminById(id int) (*Admin, error) {
	var admin Admin
	err := db.Where("admin_id = ?", id).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if admin.AdminId > -1 {
		return &admin, nil
	}

	return nil, gorm.ErrRecordNotFound
}

func GetAdminByUsername(username string) (*Admin, error) {
	var admin Admin
	err := db.Where("username = ?", username).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if admin.AdminId > -1 {
		return &admin, nil
	}

	return nil, gorm.ErrRecordNotFound
}

func GetAdmins() ([]Admin, error) {
	var admins []Admin
	err := db.Find(&admins).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return admins, nil
}

func UpdateAdmin(id int, data interface{}) error {
	if err := db.Model(&Admin{}).Where("admin_id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddAdmin(data map[string]interface{}) error {
	admin := Admin{
		AdminId:  data["admin_id"].(int),
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAdminInAdmin(id int) error {
	if err := db.Where("admin_id = ?", id).Delete(Admin{}).Error; err != nil {
		return err
	}

	return nil
}
