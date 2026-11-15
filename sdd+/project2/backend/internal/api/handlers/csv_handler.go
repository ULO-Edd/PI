package handlers

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type CsvHandler struct {
	db *sql.DB
}

func NewCsvHandler(db *sql.DB) *CsvHandler {
	return &CsvHandler{db: db}
}

func (h *CsvHandler) ExportEventAttendees(w http.ResponseWriter, r *http.Request) {
    eventID, err := uuid.Parse(r.PathValue("eventId"))
	if err != nil {
		http.Error(w, "Invalid event ID", http.StatusBadRequest)
		return
	}

    query := `
        SELECT u.username, u.email, r.status, r.created_at
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

    w.Header().Set("Content-Type", "text/csv")
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"attendees_%s.csv\"", eventID.String()))

    csvWriter := csv.NewWriter(w)
    // Write CSV header
    if err := csvWriter.Write([]string{"Username", "Email", "Status", "RegisteredAt"}); err != nil {
        http.Error(w, "Failed to write CSV header", http.StatusInternalServerError)
        return
    }

    for rows.Next() {
        var username, email, status, registeredAt string
        if err := rows.Scan(&username, &email, &status, &registeredAt); err != nil {
            http.Error(w, "Failed to scan row", http.StatusInternalServerError)
            return
        }
        if err := csvWriter.Write([]string{username, email, status, registeredAt}); err != nil {
            http.Error(w, "Failed to write CSV row", http.StatusInternalServerError)
            return
        }
    }
    csvWriter.Flush()
}
