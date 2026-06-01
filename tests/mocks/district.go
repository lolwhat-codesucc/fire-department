package mocks

import (
	"context"
	"fire-department/modules/models"
	"github.com/stretchr/testify/mock"
)

type DistrictRepository struct {
	mock.Mock
}

func (m *DistrictRepository) Create(ctx context.Context, d *models.District) (int, error) {
	args := m.Called(ctx, d)
	return args.Int(0), args.Error(1)
}

func (m *DistrictRepository) GetByID(ctx context.Context, id int) (*models.District, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.District), args.Error(1)
}

func (m *DistrictRepository) GetAll(ctx context.Context) ([]models.District, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.District), args.Error(1)
}

func (m *DistrictRepository) Update(ctx context.Context, d *models.District) error {
	args := m.Called(ctx, d)
	return args.Error(0)
}

func (m *DistrictRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DistrictRepository) AddSpecialization(ctx context.Context, districtID, specID int) error {
	args := m.Called(ctx, districtID, specID)
	return args.Error(0)
}

func (m *DistrictRepository) RemoveSpecialization(ctx context.Context, districtID, specID int) error {
	args := m.Called(ctx, districtID, specID)
	return args.Error(0)
}

func (m *DistrictRepository) GetSpecializations(ctx context.Context, districtID int) ([]models.Specialization, error) {
	args := m.Called(ctx, districtID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Specialization), args.Error(1)
}
