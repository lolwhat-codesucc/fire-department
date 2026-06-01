package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"fire-department/modules/handler"
	"fire-department/modules/models"
	"fire-department/tests/mocks"
)

func setupTeamRouter(repo *mocks.TeamRepository) *chi.Mux {
	h := handler.NewTeamHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/teams", h.Create)
	r.Get("/api/teams", h.GetAll)
	r.Get("/api/teams/{number}", h.GetByNumber)
	r.Put("/api/teams/{number}", h.Update)
	r.Delete("/api/teams/{number}", h.Delete)
	return r
}

func TestCreateTeam_Success(t *testing.T) {
	repo := new(mocks.TeamRepository)
	input := models.Team{SpecializationID: 1, DistrictID: 2}
	repo.On("Create", mock.Anything, &input).Return(10, nil)

	r := setupTeamRouter(repo)
	body, _ := json.Marshal(map[string]int{
		"specialization_id": input.SpecializationID,
		"district_id":       input.DistrictID,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/teams", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Team
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 10, resp.Number)
}

func TestGetTeamByNumber_NotFound(t *testing.T) {
	repo := new(mocks.TeamRepository)
	repo.On("GetByNumber", mock.Anything, 99).Return(nil, errors.New("not found"))
	r := setupTeamRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/teams/99", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetTeamsByDistrict(t *testing.T) {
	repo := new(mocks.TeamRepository)
	repo.On("GetByDistrict", mock.Anything, 3).Return([]models.Team{{Number: 4}}, nil)

	r := setupTeamRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/teams?district=3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
