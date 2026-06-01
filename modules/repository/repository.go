package repository

import (
	"context"
	"fire-department/modules/models"
)

type RankRepository interface {
	Create(ctx context.Context, r *models.Rank) (int, error)
	GetByID(ctx context.Context, id int) (*models.Rank, error)
	GetAll(ctx context.Context) ([]models.Rank, error)
	Update(ctx context.Context, r *models.Rank) error
	Delete(ctx context.Context, id int) error
}

type SpecializationRepository interface {
	Create(ctx context.Context, s *models.Specialization) (int, error)
	GetByID(ctx context.Context, id int) (*models.Specialization, error)
	GetAll(ctx context.Context) ([]models.Specialization, error)
	Update(ctx context.Context, s *models.Specialization) error
	Delete(ctx context.Context, id int) error
}

type DistrictRepository interface {
	Create(ctx context.Context, d *models.District) (int, error)
	GetByID(ctx context.Context, id int) (*models.District, error)
	GetAll(ctx context.Context) ([]models.District, error)
	Update(ctx context.Context, d *models.District) error
	Delete(ctx context.Context, id int) error
	AddSpecialization(ctx context.Context, districtID, specID int) error
	RemoveSpecialization(ctx context.Context, districtID, specID int) error
	GetSpecializations(ctx context.Context, districtID int) ([]models.Specialization, error)
}

type CarModelRepository interface {
	Create(ctx context.Context, cm *models.CarModel) (int, error)
	GetByID(ctx context.Context, id int) (*models.CarModel, error)
	GetAll(ctx context.Context) ([]models.CarModel, error)
	Update(ctx context.Context, cm *models.CarModel) error
	Delete(ctx context.Context, id int) error
}

type CallStatusRepository interface {
	Create(ctx context.Context, cs *models.CallStatus) (int, error)
	GetByID(ctx context.Context, id int) (*models.CallStatus, error)
	GetAll(ctx context.Context) ([]models.CallStatus, error)
	Update(ctx context.Context, cs *models.CallStatus) error
	Delete(ctx context.Context, id int) error
}

type CarRepository interface {
	Create(ctx context.Context, c *models.Car) (int, error)
	GetByID(ctx context.Context, id int) (*models.Car, error)
	GetAll(ctx context.Context) ([]models.Car, error)
	GetByModel(ctx context.Context, modelID int) ([]models.Car, error)
	Update(ctx context.Context, c *models.Car) error
	Delete(ctx context.Context, id int) error
}

type TeamRepository interface {
	Create(ctx context.Context, t *models.Team) (int, error)
	GetByNumber(ctx context.Context, number int) (*models.Team, error)
	GetAll(ctx context.Context) ([]models.Team, error)
	GetByDistrict(ctx context.Context, districtID int) ([]models.Team, error)
	GetBySpecialization(ctx context.Context, specID int) ([]models.Team, error)
	Update(ctx context.Context, t *models.Team) error
	Delete(ctx context.Context, number int) error
}

type FirefighterRepository interface {
	Create(ctx context.Context, f *models.Firefighter) (int, error)
	GetByID(ctx context.Context, id int) (*models.Firefighter, error)
	GetAll(ctx context.Context) ([]models.Firefighter, error)
	GetByTeam(ctx context.Context, teamNumber int) ([]models.Firefighter, error)
	Update(ctx context.Context, f *models.Firefighter) error
	Delete(ctx context.Context, id int) error
}

type CallRepository interface {
	Create(ctx context.Context, c *models.Call) (int, error)
	GetByID(ctx context.Context, id int) (*models.Call, error)
	GetAll(ctx context.Context) ([]models.Call, error)
	GetByTeam(ctx context.Context, teamNumber int) ([]models.Call, error)
	GetByCar(ctx context.Context, carID int) ([]models.Call, error)
	GetByDistrict(ctx context.Context, districtID int) ([]models.Call, error)
	GetByStatus(ctx context.Context, statusID int) ([]models.Call, error)
	Update(ctx context.Context, c *models.Call) error
	Delete(ctx context.Context, id int) error
}
