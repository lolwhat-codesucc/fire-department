package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "fire-department/modules/models"
    "fire-department/modules/repository"

    "github.com/go-chi/chi/v5"
)

type FirefighterHandler struct {
    repo repository.FirefighterRepository
}

func NewFirefighterHandler(repo repository.FirefighterRepository) *FirefighterHandler {
    return &FirefighterHandler{repo: repo}
}

func (h *FirefighterHandler) Create(w http.ResponseWriter, r *http.Request) {
    var input models.Firefighter
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondError(w, http.StatusBadRequest, "некорректный JSON: "+err.Error())
        return
    }
    if input.Name == "" || input.RankID == 0 {
        respondError(w, http.StatusBadRequest, "name и rank_id обязательны")
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

func (h *FirefighterHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        respondError(w, http.StatusBadRequest, "некорректный id")
        return
    }
    ff, err := h.repo.GetByID(r.Context(), id)
    if err != nil {
        respondError(w, http.StatusNotFound, "пожарный не найден")
        return
    }
    respondJSON(w, http.StatusOK, ff)
}

func (h *FirefighterHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    teamStr := r.URL.Query().Get("team")
    if teamStr != "" {
        teamNumber, err := strconv.Atoi(teamStr)
        if err != nil {
            respondError(w, http.StatusBadRequest, "некорректный номер команды")
            return
        }
        list, err := h.repo.GetByTeam(r.Context(), teamNumber)
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

func (h *FirefighterHandler) Update(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        respondError(w, http.StatusBadRequest, "некорректный id")
        return
    }
    var input models.Firefighter
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

func (h *FirefighterHandler) Delete(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
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
