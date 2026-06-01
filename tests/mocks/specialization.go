package mocks

import (
	"context"
	"fire-department/modules/models"
	"github.com/stretchr/testify/mock"
)

type SpecializationRepository struct {
	mock.Mock
}

func (m *SpecializationRepository) Create(ctx context.Context, s *models.Specialization) (int, error) {
	args := m.Called(ctx, s)
	return args.Int(0), args.Error(1)
}

func (m *SpecializationRepository) GetByID(ctx context.Context, id int) (*models.Specialization, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Specialization), args.Error(1)
}

func (m *SpecializationRepository) GetAll(ctx context.Context) ([]models.Specialization, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Specialization), args.Error(1)
}

func (m *SpecializationRepository) Update(ctx context.Context, s *models.Specialization) error {
	args := m.Called(ctx, s)
	return args.Error(0)
}

func (m *SpecializationRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
