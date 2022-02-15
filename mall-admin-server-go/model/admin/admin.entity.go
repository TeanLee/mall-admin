package admin

type Admin struct {
	adminId  string `json:"admin_id"`
	username string `json:"username"`
	password string `json:"password"`
}

func (Admin) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为testgorms（结构体+s）
	return "admin"
}
