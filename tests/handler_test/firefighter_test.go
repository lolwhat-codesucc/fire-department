package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"fire-department/modules/handler"
	"fire-department/modules/models"
	"fire-department/tests/mocks"
)

func setupFirefighterRouter(repo *mocks.FirefighterRepository) *chi.Mux {
	h := handler.NewFirefighterHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/firefighters", h.Create)
	r.Get("/api/firefighters", h.GetAll)
	r.Get("/api/firefighters/{id}", h.GetByID)
	r.Put("/api/firefighters/{id}", h.Update)
	r.Delete("/api/firefighters/{id}", h.Delete)
	return r
}

func TestCreateFirefighter_Success(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	input := models.Firefighter{
		Name:          "Иванов",
		YearOfBirth:   models.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
		RankID:        1,
		Qualification: 3,
	}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupFirefighterRouter(repo)
	body, _ := json.Marshal(map[string]any{
		"name":          input.Name,
		"year_of_birth": "1990-01-01",
		"rank_id":       input.RankID,
		"qualification": input.Qualification,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/firefighters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Firefighter
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Иванов", resp.Name)
	repo.AssertExpectations(t)
}

func TestCreateFirefighter_MissingFields(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	r := setupFirefighterRouter(repo)
	body := `{"year_of_birth":"1990-01-01","rank_id":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/firefighters", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "name")
}

func TestGetFirefighterByID_NotFound(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	repo.On("GetByID", mock.Anything, 999).Return(nil, errors.New("not found"))
	r := setupFirefighterRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/firefighters/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetAllFirefighters_WithTeamFilter(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	expected := []models.Firefighter{{ID: 1, Name: "Петров", TeamID: new(5)}}
	repo.On("GetByTeam", mock.Anything, 5).Return(expected, nil)

	r := setupFirefighterRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/firefighters?team=5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var list []models.Firefighter
	err := json.Unmarshal(w.Body.Bytes(), &list)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	assert.Equal(t, "Петров", list[0].Name)
}

func TestUpdateFirefighter_Success(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	updated := models.Firefighter{
		ID:            2,
		Name:          "Сидоров",
		YearOfBirth:   models.Date{Time: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)},
		RankID:        2,
		Qualification: 4,
	}
	repo.On("Update", mock.Anything, &updated).Return(nil)

	r := setupFirefighterRouter(repo)
	body, _ := json.Marshal(map[string]any{
		"name":          updated.Name,
		"year_of_birth": "1990-01-01",
		"rank_id":       updated.RankID,
		"qualification": updated.Qualification,
	})
	req := httptest.NewRequest(http.MethodPut, "/api/firefighters/2", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFirefighter_Success(t *testing.T) {
	repo := new(mocks.FirefighterRepository)
	repo.On("Delete", mock.Anything, 1).Return(nil)
	r := setupFirefighterRouter(repo)
	req := httptest.NewRequest(http.MethodDelete, "/api/firefighters/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
