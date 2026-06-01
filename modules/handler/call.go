package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"fire-department/modules/models"
	"fire-department/modules/repository"

	"github.com/go-chi/chi/v5"
)

type CallHandler struct {
	repo repository.CallRepository
}

func NewCallHandler(repo repository.CallRepository) *CallHandler {
	return &CallHandler{repo: repo}
}

func (h *CallHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Call
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
		return
	}
	if input.TeamID == 0 || input.CarID == 0 || input.StatusID == 0 {
		respondError(w, http.StatusBadRequest, "team_id, car_id и status_id обязательны")
		return
	}
	if input.Time.IsZero() {
		respondError(w, http.StatusBadRequest, "time обязателен")
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

func (h *CallHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	teamStr := r.URL.Query().Get("team")
	carStr := r.URL.Query().Get("car")
	districtStr := r.URL.Query().Get("district")
	statusStr := r.URL.Query().Get("status")

	// Проверяем, что передан только один фильтр (для простоты)
	filtersCount := 0
	if teamStr != "" { filtersCount++ }
	if carStr != "" { filtersCount++ }
	if districtStr != "" { filtersCount++ }
	if statusStr != "" { filtersCount++ }
	if filtersCount > 1 {
		respondError(w, http.StatusBadRequest, "можно указать только один фильтр")
		return
	}

	switch {
	case teamStr != "":
		teamID, err := strconv.Atoi(teamStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный team_id")
			return
		}
		list, err := h.repo.GetByTeam(r.Context(), teamID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
	case carStr != "":
		carID, err := strconv.Atoi(carStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный car_id")
			return
		}
		list, err := h.repo.GetByCar(r.Context(), carID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
	case districtStr != "":
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
	case statusStr != "":
		statusID, err := strconv.Atoi(statusStr)
		if err != nil {
			respondError(w, http.StatusBadRequest, "некорректный status_id")
			return
		}
		list, err := h.repo.GetByStatus(r.Context(), statusID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
	default:
		list, err := h.repo.GetAll(r.Context())
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondJSON(w, http.StatusOK, list)
	}
}

func (h *CallHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	call, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "вызов не найден")
		return
	}
	respondJSON(w, http.StatusOK, call)
}

func (h *CallHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "некорректный id")
		return
	}
	var input models.Call
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

func (h *CallHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
