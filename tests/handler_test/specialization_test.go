package handler_test

import (
	"bytes"
	"encoding/json"
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

func setupSpecializationRouter(repo *mocks.SpecializationRepository) *chi.Mux {
	h := handler.NewSpecializationHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/specializations", h.Create)
	r.Get("/api/specializations", h.GetAll)
	r.Get("/api/specializations/{id}", h.GetByID)
	r.Put("/api/specializations/{id}", h.Update)
	r.Delete("/api/specializations/{id}", h.Delete)
	return r
}

func TestCreateSpecialization_Success(t *testing.T) {
	repo := new(mocks.SpecializationRepository)
	input := models.Specialization{Name: "Спасательная"}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupSpecializationRouter(repo)
	body, _ := json.Marshal(map[string]string{"name": input.Name})
	req := httptest.NewRequest(http.MethodPost, "/api/specializations", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Specialization
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Спасательная", resp.Name)
}

func TestGetAllSpecializations(t *testing.T) {
	repo := new(mocks.SpecializationRepository)
	expected := []models.Specialization{{ID: 1, Name: "Пожарная"}}
	repo.On("GetAll", mock.Anything).Return(expected, nil)

	r := setupSpecializationRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/specializations", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
