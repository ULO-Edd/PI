package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type AttendanceHandler struct {
	db *sql.DB
}

func NewAttendanceHandler(db *sql.DB) *AttendanceHandler {
	return &AttendanceHandler{db: db}
}

type AttendeeRoster struct {
	RegistrationID uuid.UUID `json:"registration_id"`
	AttendeeID     uuid.UUID `json:"attendee_id"`
	AttendeeName   string    `json:"attendee_name"`
	Status         string    `json:"status"`
}

type UpdatePresenceRequest struct {
	Status string `json:"status"`
}

func (h *AttendanceHandler) GetAttendeeRoster(w http.ResponseWriter, r *http.Request) {
    eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

    query := `
        SELECT r.id, r.attendee_id, u.username, r.status
        FROM registrations r
        JOIN users u ON r.attendee_id = u.id
        WHERE r.event_id = $1
        ORDER BY u.username ASC
    `
    rows, err := h.db.Query(query, eventID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var roster []AttendeeRoster
    for rows.Next() {
        var a AttendeeRoster
        err := rows.Scan(&a.RegistrationID, &a.AttendeeID, &a.AttendeeName, &a.Status)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        roster = append(roster, a)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(roster)
}

func (h *AttendanceHandler) UpdatePresence(w http.ResponseWriter, r *http.Request) {
    regID, err := uuid.Parse(r.PathValue("registrationId"))
	if err != nil {
		http.Error(w, "Invalid registration ID", http.StatusBadRequest)
		return
	}

    var req UpdatePresenceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Basic validation for status
    status := req.Status
    if status != "Registered" && status != "Present" && status != "No-Show" && status != "Cancelled" {
        http.Error(w, "Invalid status provided", http.StatusBadRequest)
        return
    }

    query := `UPDATE registrations SET status = $1 WHERE id = $2`
    _, err = h.db.Exec(query, status, regID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Presence status updated successfully"})
}
