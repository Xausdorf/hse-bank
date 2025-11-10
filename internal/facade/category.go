package facade

import "github.com/Xausdorf/hse-bank/internal/domain"

type CategoryFacade struct {
	factory CategoryFactory
	repo    CategoryRepository
}

func (f *CategoryFacade) CreateCategory(name string, operationType domain.OperationType) (*domain.Category, error) {
	category, err := f.factory.Create(name, operationType)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (f *CategoryFacade) UpdateCategory(id, name string, opType domain.OperationType) error {
	category, err := f.factory.CreateWithID(id, name, opType)
	if err != nil {
		return err
	}
	if err := f.repo.Update(category); err != nil {
		return err
	}
	return nil
}

func (f *CategoryFacade) GetCategoryByID(id string) (*domain.Category, error) {
	category, err := f.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (f *CategoryFacade) DeleteCategory(id string) error {
	if err := f.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (f *CategoryFacade) GetAllCategories() []*domain.Category {
	return f.repo.GetAll()
}

func NewCategoryFacade(factory CategoryFactory, repo CategoryRepository) *CategoryFacade {
	return &CategoryFacade{
		factory: factory,
		repo:    repo,
	}
}

func (f *CategoryFacade) CreateCategoryWithID(id, name string, operationType domain.OperationType) (*domain.Category, error) {
	category, err := f.factory.CreateWithID(id, name, operationType)
	if err != nil {
		return nil, err
	}
	if err := f.repo.Save(category); err != nil {
		return nil, err
	}
	return category, nil
}
