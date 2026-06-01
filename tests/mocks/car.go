package mocks

import (
    "context"
    "fire-department/modules/models"
    "github.com/stretchr/testify/mock"
)

type CarRepository struct {
    mock.Mock
}

func (m *CarRepository) Create(ctx context.Context, c *models.Car) (int, error) {
    args := m.Called(ctx, c)
    return args.Int(0), args.Error(1)
}

func (m *CarRepository) GetByID(ctx context.Context, id int) (*models.Car, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.Car), args.Error(1)
}

func (m *CarRepository) GetAll(ctx context.Context) ([]models.Car, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Car), args.Error(1)
}

func (m *CarRepository) GetByModel(ctx context.Context, modelID int) ([]models.Car, error) {
    args := m.Called(ctx, modelID)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Car), args.Error(1)
}

func (m *CarRepository) Update(ctx context.Context, c *models.Car) error {
    args := m.Called(ctx, c)
    return args.Error(0)
}

func (m *CarRepository) Delete(ctx context.Context, id int) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}
