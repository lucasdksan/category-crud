package repositories

import (
	"category/internal/entities"
	"errors"
)

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

func (r *InMemoryCategoryRepository) Delete(index uint) ([]*entities.Category, error) {
	if index >= uint(len(r.db)) {
		return nil, errors.New("index out of range")
	}

	r.db = append(r.db[:index], r.db[index+1:]...)

	return r.db, nil
}

func (r *InMemoryCategoryRepository) Update(index uint, name string) error {
	if index >= uint(len(r.db)) {
		return errors.New("index out of range")
	}

	r.db[index].Name = name

	return nil
}

func (r *InMemoryCategoryRepository) Read(index uint) (*entities.Category, error) {
	if index >= uint(len(r.db)) {
		return nil, errors.New("index out of range")
	}

	return r.db[index], nil
}
