package model

type Category struct {
	CategoryId int    `json:"category_id"`
	Category   string `json:"category"`
	Icon       string `json:"icon"`
	Color      string `json:"color"`
	Name       string `json:"name"`
}

func (Category) TableName() string {
	return "category"
}

func GetCategories() ([]Category, error) {
	var categories []Category
	if err := db.Find(&categories).Error; err != nil {
		return categories, err
	}
	return categories, nil
}

func GetCategoryById(categoryId int) (Category, error) {
	var category Category
	if err := db.Where("categoryId = ?", categoryId, 0).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func UpdateCategory(categoryId int, data Category) error {
	var oldCategory Category
	if err := db.Where("category_id = ?", categoryId).First(&oldCategory).Error; err != nil {
		return err
	}
	/**
	Updates multiple columns

	// Update attributes with `struct`, will only update non-zero fields（不会更新主键 id）
	db.Model(&user).Where("category_id = ?", categoryId).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	*/
	if err := db.Model(&oldCategory).Where("category_id = ?", categoryId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
