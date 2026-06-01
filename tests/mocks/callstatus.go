package mocks

import (
    "context"
    "fire-department/modules/models"
    "github.com/stretchr/testify/mock"
)

type CallStatusRepository struct {
    mock.Mock
}

func (m *CallStatusRepository) Create(ctx context.Context, cs *models.CallStatus) (int, error) {
    args := m.Called(ctx, cs)
    return args.Int(0), args.Error(1)
}

func (m *CallStatusRepository) GetByID(ctx context.Context, id int) (*models.CallStatus, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.CallStatus), args.Error(1)
}

func (m *CallStatusRepository) GetAll(ctx context.Context) ([]models.CallStatus, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.CallStatus), args.Error(1)
}

func (m *CallStatusRepository) Update(ctx context.Context, cs *models.CallStatus) error {
    args := m.Called(ctx, cs)
    return args.Error(0)
}

func (m *CallStatusRepository) Delete(ctx context.Context, id int) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}
