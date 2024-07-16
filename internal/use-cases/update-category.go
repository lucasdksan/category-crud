package use_cases

import (
	"category/internal/repositories"
)

type updateCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewUpdateCategoryUseCase(repository repositories.ICategoryRepository) *updateCategoryUseCase {
	return &updateCategoryUseCase{repository}
}

func (u *updateCategoryUseCase) Execute(index uint, name string) error {
	if err := u.repository.Update(index, name); err != nil {
		return err
	}

	return nil
}
