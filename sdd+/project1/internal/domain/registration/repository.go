package registration

import (
	"database/sql"
	"github.com/google/uuid"
)

type dbRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &dbRepository{db: db}
}

func (r *dbRepository) Create(reg *Registration) error {
	query := `INSERT INTO registrations (id, event_id, attendee_id, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, reg.ID, reg.EventID, reg.AttendeeID, reg.CreatedAt)
	return err
}

func (r *dbRepository) Delete(eventID, attendeeID uuid.UUID) error {
	query := `DELETE FROM registrations WHERE event_id = $1 AND attendee_id = $2`
	_, err := r.db.Exec(query, eventID, attendeeID)
	return err
}

func (r *dbRepository) CountByEventID(eventID uuid.UUID) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM registrations WHERE event_id = $1`
	err := r.db.QueryRow(query, eventID).Scan(&count)
	return count, err
}

func (r *dbRepository) FindByEventAndAttendee(eventID, attendeeID uuid.UUID) (*Registration, error) {
	reg := &Registration{}
	query := `SELECT id, event_id, attendee_id, created_at FROM registrations WHERE event_id = $1 AND attendee_id = $2`
	err := r.db.QueryRow(query, eventID, attendeeID).Scan(&reg.ID, &reg.EventID, &reg.AttendeeID, &reg.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found is not an error here
		}
		return nil, err
	}
	return reg, nil
}

