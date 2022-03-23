package service

import (
	"github.com/rs/zerolog/log"
	"mall-admin-server-go/model"
)

type OrderSrv struct{}

type ProductItem struct {
	Title      string  `json:"title"`
	Banner     string  `json:"banner"`
	Price      float64 `json:"price"`
	CategoryId int     `json:"category_id"`
}

type OrderItem struct {
	Count       int         `json:"count"`
	ProductItem ProductItem `json:"product,inline"`
}

func newOrderItem(productItem ProductItem) OrderItem {
	return OrderItem{
		Count:       0,
		ProductItem: productItem,
	}
}

type OrderResType struct {
	OrderItems []OrderItem `json:"order_items"`
	OrderTime  string      `json:"order_time"`
	UserName   string      `json:"user_name"`
	Status     int         `json:"status"`
	Address    string      `json:"address"`
	Phone      string      `json:"phone"`
	Receiver   string      `json:"receiver"`
}

func GetOrders(offset int, pageSize int) ([]OrderResType, int64, error) {
	var orders []model.Order
	orders, err := model.GetOrders(offset, pageSize)

	var resOrder []OrderResType
	total := model.CountOrder()

	// 根据订单中的 productId 查找对应的商品
	for _, order := range orders {
		var tempOrder OrderResType
		for _, item := range order.OrderItem {
			var tempItem ProductItem
			var product model.Product
			product, err = model.GetProductById(item.ProductID)
			if err != nil {
				return nil, 0, err
			}
			tempItem.Title = product.Title
			tempItem.Banner = product.Banner
			tempItem.Price = product.Price
			tempItem.CategoryId = product.CategoryId
			log.Info().Msgf("tempItem：%v", tempItem)

			tempOrderItem := newOrderItem(tempItem)
			tempOrderItem.Count = item.Count
			tempOrder.OrderItems = append(tempOrder.OrderItems, tempOrderItem)

			user, err := GetUserById(order.Uid)
			if err != nil {
				return nil, 0, err
			}

			tempOrder.UserName = user.Username
			tempOrder.Phone = user.Phone
			tempOrder.Address = user.Address
			tempOrder.Receiver = user.Receiver
			tempOrder.OrderTime = order.OrderTime
		}
		resOrder = append(resOrder, tempOrder)
	}

	if err != nil {
		return nil, total, err
	}
	return resOrder, total, nil
}
