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

func setupRankRouter(repo *mocks.RankRepository) *chi.Mux {
	h := handler.NewRankHandler(repo)
	r := chi.NewRouter()
	r.Post("/api/ranks", h.Create)
	r.Get("/api/ranks", h.GetAll)
	r.Get("/api/ranks/{id}", h.GetByID)
	r.Put("/api/ranks/{id}", h.Update)
	r.Delete("/api/ranks/{id}", h.Delete)
	return r
}

func TestCreateRank_Success(t *testing.T) {
	repo := new(mocks.RankRepository)
	input := models.Rank{Name: "Капитан"}
	repo.On("Create", mock.Anything, &input).Return(1, nil)

	r := setupRankRouter(repo)
	body, _ := json.Marshal(map[string]string{"name": input.Name})
	req := httptest.NewRequest(http.MethodPost, "/api/ranks", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var resp models.Rank
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Капитан", resp.Name)
}

func TestGetAllRanks(t *testing.T) {
	repo := new(mocks.RankRepository)
	expected := []models.Rank{{ID: 1, Name: "Лейтенант"}}
	repo.On("GetAll", mock.Anything).Return(expected, nil)

	r := setupRankRouter(repo)
	req := httptest.NewRequest(http.MethodGet, "/api/ranks", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var list []models.Rank
	err := json.Unmarshal(w.Body.Bytes(), &list)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	assert.Equal(t, "Лейтенант", list[0].Name)
}
