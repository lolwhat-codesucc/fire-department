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

func setupCallStatusRouter(repo *mocks.CallStatusRepository) *chi.Mux {
	h := handler.NewCallStatusHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/call-statuses", h.Create)
	r.Get("/api/call-statuses", h.GetAll)
	r.Get("/api/call-statuses/{id}", h.GetByID)
	r.Put("/api/call-statuses/{id}", h.Update)
	r.Delete("/api/call-statuses/{id}", h.Delete)
	return r
}

func TestCreateCallStatus_Success(t *testing.T) {
	repo := new(mocks.CallStatusRepository)
	input := models.CallStatus{Name: "Поступил"}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupCallStatusRouter(repo)
	body, _ := json.Marshal(map[string]string{"name": input.Name})
	req := httptest.NewRequest(http.MethodPost, "/api/call-statuses", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.CallStatus
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Поступил", resp.Name)
}
