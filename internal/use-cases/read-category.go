package use_cases

import (
	"category/internal/entities"
	"category/internal/repositories"
)

type readCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewReadCategoryUseCase(repository repositories.ICategoryRepository) *readCategoryUseCase {
	return &readCategoryUseCase{repository}
}

func (u *readCategoryUseCase) Execute(index uint) (*entities.Category, error) {
	category, err := u.repository.Read(index)

	if err != nil {
		return nil, err
	}

	return category, nil
}
