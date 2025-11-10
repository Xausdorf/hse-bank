package inmemory

import (
	"errors"

	"github.com/Xausdorf/hse-bank/internal/domain"
)

type OperationRepository struct {
	operationByID map[string]*domain.Operation
}

func NewOperationRepository() *OperationRepository {
	return &OperationRepository{
		operationByID: make(map[string]*domain.Operation),
	}
}

func (r *OperationRepository) Save(operation *domain.Operation) error {
	if operation == nil {
		return errors.New("operation is nil")
	}
	if _, ok := r.operationByID[operation.ID()]; !ok {
		r.operationByID[operation.ID()] = operation
	}
	return nil
}

func (r *OperationRepository) GetByID(id string) (*domain.Operation, error) {
	operation, ok := r.operationByID[id]
	if !ok {
		return nil, errors.New("operation not found")
	}
	return operation, nil
}

func (r *OperationRepository) Update(operation *domain.Operation) error {
	if operation == nil {
		return errors.New("operation is nil")
	}
	if _, ok := r.operationByID[operation.ID()]; !ok {
		return errors.New("operation not found")
	}
	r.operationByID[operation.ID()] = operation
	return nil
}

func (r *OperationRepository) Delete(id string) error {
	if _, ok := r.operationByID[id]; !ok {
		return errors.New("operation not found")
	}
	delete(r.operationByID, id)
	return nil
}

func (r *OperationRepository) GetAll() []*domain.Operation {
	operations := make([]*domain.Operation, 0, len(r.operationByID))
	for _, operation := range r.operationByID {
		operations = append(operations, operation)
	}
	return operations
}
