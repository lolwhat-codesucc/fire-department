package handler_test

import (
	"bytes"
	"encoding/json"
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

func setupCallRouter(repo *mocks.CallRepository) *chi.Mux {
	h := handler.NewCallHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/calls", h.Create)
	r.Get("/api/calls", h.GetAll)
	r.Get("/api/calls/{id}", h.GetByID)
	r.Put("/api/calls/{id}", h.Update)
	r.Delete("/api/calls/{id}", h.Delete)
	return r
}

func TestCreateCall_Success(t *testing.T) {
	repo := new(mocks.CallRepository)
	now := time.Now().Truncate(time.Second)
	input := models.Call{
		TeamID:   1,
		CarID:    2,
		Time:     now,
		StatusID: 1,
	}
	repo.On("Create", mock.Anything, &input).Return(7, nil)

	r := setupCallRouter(repo)
	body, _ := json.Marshal(map[string]any{
		"team_id":   input.TeamID,
		"car_id":    input.CarID,
		"time":      input.Time.Format(time.RFC3339),
		"status_id": input.StatusID,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/calls", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Call
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 7, resp.ID)
}

func TestCreateCall_MissingFields(t *testing.T) {
	repo := new(mocks.CallRepository)
	r := setupCallRouter(repo)
	body := `{"team_id":1}` 
	req := httptest.NewRequest(http.MethodPost, "/api/calls", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCallsMultipleFilters(t *testing.T) {
	repo := new(mocks.CallRepository)
	r := setupCallRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/calls?team=1&car=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "только один фильтр")
}
