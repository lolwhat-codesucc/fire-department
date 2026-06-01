package models

import (
	"time"
)

type Rank struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Specialization struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type District struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type CarModel struct {
	ID                   int    `db:"id" json:"id"`
	Name                 string `db:"name" json:"name"`
	MaintenancePeriodDays int   `db:"maintenance_period_days" json:"maintenance_period_days"`
}

type CallStatus struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Car struct {
	ID               int       `db:"id" json:"id"`
	ModelID          int       `db:"model_id" json:"model_id"`
	AcquisitionDate  Date `db:"acquisition_date" json:"acquisition_date"`
	LastMaintenance  *Date `db:"last_maintenance" json:"last_maintenance"` // nullable
	Ready            bool      `db:"ready" json:"ready"`
}

type Team struct {
	Number           int    `db:"number" json:"number"`
	SpecializationID int    `db:"specialization_id" json:"specialization_id"`
	DistrictID       int    `db:"district_id" json:"district_id"`
	CarID            *int   `db:"car_id" json:"car_id"` // nullable + unique
}

type Firefighter struct {
	ID            int       `db:"id" json:"id"`
	Name          string    `db:"name" json:"name"`
	YearOfBirth   Date `db:"year_of_birth" json:"year_of_birth"` // date
	RankID        int       `db:"rank_id" json:"rank_id"`
	Qualification int       `db:"qualification" json:"qualification"`
	TeamID        *int      `db:"team_id" json:"team_id"` // nullable
}

type Call struct {
	ID         int       `db:"id" json:"id"`
	TeamID     int       `db:"team_id" json:"team_id"`
	DistrictID *int      `db:"district_id" json:"district_id"` // nullable
	CarID      int       `db:"car_id" json:"car_id"`
	Time       time.Time `db:"time" json:"time"`
	StatusID   int       `db:"status_id" json:"status_id"`
	Comment    *string   `db:"comment" json:"comment"` // nullable
}

type DistrictSpecialization struct {
	DistrictID      int `db:"district_id" json:"district_id"`
	SpecializationID int `db:"specialization_id" json:"specialization_id"`
}
