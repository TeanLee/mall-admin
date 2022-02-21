package service

import (
	"encoding/json"
	"mall-admin-server-go/model"
)

type CategorySrv struct{}

func GetCategories() ([]model.Category, error) {
	return model.GetCategories()
}

func GetCategoriesById(categoryId int) (model.Category, error) {
	return model.GetCategoryById(categoryId)
}

func UpdateCategory(categoryId int, data interface{}) error {
	marshalData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var category model.Category
	err = json.Unmarshal(marshalData, &category)
	if err != nil {
		return err
	}
	category.CategoryId = categoryId
	return model.UpdateCategory(categoryId, category)
}
