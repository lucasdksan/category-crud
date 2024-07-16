package use_cases

import (
	"category/internal/entities"
	"category/internal/repositories"
)

type deleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{repository}
}

func (u *deleteCategoryUseCase) Execute(index uint) ([]*entities.Category, error) {
	categories, err := u.repository.Delete(index)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
