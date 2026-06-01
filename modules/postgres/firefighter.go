package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FirefighterRepo struct {
	pool *pgxpool.Pool
}

func NewFirefighterRepo(pool *pgxpool.Pool) repository.FirefighterRepository {
	return &FirefighterRepo{pool: pool}
}

func (r *FirefighterRepo) Create(ctx context.Context, f *models.Firefighter) (int, error) {
	query := `INSERT INTO Firefighter (name, year_of_birth, rank_id, qualification, team_id)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query,
		f.Name, f.YearOfBirth, f.RankID, f.Qualification, f.TeamID,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create firefighter: %w", err)
	}
	return id, nil
}

func (r *FirefighterRepo) GetByID(ctx context.Context, id int) (*models.Firefighter, error) {
	query := `SELECT id, name, year_of_birth, rank_id, qualification, team_id
	          FROM Firefighter WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var ff models.Firefighter
	err := row.Scan(&ff.ID, &ff.Name, &ff.YearOfBirth, &ff.RankID, &ff.Qualification, &ff.TeamID)
	if err != nil {
		return nil, fmt.Errorf("get firefighter by id: %w", err)
	}
	return &ff, nil
}

func (r *FirefighterRepo) GetAll(ctx context.Context) ([]models.Firefighter, error) {
	query := `SELECT id, name, year_of_birth, rank_id, qualification, team_id FROM Firefighter`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all firefighters: %w", err)
	}
	defer rows.Close()
	var result []models.Firefighter
	for rows.Next() {
		var ff models.Firefighter
		if err := rows.Scan(&ff.ID, &ff.Name, &ff.YearOfBirth, &ff.RankID, &ff.Qualification, &ff.TeamID); err != nil {
			return nil, fmt.Errorf("scan firefighter: %w", err)
		}
		result = append(result, ff)
	}
	return result, rows.Err()
}

func (r *FirefighterRepo) GetByTeam(ctx context.Context, teamNumber int) ([]models.Firefighter, error) {
	query := `SELECT id, name, year_of_birth, rank_id, qualification, team_id FROM Firefighter WHERE team_id = $1`
	rows, err := r.pool.Query(ctx, query, teamNumber)
	if err != nil {
		return nil, fmt.Errorf("get firefighters by team: %w", err)
	}
	defer rows.Close()
	var result []models.Firefighter
	for rows.Next() {
		var ff models.Firefighter
		if err := rows.Scan(&ff.ID, &ff.Name, &ff.YearOfBirth, &ff.RankID, &ff.Qualification, &ff.TeamID); err != nil {
			return nil, err
		}
		result = append(result, ff)
	}
	return result, rows.Err()
}

func (r *FirefighterRepo) Update(ctx context.Context, f *models.Firefighter) error {
	query := `UPDATE Firefighter SET name=$1, year_of_birth=$2, rank_id=$3, qualification=$4, team_id=$5 WHERE id=$6`
	tag, err := r.pool.Exec(ctx, query,
		f.Name, f.YearOfBirth, f.RankID, f.Qualification, f.TeamID, f.ID,
	)
	if err != nil {
		return fmt.Errorf("update firefighter: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("firefighter with id %d not found", f.ID)
	}
	return nil
}

func (r *FirefighterRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Firefighter WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete firefighter: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("firefighter with id %d not found", id)
	}
	return nil
}
