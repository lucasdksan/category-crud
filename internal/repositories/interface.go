package repositories

import "category/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
	Delete(index uint) ([]*entities.Category, error)
	Update(index uint, name string) error
	Read(index uint) (*entities.Category, error)
}
