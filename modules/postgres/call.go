package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CallRepo struct {
	pool *pgxpool.Pool
}

func NewCallRepo(pool *pgxpool.Pool) repository.CallRepository {
	return &CallRepo{pool: pool}
}

func (r *CallRepo) Create(ctx context.Context, c *models.Call) (int, error) {
	query := `INSERT INTO Call (team_id, district_id, car_id, time, status_id, comment)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query,
		c.TeamID, c.DistrictID, c.CarID, c.Time, c.StatusID, c.Comment,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create call: %w", err)
	}
	return id, nil
}

func (r *CallRepo) GetByID(ctx context.Context, id int) (*models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var c models.Call
	err := row.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment)
	if err != nil {
		return nil, fmt.Errorf("get call by id: %w", err)
	}
	return &c, nil
}

func (r *CallRepo) GetAll(ctx context.Context) ([]models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all calls: %w", err)
	}
	defer rows.Close()
	var calls []models.Call
	for rows.Next() {
		var c models.Call
		if err := rows.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment); err != nil {
			return nil, fmt.Errorf("scan call: %w", err)
		}
		calls = append(calls, c)
	}
	return calls, rows.Err()
}

func (r *CallRepo) GetByTeam(ctx context.Context, teamNumber int) ([]models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call WHERE team_id = $1`
	rows, err := r.pool.Query(ctx, query, teamNumber)
	if err != nil {
		return nil, fmt.Errorf("get calls by team: %w", err)
	}
	defer rows.Close()
	var calls []models.Call
	for rows.Next() {
		var c models.Call
		if err := rows.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment); err != nil {
			return nil, err
		}
		calls = append(calls, c)
	}
	return calls, rows.Err()
}

func (r *CallRepo) GetByCar(ctx context.Context, carID int) ([]models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call WHERE car_id = $1`
	rows, err := r.pool.Query(ctx, query, carID)
	if err != nil {
		return nil, fmt.Errorf("get calls by car: %w", err)
	}
	defer rows.Close()
	var calls []models.Call
	for rows.Next() {
		var c models.Call
		if err := rows.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment); err != nil {
			return nil, err
		}
		calls = append(calls, c)
	}
	return calls, rows.Err()
}

func (r *CallRepo) GetByDistrict(ctx context.Context, districtID int) ([]models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call WHERE district_id = $1`
	rows, err := r.pool.Query(ctx, query, districtID)
	if err != nil {
		return nil, fmt.Errorf("get calls by district: %w", err)
	}
	defer rows.Close()
	var calls []models.Call
	for rows.Next() {
		var c models.Call
		if err := rows.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment); err != nil {
			return nil, err
		}
		calls = append(calls, c)
	}
	return calls, rows.Err()
}

func (r *CallRepo) GetByStatus(ctx context.Context, statusID int) ([]models.Call, error) {
	query := `SELECT id, team_id, district_id, car_id, time, status_id, comment FROM Call WHERE status_id = $1`
	rows, err := r.pool.Query(ctx, query, statusID)
	if err != nil {
		return nil, fmt.Errorf("get calls by status: %w", err)
	}
	defer rows.Close()
	var calls []models.Call
	for rows.Next() {
		var c models.Call
		if err := rows.Scan(&c.ID, &c.TeamID, &c.DistrictID, &c.CarID, &c.Time, &c.StatusID, &c.Comment); err != nil {
			return nil, err
		}
		calls = append(calls, c)
	}
	return calls, rows.Err()
}

func (r *CallRepo) Update(ctx context.Context, c *models.Call) error {
	query := `UPDATE Call SET team_id=$1, district_id=$2, car_id=$3, time=$4, status_id=$5, comment=$6 WHERE id=$7`
	tag, err := r.pool.Exec(ctx, query,
		c.TeamID, c.DistrictID, c.CarID, c.Time, c.StatusID, c.Comment, c.ID,
	)
	if err != nil {
		return fmt.Errorf("update call: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("call with id %d not found", c.ID)
	}
	return nil
}

func (r *CallRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Call WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete call: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("call with id %d not found", id)
	}
	return nil
}
