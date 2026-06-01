package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CarModelRepo struct {
	pool *pgxpool.Pool
}

func NewCarModelRepo(pool *pgxpool.Pool) repository.CarModelRepository {
	return &CarModelRepo{pool: pool}
}

func (r *CarModelRepo) Create(ctx context.Context, cm *models.CarModel) (int, error) {
	query := `INSERT INTO Car_model (name, maintenance_period_days) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query, cm.Name, cm.MaintenancePeriodDays).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create car model: %w", err)
	}
	return id, nil
}

func (r *CarModelRepo) GetByID(ctx context.Context, id int) (*models.CarModel, error) {
	query := `SELECT id, name, maintenance_period_days FROM Car_model WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var cm models.CarModel
	err := row.Scan(&cm.ID, &cm.Name, &cm.MaintenancePeriodDays)
	if err != nil {
		return nil, fmt.Errorf("get car model by id: %w", err)
	}
	return &cm, nil
}

func (r *CarModelRepo) GetAll(ctx context.Context) ([]models.CarModel, error) {
	query := `SELECT id, name, maintenance_period_days FROM Car_model`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all car models: %w", err)
	}
	defer rows.Close()
	var carModels []models.CarModel
	for rows.Next() {
		var cm models.CarModel
		if err := rows.Scan(&cm.ID, &cm.Name, &cm.MaintenancePeriodDays); err != nil {
			return nil, fmt.Errorf("scan car model: %w", err)
		}
		carModels = append(carModels, cm)
	}
	return carModels, rows.Err()
}

func (r *CarModelRepo) Update(ctx context.Context, cm *models.CarModel) error {
	query := `UPDATE Car_model SET name = $1, maintenance_period_days = $2 WHERE id = $3`
	tag, err := r.pool.Exec(ctx, query, cm.Name, cm.MaintenancePeriodDays, cm.ID)
	if err != nil {
		return fmt.Errorf("update car model: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car model with id %d not found", cm.ID)
	}
	return nil
}

func (r *CarModelRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Car_model WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete car model: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car model with id %d not found", id)
	}
	return nil
}
