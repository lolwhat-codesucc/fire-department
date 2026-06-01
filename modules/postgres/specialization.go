package postgres

import (
	"context"
	"fire-department/modules/models"
	"fire-department/modules/repository"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SpecializationRepo struct {
	pool *pgxpool.Pool
}

func NewSpecializationRepo(pool *pgxpool.Pool) repository.SpecializationRepository {
	return &SpecializationRepo{pool: pool}
}

func (r *SpecializationRepo) Create(ctx context.Context, s *models.Specialization) (int, error) {
	query := `INSERT INTO Specialization (name) VALUES ($1) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query, s.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create specialization: %w", err)
	}
	return id, nil
}

func (r *SpecializationRepo) GetByID(ctx context.Context, id int) (*models.Specialization, error) {
	query := `SELECT id, name FROM Specialization WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var s models.Specialization
	err := row.Scan(&s.ID, &s.Name)
	if err != nil {
		return nil, fmt.Errorf("get specialization by id: %w", err)
	}
	return &s, nil
}

func (r *SpecializationRepo) GetAll(ctx context.Context) ([]models.Specialization, error) {
	query := `SELECT id, name FROM Specialization`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all specializations: %w", err)
	}
	defer rows.Close()
	var specs []models.Specialization
	for rows.Next() {
		var s models.Specialization
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			return nil, fmt.Errorf("scan specialization: %w", err)
		}
		specs = append(specs, s)
	}
	return specs, rows.Err()
}

func (r *SpecializationRepo) Update(ctx context.Context, s *models.Specialization) error {
	query := `UPDATE Specialization SET name = $1 WHERE id = $2`
	tag, err := r.pool.Exec(ctx, query, s.Name, s.ID)
	if err != nil {
		return fmt.Errorf("update specialization: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("specialization with id %d not found", s.ID)
	}
	return nil
}

func (r *SpecializationRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM Specialization WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete specialization: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("specialization with id %d not found", id)
	}
	return nil
}
