package repository

import "github.com/tubagusmf/category-service/internal/model"

type categoryRepository struct {
}

func NewCategoryRepository() model.CategoryRepository {
	return &categoryRepository{}
}

func (cr *categoryRepository) FindCategoryById(id string) (model.Category, error) {
	var category model.Category

	if id != "1" {
		return category, nil
	}

	return model.Category{
		Id:        "1",
		Name:      "Hiburan",
		CreatedAt: "2024-02-11T15:04:05Z",
		UpdatedAt: "2024-02-11T15:04:05Z",
	}, nil
}
