package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepo struct {
	pool *pgxpool.Pool
}

func NewTeamRepo(pool *pgxpool.Pool) repository.TeamRepository {
	return &TeamRepo{pool: pool}
}

func (r *TeamRepo) Create(ctx context.Context, t *models.Team) (int, error) {
	query := `INSERT INTO Team (specialization_id, district_id, car_id)
	          VALUES ($1, $2, $3) RETURNING number`
	var number int
	err := r.pool.QueryRow(ctx, query,
		t.SpecializationID, t.DistrictID, t.CarID,
	).Scan(&number)
	if err != nil {
		return 0, fmt.Errorf("create team: %w", err)
	}
	return number, nil
}

func (r *TeamRepo) GetByNumber(ctx context.Context, number int) (*models.Team, error) {
	query := `SELECT number, specialization_id, district_id, car_id FROM Team WHERE number = $1`
	row := r.pool.QueryRow(ctx, query, number)
	var t models.Team
	err := row.Scan(&t.Number, &t.SpecializationID, &t.DistrictID, &t.CarID)
	if err != nil {
		return nil, fmt.Errorf("get team by number: %w", err)
	}
	return &t, nil
}

func (r *TeamRepo) GetAll(ctx context.Context) ([]models.Team, error) {
	query := `SELECT number, specialization_id, district_id, car_id FROM Team`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all teams: %w", err)
	}
	defer rows.Close()
	var teams []models.Team
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.Number, &t.SpecializationID, &t.DistrictID, &t.CarID); err != nil {
			return nil, fmt.Errorf("scan team: %w", err)
		}
		teams = append(teams, t)
	}
	return teams, rows.Err()
}

func (r *TeamRepo) GetByDistrict(ctx context.Context, districtID int) ([]models.Team, error) {
	query := `SELECT number, specialization_id, district_id, car_id FROM Team WHERE district_id = $1`
	rows, err := r.pool.Query(ctx, query, districtID)
	if err != nil {
		return nil, fmt.Errorf("get teams by district: %w", err)
	}
	defer rows.Close()
	var teams []models.Team
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.Number, &t.SpecializationID, &t.DistrictID, &t.CarID); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, rows.Err()
}

func (r *TeamRepo) GetBySpecialization(ctx context.Context, specID int) ([]models.Team, error) {
	query := `SELECT number, specialization_id, district_id, car_id FROM Team WHERE specialization_id = $1`
	rows, err := r.pool.Query(ctx, query, specID)
	if err != nil {
		return nil, fmt.Errorf("get teams by specialization: %w", err)
	}
	defer rows.Close()
	var teams []models.Team
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.Number, &t.SpecializationID, &t.DistrictID, &t.CarID); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}
	return teams, rows.Err()
}

func (r *TeamRepo) Update(ctx context.Context, t *models.Team) error {
	query := `UPDATE Team SET specialization_id=$1, district_id=$2, car_id=$3 WHERE number=$4`
	tag, err := r.pool.Exec(ctx, query,
		t.SpecializationID, t.DistrictID, t.CarID, t.Number,
	)
	if err != nil {
		return fmt.Errorf("update team: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("team with number %d not found", t.Number)
	}
	return nil
}

func (r *TeamRepo) Delete(ctx context.Context, number int) error {
	query := `DELETE FROM Team WHERE number = $1`
	tag, err := r.pool.Exec(ctx, query, number)
	if err != nil {
		return fmt.Errorf("delete team: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("team with number %d not found", number)
	}
	return nil
}
