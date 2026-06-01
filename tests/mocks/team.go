package mocks

import (
	"context"
	"fire-department/modules/models"
	"github.com/stretchr/testify/mock"
)

type TeamRepository struct {
	mock.Mock
}

func (m *TeamRepository) Create(ctx context.Context, t *models.Team) (int, error) {
	args := m.Called(ctx, t)
	return args.Int(0), args.Error(1)
}

func (m *TeamRepository) GetByNumber(ctx context.Context, number int) (*models.Team, error) {
	args := m.Called(ctx, number)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Team), args.Error(1)
}

func (m *TeamRepository) GetAll(ctx context.Context) ([]models.Team, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Team), args.Error(1)
}

func (m *TeamRepository) GetByDistrict(ctx context.Context, districtID int) ([]models.Team, error) {
	args := m.Called(ctx, districtID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Team), args.Error(1)
}

func (m *TeamRepository) GetBySpecialization(ctx context.Context, specID int) ([]models.Team, error) {
	args := m.Called(ctx, specID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Team), args.Error(1)
}

func (m *TeamRepository) Update(ctx context.Context, t *models.Team) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *TeamRepository) Delete(ctx context.Context, number int) error {
	args := m.Called(ctx, number)
	return args.Error(0)
}
