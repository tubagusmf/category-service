package usecase

import "github.com/tubagusmf/category-service/internal/model"

type categoryUsecase struct {
	categoryUsecase model.CategoryRepository
}

func NewCategoryUsecase(cr model.CategoryRepository) model.CategoryUsecase {
	return &categoryUsecase{categoryUsecase: cr}
}

func (cu *categoryUsecase) FindCategoryById(id string) (model.Category, error) {
	return cu.categoryUsecase.FindCategoryById(id)
}
