package mocks

import (
	"context"
	"fire-department/modules/models"
	"github.com/stretchr/testify/mock"
)

type CarModelRepository struct {
	mock.Mock
}

func (m *CarModelRepository) Create(ctx context.Context, cm *models.CarModel) (int, error) {
	args := m.Called(ctx, cm)
	return args.Int(0), args.Error(1)
}

func (m *CarModelRepository) GetByID(ctx context.Context, id int) (*models.CarModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CarModel), args.Error(1)
}

func (m *CarModelRepository) GetAll(ctx context.Context) ([]models.CarModel, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CarModel), args.Error(1)
}

func (m *CarModelRepository) Update(ctx context.Context, cm *models.CarModel) error {
	args := m.Called(ctx, cm)
	return args.Error(0)
}

func (m *CarModelRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
