package handlers

import (
	"encoding/json"
	"event-api/internal/domain/event"
	"event-api/internal/platform/auth"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type EventHandler struct {
	service *event.Service
}

func NewEventHandler(service *event.Service) *EventHandler {
	return &EventHandler{service: service}
}

type CreateEventRequest struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Capacity     int       `json:"capacity"`
	CategoryID   uuid.UUID `json:"category_id"`
	ContactEmail string    `json:"contact_email"`
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event (requires organiser or admin role).
// @Tags events
// @Accept  json
// @Produce  json
// @Param   event body      CreateEventRequest true  "Event info"
// @Success 201   {object}  event.Event
// @Security ApiKeyAuth
// @Router /events [post]
func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("userClaims").(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	organiserID, err := uuid.Parse(claims.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	var req CreateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	e, err := h.service.CreateEvent(req.Title, req.Description, req.Location, req.StartTime, req.EndTime, organiserID, req.CategoryID, req.Capacity, req.ContactEmail)
	if err != nil {
		switch err {
		case event.ErrInvalidTitle, event.ErrInvalidTimes, event.ErrStartTimePast, event.ErrInvalidCapacity, event.ErrInvalidStatus:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

// GetEvents godoc
// @Summary Get all events
// @Description Get a list of all events.
// @Tags events
// @Produce  json
// @Success 200 {array} event.Event
// @Router /events [get]
func (h *EventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	location := r.URL.Query().Get("location")
	organiserID := r.URL.Query().Get("organiser_id")

	events, err := h.service.GetAllEvents(status, location, organiserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// GetEventByID godoc
// @Summary Get a single event by its ID
// @Description Get a single event by its ID.
// @Tags events
// @Produce  json
// @Param   eventId     path      string  true  "Event ID"
// @Success 200 {object} event.Event
// @Router /events/{eventId} [get]
func (h *EventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	e, err := h.service.GetEventByID(eventID)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

type UpdateEventRequest struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Capacity     int       `json:"capacity"`
	Status       string    `json:"status"`
	CategoryID   uuid.UUID `json:"category_id"`
	ContactEmail string    `json:"contact_email"`
}

// UpdateEvent godoc
// @Summary Update an existing event
// @Description Update an existing event. Requires organiser or admin role. User must own the event.
// @Tags events
// @Accept  json
// @Produce  json
// @Param   eventId     path      string  true  "Event ID"
// @Param   event body      UpdateEventRequest true  "Event info"
// @Success 200   {object}  event.Event
// @Security ApiKeyAuth
// @Router /events/{eventId} [put]
func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("userClaims").(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	organiserID, err := uuid.Parse(claims.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	var req UpdateEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e, err := h.service.GetEventByID(eventID)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Verify the user is the owner or an admin
	if claims.Role != "administrator" && e.OrganiserID != organiserID {
		http.Error(w, "Forbidden: you do not own this event", http.StatusForbidden)
		return
	}
	
	e.Title = req.Title
	e.Description = req.Description
	e.Location = req.Location
	e.StartTime = req.StartTime
	e.EndTime = req.EndTime
	e.Capacity = req.Capacity
	e.Status = req.Status
	e.CategoryID = req.CategoryID
	e.ContactEmail = req.ContactEmail

	if err := h.service.UpdateEvent(e); err != nil {
		switch err {
		case event.ErrInvalidTitle, event.ErrInvalidTimes, event.ErrStartTimePast, event.ErrInvalidCapacity, event.ErrInvalidStatus:
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an event. Requires organiser or admin role. User must own the event.
// @Tags events
// @Param   eventId     path      string  true  "Event ID"
// @Success 204
// @Security ApiKeyAuth
// @Router /events/{eventId} [delete]
func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("userClaims").(*auth.Claims)
	if !ok {
		http.Error(w, "Could not retrieve user claims", http.StatusInternalServerError)
		return
	}

	organiserID, err := uuid.Parse(claims.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

	e, err := h.service.GetEventByID(eventID)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	// Verify user owns this event or is admin
	if claims.Role != "administrator" && e.OrganiserID != organiserID {
		http.Error(w, "Forbidden: you do not own this event", http.StatusForbidden)
		return
	}

	if err := h.service.DeleteEvent(eventID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
