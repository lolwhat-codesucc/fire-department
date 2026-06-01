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

func setupCarModelRouter(repo *mocks.CarModelRepository) *chi.Mux {
	h := handler.NewCarModelHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/car-models", h.Create)
	r.Get("/api/car-models", h.GetAll)
	r.Get("/api/car-models/{id}", h.GetByID)
	r.Put("/api/car-models/{id}", h.Update)
	r.Delete("/api/car-models/{id}", h.Delete)
	return r
}

func TestCreateCarModel_Success(t *testing.T) {
	repo := new(mocks.CarModelRepository)
	input := models.CarModel{Name: "АЦ-40", MaintenancePeriodDays: 180}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupCarModelRouter(repo)
	body, _ := json.Marshal(map[string]any{
		"name":                   input.Name,
		"maintenance_period_days": input.MaintenancePeriodDays,
	})
	req := httptest.NewRequest(http.MethodPost, "/api/car-models", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.CarModel
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "АЦ-40", resp.Name)
}
