package repositories

import "category/internal/entities"

type InMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *InMemoryCategoryRepository) Save(category *entities.Category) error {
	r.db = append(r.db, category)

	return nil
}

func (r *InMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}
