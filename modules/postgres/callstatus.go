package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CallStatusRepo struct {
	pool *pgxpool.Pool
}

func NewCallStatusRepo(pool *pgxpool.Pool) repository.CallStatusRepository {
	return &CallStatusRepo{pool: pool}
}

func (r *CallStatusRepo) Create(ctx context.Context, cs *models.CallStatus) (int, error) {
	query := `INSERT INTO Call_status (name) VALUES ($1) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query, cs.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create call status: %w", err)
	}
	return id, nil
}

func (r *CallStatusRepo) GetByID(ctx context.Context, id int) (*models.CallStatus, error) {
	query := `SELECT id, name FROM Call_status WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var cs models.CallStatus
	err := row.Scan(&cs.ID, &cs.Name)
	if err != nil {
		return nil, fmt.Errorf("get call status by id: %w", err)
	}
	return &cs, nil
}

func (r *CallStatusRepo) GetAll(ctx context.Context) ([]models.CallStatus, error) {
	query := `SELECT id, name FROM Call_status`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all call statuses: %w", err)
	}
	defer rows.Close()
	var statuses []models.CallStatus
	for rows.Next() {
		var cs models.CallStatus
		if err := rows.Scan(&cs.ID, &cs.Name); err != nil {
			return nil, fmt.Errorf("scan call status: %w", err)
		}
		statuses = append(statuses, cs)
	}
	return statuses, rows.Err()
}

func (r *CallStatusRepo) Update(ctx context.Context, cs *models.CallStatus) error {
	query := `UPDATE Call_status SET name = $1 WHERE id = $2`
	tag, err := r.pool.Exec(ctx, query, cs.Name, cs.ID)
	if err != nil {
		return fmt.Errorf("update call status: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("call status with id %d not found", cs.ID)
	}
	return nil
}

func (r *CallStatusRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Call_status WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete call status: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("call status with id %d not found", id)
	}
	return nil
}
