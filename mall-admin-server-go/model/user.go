package model

import "gorm.io/gorm"

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Receiver string `json:"receiver"`
}

func (User) TableName() string {
	return "user"
}

func GetUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}

func GetUsersByPagination(offset int, pageSize int, username string, address string, phone string, receiver string) ([]User, error) {
	var users []User
	if err := db.Offset(offset).Limit(pageSize).Where("username LIKE ? and address LIKE ? and phone LIKE ? and receiver LIKE ?", "%"+username+"%", "%"+address+"%", "%"+phone+"%", "%"+receiver+"%").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func GetUserById(userId int) (User, error) {
	var user User
	if err := db.Where("uid = ?", userId).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(userId int, data User) error {
	var oldUser User
	if err := db.Where(" uid = ?", userId).First(&oldUser).Error; err != nil {
		return err
	}
	/**
	Updates multiple columns

	// Update attributes with `struct`, will only update non-zero fields（不会更新主键 id）
	db.Model(&user).Where("category_id = ?", categoryId).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	*/
	if err := db.Model(&oldUser).Where(" uid = ?", userId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func CountUser() int64 {
	var total int64 = 0
	if err := db.Model(User{}).Count(&total).Error; err != nil {
		return -1
	}
	return total
}

func AddUser(user User) (int, error) {
	sql := db.Create(&user)
	if err := sql.Error; err != nil {
		return -1, err
	}
	return 0, nil
}
