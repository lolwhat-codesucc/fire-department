package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type DistrictHandler struct {
	repo repository.DistrictRepository
}

func NewDistrictHandler(repo repository.DistrictRepository) *DistrictHandler {
	return &DistrictHandler{repo: repo}
}

func (h *DistrictHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.District
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

func (h *DistrictHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.GetAll(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}

func (h *DistrictHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	district, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "район не найден")
		return
	}
	respondJSON(w, http.StatusOK, district)
}

func (h *DistrictHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	var input models.District
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

func (h *DistrictHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *DistrictHandler) AddSpecialization(w http.ResponseWriter, r *http.Request) {
	districtID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id района")
		return
	}
	specID, err := strconv.Atoi(chi.URLParam(r, "specId"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный specId")
		return
	}
	if err := h.repo.AddSpecialization(r.Context(), districtID, specID); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *DistrictHandler) RemoveSpecialization(w http.ResponseWriter, r *http.Request) {
	districtID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id района")
		return
	}
	specID, err := strconv.Atoi(chi.URLParam(r, "specId"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный specId")
		return
	}
	if err := h.repo.RemoveSpecialization(r.Context(), districtID, specID); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *DistrictHandler) GetSpecializations(w http.ResponseWriter, r *http.Request) {
	districtID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id района")
		return
	}
	specs, err := h.repo.GetSpecializations(r.Context(), districtID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, specs)
}
