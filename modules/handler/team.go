package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type TeamHandler struct {
	repo repository.TeamRepository
}

func NewTeamHandler(repo repository.TeamRepository) *TeamHandler {
	return &TeamHandler{repo: repo}
}

func (h *TeamHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Team
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	if input.SpecializationID == 0 || input.DistrictID == 0 {
		respondError(w, http.StatusBadRequest, "specialization_id и district_id обязательны")
		return
	}
	number, err := h.repo.Create(r.Context(), &input)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "ошибка создания: "+err.Error())
		return
	}
	input.Number = number
	respondJSON(w, http.StatusCreated, input)
}

func (h *TeamHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Фильтрация по району и специализации
	districtStr := r.URL.Query().Get("district")
	specStr := r.URL.Query().Get("specialization")

	if districtStr != "" && specStr != "" {
		respondError(w, http.StatusBadRequest, "можно указать только один фильтр")
		return
	}

	if districtStr != "" {
		districtID, err := strconv.Atoi(districtStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный district_id")
			return
		}
		list, err := h.repo.GetByDistrict(r.Context(), districtID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
		return
	}

	if specStr != "" {
		specID, err := strconv.Atoi(specStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный specialization_id")
			return
		}
		list, err := h.repo.GetBySpecialization(r.Context(), specID)
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

func (h *TeamHandler) GetByNumber(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный номер команды")
		return
	}
	team, err := h.repo.GetByNumber(r.Context(), number)
	if err != nil {
		respondError(w, http.StatusNotFound, "команда не найдена")
		return
	}
	respondJSON(w, http.StatusOK, team)
}

func (h *TeamHandler) Update(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный номер команды")
		return
	}
	var input models.Team
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	input.Number = number
	if err := h.repo.Update(r.Context(), &input); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, input)
}

func (h *TeamHandler) Delete(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(chi.URLParam(r, "number"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный номер команды")
		return
	}
	if err := h.repo.Delete(r.Context(), number); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
