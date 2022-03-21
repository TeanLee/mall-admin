package model

type (
	StringInterfaceMap map[string]interface{}
	Order              struct {
		OrderItem string `json:"order_item"`
		OrderTime string `json:"order_time"`
		Uid       int    `json:"uid"` // float64 实际对应的是 double
		Status    int    `json:"status"`
	}
)

func (Order) TableName() string {
	return "order"
}

func GetOrders(offset int, pageSize int, productName string, category string) ([]Order, error) {
	var orders []Order
	//if err := db.Debug().Offset(offset).Limit(pageSize).Where("title LIKE ? and category_id LIKE ?", "%"+productName+"%", category+"%").Find(&orders).Error; err != nil {
	//	return orders, err
	//}
	if err := db.Debug().Offset(offset).Limit(pageSize).Find(&orders).Error; err != nil {
		return orders, err
	}
	return orders, nil
}

func CountOrder() int64 {
	var total int64 = 0
	if err := db.Debug().Model(Order{}).Count(&total).Error; err != nil {
		return -1
	}
	return total
}
