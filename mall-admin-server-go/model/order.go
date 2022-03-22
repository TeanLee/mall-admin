package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Order struct {
	OrderItem OrderItemType `gorm:"TYPE:json" json:"order_item"`
	OrderTime string        `json:"order_time"`
	Uid       int           `json:"uid"` // float64 实际对应的是 double
	Status    int           `json:"status"`
}

type OrderItemType []struct {
	Count     int `json:"count"`
	ProductID int `json:"productId"`
}

// Scan 实现 sql.Scanner 接口，Scan 将 value 扫描至 OrderItemType
func (c *OrderItemType) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// Value 实现 driver.Valuer 接口，Value 返回 OrderItemType value
func (c OrderItemType) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

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
