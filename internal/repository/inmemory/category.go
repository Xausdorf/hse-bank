package inmemory

import (
	"errors"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type CategoryRepository struct {
	categoryByID map[string]*domain.Category
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categoryByID: make(map[string]*domain.Category),
	}
}

func (r *CategoryRepository) Save(category *domain.Category) error {
	if category == nil {
		return errors.New("category is nil")
	}
	if _, ok := r.categoryByID[category.ID()]; !ok {
		r.categoryByID[category.ID()] = category
	}
	return nil
}

func (r *CategoryRepository) GetByID(id string) (*domain.Category, error) {
	category, ok := r.categoryByID[id]
	if !ok {
		return nil, errors.New("category not found")
	}
	return category, nil
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	if category == nil {
		return errors.New("category is nil")
	}
	if _, ok := r.categoryByID[category.ID()]; !ok {
		return errors.New("category not found")
	}
	r.categoryByID[category.ID()] = category
	return nil
}

func (r *CategoryRepository) Delete(id string) error {
	if _, ok := r.categoryByID[id]; !ok {
		return errors.New("category not found")
	}
	delete(r.categoryByID, id)
	return nil
}

func (r *CategoryRepository) GetAll() []*domain.Category {
	categories := make([]*domain.Category, 0, len(r.categoryByID))
	for _, category := range r.categoryByID {
		categories = append(categories, category)
	}
	return categories
}
