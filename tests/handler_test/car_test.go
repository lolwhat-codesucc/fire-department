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

func setupCarRouter(repo *mocks.CarRepository) *chi.Mux {
	h := handler.NewCarHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/cars", h.Create)
	r.Get("/api/cars", h.GetAll)
	r.Get("/api/cars/{id}", h.GetByID)
	r.Put("/api/cars/{id}", h.Update)
	r.Delete("/api/cars/{id}", h.Delete)
	return r
}

func TestCreateCar_Success(t *testing.T) {
	repo := new(mocks.CarRepository)
	acqDate := models.Date{Time: time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC)}
	input := models.Car{
		ModelID:         1,
		AcquisitionDate: acqDate,
		Ready:           true,
	}
	repo.On("Create", mock.Anything, &input).Return(5, nil)

	r := setupCarRouter(repo)
	body, _ := json.Marshal(map[string]interface{}{
		"model_id":         input.ModelID,
		"acquisition_date": "2023-01-15",
		"ready":            input.Ready,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/cars", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Car
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 5, resp.ID)
}

func TestGetCarsByModelFilter(t *testing.T) {
	repo := new(mocks.CarRepository)
	repo.On("GetByModel", mock.Anything, 2).Return([]models.Car{{ID: 3}}, nil)

	r := setupCarRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/cars?model=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
