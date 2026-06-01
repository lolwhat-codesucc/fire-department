package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RankRepo struct {
	pool *pgxpool.Pool
}

func NewRankRepo(pool *pgxpool.Pool) repository.RankRepository {
	return &RankRepo{pool: pool}
}

func (r *RankRepo) Create(ctx context.Context, rank *models.Rank) (int, error) {
	query := `INSERT INTO Rank (name) VALUES ($1) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query, rank.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create rank: %w", err)
	}
	return id, nil
}

func (r *RankRepo) GetByID(ctx context.Context, id int) (*models.Rank, error) {
	query := `SELECT id, name FROM Rank WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var rank models.Rank
	err := row.Scan(&rank.ID, &rank.Name)
	if err != nil {
		return nil, fmt.Errorf("get rank by id: %w", err)
	}
	return &rank, nil
}

func (r *RankRepo) GetAll(ctx context.Context) ([]models.Rank, error) {
	query := `SELECT id, name FROM Rank`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all ranks: %w", err)
	}
	defer rows.Close()
	var ranks []models.Rank
	for rows.Next() {
		var rank models.Rank
		if err := rows.Scan(&rank.ID, &rank.Name); err != nil {
			return nil, fmt.Errorf("scan rank: %w", err)
		}
		ranks = append(ranks, rank)
	}
	return ranks, rows.Err()
}

func (r *RankRepo) Update(ctx context.Context, rank *models.Rank) error {
	query := `UPDATE Rank SET name = $1 WHERE id = $2`
	tag, err := r.pool.Exec(ctx, query, rank.Name, rank.ID)
	if err != nil {
		return fmt.Errorf("update rank: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("rank with id %d not found", rank.ID)
	}
	return nil
}

func (r *RankRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Rank WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete rank: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("rank with id %d not found", id)
	}
	return nil
}
