package mocks

import (
	"context"
	"fire-department/modules/models" 
	"github.com/stretchr/testify/mock"
)

type RankRepository struct {
	mock.Mock
}

func (m *RankRepository) Create(ctx context.Context, r *models.Rank) (int, error) {
	args := m.Called(ctx, r)
	return args.Int(0), args.Error(1)
}

func (m *RankRepository) GetByID(ctx context.Context, id int) (*models.Rank, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Rank), args.Error(1)
}

func (m *RankRepository) GetAll(ctx context.Context) ([]models.Rank, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Rank), args.Error(1)
}

func (m *RankRepository) Update(ctx context.Context, r *models.Rank) error {
	args := m.Called(ctx, r)
	return args.Error(0)
}

func (m *RankRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
