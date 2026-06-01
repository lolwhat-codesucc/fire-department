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

func setupDistrictRouter(repo *mocks.DistrictRepository) *chi.Mux {
	h := handler.NewDistrictHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/districts", h.Create)
	r.Get("/api/districts", h.GetAll)
	r.Get("/api/districts/{id}", h.GetByID)
	r.Put("/api/districts/{id}", h.Update)
	r.Delete("/api/districts/{id}", h.Delete)
	r.Post("/api/districts/{id}/specializations/{specId}", h.AddSpecialization)
	r.Delete("/api/districts/{id}/specializations/{specId}", h.RemoveSpecialization)
	r.Get("/api/districts/{id}/specializations", h.GetSpecializations)
	return r
}

func TestCreateDistrict_Success(t *testing.T) {
	repo := new(mocks.DistrictRepository)
	input := models.District{Name: "Центральный"}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupDistrictRouter(repo)
	body, _ := json.Marshal(map[string]string{"name": input.Name})
	req := httptest.NewRequest(http.MethodPost, "/api/districts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAddSpecializationToDistrict(t *testing.T) {
	repo := new(mocks.DistrictRepository)
	repo.On("AddSpecialization", mock.Anything, 1, 2).Return(nil)

	r := setupDistrictRouter(repo)
	req := httptest.NewRequest(http.MethodPost, "/api/districts/1/specializations/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	repo.AssertCalled(t, "AddSpecialization", mock.Anything, 1, 2)
}

func TestGetSpecializationsOfDistrict(t *testing.T) {
	repo := new(mocks.DistrictRepository)
	expected := []models.Specialization{{ID: 2, Name: "Спасательная"}}
	repo.On("GetSpecializations", mock.Anything, 1).Return(expected, nil)

	r := setupDistrictRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/districts/1/specializations", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var specs []models.Specialization
	json.Unmarshal(w.Body.Bytes(), &specs)
	assert.Len(t, specs, 1)
}
