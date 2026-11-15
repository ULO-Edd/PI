package handlers

import (
	"encoding/json"
	"event-api/internal/domain/registration"
	"event-api/internal/platform/auth"
	"net/http"

	"github.com/google/uuid"
)

type RegistrationHandler struct {
	service *registration.Service
}

func NewRegistrationHandler(service *registration.Service) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

func (h *RegistrationHandler) RegisterForEvent(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("userClaims").(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	attendeeID, err := uuid.Parse(claims.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	reg, err := h.service.RegisterForEvent(eventID, attendeeID)
	if err != nil {
		switch err {
		case registration.ErrEventNotPublished:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case registration.ErrEventFull, registration.ErrAlreadyRegistered:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reg)
}

func (h *RegistrationHandler) UnregisterFromEvent(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("userClaims").(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	attendeeID, err := uuid.Parse(claims.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	if err := h.service.UnregisterFromEvent(eventID, attendeeID); err != nil {
		if err == registration.ErrNotRegistered {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
