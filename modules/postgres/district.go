package postgres

import (
	"context"
	"fmt"
	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DistrictRepo struct {
	pool *pgxpool.Pool
}

func NewDistrictRepo(pool *pgxpool.Pool) repository.DistrictRepository {
	return &DistrictRepo{pool: pool}
}

func (r *DistrictRepo) Create(ctx context.Context, d *models.District) (int, error) {
	query := `INSERT INTO District (name) VALUES ($1) RETURNING id`
	var id int
	err := r.pool.QueryRow(ctx, query, d.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create district: %w", err)
	}
	return id, nil
}

func (r *DistrictRepo) GetByID(ctx context.Context, id int) (*models.District, error) {
	query := `SELECT id, name FROM District WHERE id = $1`
	row := r.pool.QueryRow(ctx, query, id)
	var d models.District
	err := row.Scan(&d.ID, &d.Name)
	if err != nil {
		return nil, fmt.Errorf("get district by id: %w", err)
	}
	return &d, nil
}

func (r *DistrictRepo) GetAll(ctx context.Context) ([]models.District, error) {
	query := `SELECT id, name FROM District`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all districts: %w", err)
	}
	defer rows.Close()
	var districts []models.District
	for rows.Next() {
		var d models.District
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			return nil, fmt.Errorf("scan district: %w", err)
		}
		districts = append(districts, d)
	}
	return districts, rows.Err()
}

func (r *DistrictRepo) Update(ctx context.Context, d *models.District) error {
	query := `UPDATE District SET name = $1 WHERE id = $2`
	tag, err := r.pool.Exec(ctx, query, d.Name, d.ID)
	if err != nil {
		return fmt.Errorf("update district: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("district with id %d not found", d.ID)
	}
	return nil
}

func (r *DistrictRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM District WHERE id = $1`
	tag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete district: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("district with id %d not found", id)
	}
	return nil
}

func (r *DistrictRepo) AddSpecialization(ctx context.Context, districtID, specID int) error {
	query := `INSERT INTO District_specializations (district_id, specialization_id) VALUES ($1, $2)`
	_, err := r.pool.Exec(ctx, query, districtID, specID)
	if err != nil {
		return fmt.Errorf("add specialization to district: %w", err)
	}
	return nil
}

func (r *DistrictRepo) RemoveSpecialization(ctx context.Context, districtID, specID int) error {
	query := `DELETE FROM District_specializations WHERE district_id = $1 AND specialization_id = $2`
	tag, err := r.pool.Exec(ctx, query, districtID, specID)
	if err != nil {
		return fmt.Errorf("remove specialization from district: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("specialization %d not found for district %d", specID, districtID)
	}
	return nil
}

func (r *DistrictRepo) GetSpecializations(ctx context.Context, districtID int) ([]models.Specialization, error) {
	query := `SELECT s.id, s.name
	          FROM Specialization s
	          JOIN District_specializations ds ON s.id = ds.specialization_id
	          WHERE ds.district_id = $1`
	rows, err := r.pool.Query(ctx, query, districtID)
	if err != nil {
		return nil, fmt.Errorf("get district specializations: %w", err)
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
