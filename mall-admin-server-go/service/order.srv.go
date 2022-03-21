package service

import "mall-admin-server-go/model"

type OrderSrv struct{}

func GetOrders(offset int, pageSize int, productName string, category string) ([]model.Order, int64, error) {
	products, err := model.GetOrders(offset, pageSize, productName, category)
	total := model.CountOrder()
	if err != nil {
		return nil, total, err
	}
	return products, total, nil
}
