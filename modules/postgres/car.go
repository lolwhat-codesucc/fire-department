package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CarRepo struct {
	pool *pgxpool.Pool
}

func NewCarRepo(pool *pgxpool.Pool) repository.CarRepository {
	return &CarRepo{pool: pool}
}

func (r *CarRepo) Create(ctx context.Context, c *models.Car) (int, error) {
	query := `INSERT INTO Car (model_id, acquisition_date, last_maintenance, ready)
	          VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query,
		c.ModelID, c.AcquisitionDate, c.LastMaintenance, c.Ready,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create car: %w", err)
	}
	return id, nil
}

func (r *CarRepo) GetByID(ctx context.Context, id int) (*models.Car, error) {
	query := `SELECT id, model_id, acquisition_date, last_maintenance, ready FROM Car WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var c models.Car
	err := row.Scan(&c.ID, &c.ModelID, &c.AcquisitionDate, &c.LastMaintenance, &c.Ready)
	if err != nil {
		return nil, fmt.Errorf("get car by id: %w", err)
	}
	return &c, nil
}

func (r *CarRepo) GetAll(ctx context.Context) ([]models.Car, error) {
	query := `SELECT id, model_id, acquisition_date, last_maintenance, ready FROM Car`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all cars: %w", err)
	}
	defer rows.Close()
	var cars []models.Car
	for rows.Next() {
		var c models.Car
		if err := rows.Scan(&c.ID, &c.ModelID, &c.AcquisitionDate, &c.LastMaintenance, &c.Ready); err != nil {
			return nil, fmt.Errorf("scan car: %w", err)
		}
		cars = append(cars, c)
	}
	return cars, rows.Err()
}

func (r *CarRepo) GetByModel(ctx context.Context, modelID int) ([]models.Car, error) {
	query := `SELECT id, model_id, acquisition_date, last_maintenance, ready FROM Car WHERE model_id = $1`
	rows, err := r.pool.Query(ctx, query, modelID)
	if err != nil {
		return nil, fmt.Errorf("get cars by model: %w", err)
	}
	defer rows.Close()
	var cars []models.Car
	for rows.Next() {
		var c models.Car
		if err := rows.Scan(&c.ID, &c.ModelID, &c.AcquisitionDate, &c.LastMaintenance, &c.Ready); err != nil {
			return nil, fmt.Errorf("scan car: %w", err)
		}
		cars = append(cars, c)
	}
	return cars, rows.Err()
}

func (r *CarRepo) Update(ctx context.Context, c *models.Car) error {
	query := `UPDATE Car SET model_id=$1, acquisition_date=$2, last_maintenance=$3, ready=$4 WHERE id=$5`
	tag, err := r.pool.Exec(ctx, query,
		c.ModelID, c.AcquisitionDate, c.LastMaintenance, c.Ready, c.ID,
	)
	if err != nil {
		return fmt.Errorf("update car: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car with id %d not found", c.ID)
	}
	return nil
}

func (r *CarRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Car WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete car: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("car with id %d not found", id)
	}
	return nil
}
