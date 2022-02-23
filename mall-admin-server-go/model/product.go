package model

type Product struct {
	ProductId  int     `json:"product_id"`
	Banner     string  `json:"banner"`
	Price      float64 `json:"price"` // float64 实际对应的是 double
	OldPrice   float64 `json:"old_price"`
	CategoryId int     `json:"category_id"`
	Unit       string  `json:"unit"`
	Title      string  `json:"title"`
	SubTitle   string  `json:"sub_title"`
}

func (Product) TableName() string {
	return "product"
}

func GetProducts() ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func GetProductsByPagination(offset int, pageSize int) ([]Product, error) {
	var products []Product
	if err := db.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func GetProductById(productId int) (Product, error) {
	var product Product
	if err := db.Where("product_id = ?", productId).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(productId int, data Product) error {
	var oldProduct Product
	if err := db.Where("category_id = ?", productId).First(&oldProduct).Error; err != nil {
		return err
	}
	/**
	Updates multiple columns

	// Update attributes with `struct`, will only update non-zero fields（不会更新主键 id）
	db.Model(&user).Where("category_id = ?", categoryId).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	*/
	if err := db.Model(&oldProduct).Where("category_id = ?", productId).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func CountProduct() int64 {
	var total int64 = 0
	if err := db.Model(Product{}).Count(&total).Error; err != nil {
		return -1
	}
	return total
}
