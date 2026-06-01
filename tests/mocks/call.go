package mocks

import (
    "context"
    "fire-department/modules/models"
    "github.com/stretchr/testify/mock"
)

type CallRepository struct {
    mock.Mock
}

func (m *CallRepository) Create(ctx context.Context, c *models.Call) (int, error) {
    args := m.Called(ctx, c)
    return args.Int(0), args.Error(1)
}

func (m *CallRepository) GetByID(ctx context.Context, id int) (*models.Call, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.Call), args.Error(1)
}

func (m *CallRepository) GetAll(ctx context.Context) ([]models.Call, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Call), args.Error(1)
}

func (m *CallRepository) GetByTeam(ctx context.Context, teamNumber int) ([]models.Call, error) {
    args := m.Called(ctx, teamNumber)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Call), args.Error(1)
}

func (m *CallRepository) GetByCar(ctx context.Context, carID int) ([]models.Call, error) {
    args := m.Called(ctx, carID)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Call), args.Error(1)
}

func (m *CallRepository) GetByDistrict(ctx context.Context, districtID int) ([]models.Call, error) {
    args := m.Called(ctx, districtID)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Call), args.Error(1)
}

func (m *CallRepository) GetByStatus(ctx context.Context, statusID int) ([]models.Call, error) {
    args := m.Called(ctx, statusID)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Call), args.Error(1)
}

func (m *CallRepository) Update(ctx context.Context, c *models.Call) error {
    args := m.Called(ctx, c)
    return args.Error(0)
}

func (m *CallRepository) Delete(ctx context.Context, id int) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}
