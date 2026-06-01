package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type CarHandler struct {
	repo repository.CarRepository
}

func NewCarHandler(repo repository.CarRepository) *CarHandler {
	return &CarHandler{repo: repo}
}

func (h *CarHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Car
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	if input.ModelID == 0 {
		respondError(w, http.StatusBadRequest, "model_id обязателен")
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

func (h *CarHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	modelStr := r.URL.Query().Get("model")
	if modelStr != "" {
		modelID, err := strconv.Atoi(modelStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный model_id")
			return
		}
		list, err := h.repo.GetByModel(r.Context(), modelID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
		return
	}
	list, err := h.repo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

func (h *CarHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	car, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "автомобиль не найден")
		return
	}
	respondJSON(w, http.StatusOK, car)
}

func (h *CarHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	var input models.Car
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

func (h *CarHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
