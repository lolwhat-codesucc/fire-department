package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type RankHandler struct {
	repo repository.RankRepository
}

func NewRankHandler(repo repository.RankRepository) *RankHandler {
	return &RankHandler{repo: repo}
}

func (h *RankHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Rank
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	if input.Name == "" {
		respondError(w, http.StatusBadRequest, "name обязателен")
		return
	}
	id, err := h.repo.Create(r.Context(), &input)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "ошибка создания: "+err.Error())
		return
	}
	input.ID = id
	respondJSON(w, http.StatusCreated, input)
}

func (h *RankHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

func (h *RankHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	rank, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "звание не найдено")
		return
	}
	respondJSON(w, http.StatusOK, rank)
}

func (h *RankHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	var input models.Rank
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	input.ID = id
	if err := h.repo.Update(r.Context(), &input); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, input)
}

func (h *RankHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	if err := h.repo.Delete(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
