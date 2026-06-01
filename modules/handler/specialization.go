package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type SpecializationHandler struct {
	repo repository.SpecializationRepository
}

func NewSpecializationHandler(repo repository.SpecializationRepository) *SpecializationHandler {
	return &SpecializationHandler{repo: repo}
}

func (h *SpecializationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Specialization
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

func (h *SpecializationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

func (h *SpecializationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	spec, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "специализация не найдена")
		return
	}
	respondJSON(w, http.StatusOK, spec)
}

func (h *SpecializationHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	var input models.Specialization
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

func (h *SpecializationHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
