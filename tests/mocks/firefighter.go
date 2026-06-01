package mocks

import (
    "context"
    "fire-department/modules/models"
    "github.com/stretchr/testify/mock"
)

type FirefighterRepository struct {
    mock.Mock
}

func (m *FirefighterRepository) Create(ctx context.Context, f *models.Firefighter) (int, error) {
    args := m.Called(ctx, f)
    return args.Int(0), args.Error(1)
}

func (m *FirefighterRepository) GetByID(ctx context.Context, id int) (*models.Firefighter, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*models.Firefighter), args.Error(1)
}

func (m *FirefighterRepository) GetAll(ctx context.Context) ([]models.Firefighter, error) {
    args := m.Called(ctx)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Firefighter), args.Error(1)
}

func (m *FirefighterRepository) GetByTeam(ctx context.Context, teamNumber int) ([]models.Firefighter, error) {
    args := m.Called(ctx, teamNumber)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).([]models.Firefighter), args.Error(1)
}

func (m *FirefighterRepository) Update(ctx context.Context, f *models.Firefighter) error {
    args := m.Called(ctx, f)
    return args.Error(0)
}

func (m *FirefighterRepository) Delete(ctx context.Context, id int) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}
