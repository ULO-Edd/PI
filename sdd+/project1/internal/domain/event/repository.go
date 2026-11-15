package event

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type dbRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &dbRepository{db: db}
}

func (r *dbRepository) Create(e *Event) error {
	query := `INSERT INTO events (id, title, description, location, start_time, end_time, organiser_id, capacity, status, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.Exec(query, e.ID, e.Title, e.Description, e.Location, e.StartTime, e.EndTime, e.OrganiserID, e.Capacity, e.Status, e.CreatedAt)
	return err
}

func (r *dbRepository) FindAll(status, location, organiserID string) ([]*Event, error) {
	query := "SELECT id, title, description, location, start_time, end_time, organiser_id, capacity, status, created_at FROM events WHERE 1=1"
	var args []interface{}
	argIdx := 1

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, status)
		argIdx++
	}
	if location != "" {
		query += fmt.Sprintf(" AND location ILIKE $%d", argIdx)
		args = append(args, "%"+location+"%")
		argIdx++
	}
	if organiserID != "" {
		query += fmt.Sprintf(" AND organiser_id = $%d", argIdx)
		args = append(args, organiserID)
		argIdx++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		e := &Event{}
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.StartTime, &e.EndTime, &e.OrganiserID, &e.Capacity, &e.Status, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *dbRepository) FindByID(id uuid.UUID) (*Event, error) {
	e := &Event{}
	query := `SELECT id, title, description, location, start_time, end_time, organiser_id, capacity, status, created_at FROM events WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.StartTime, &e.EndTime, &e.OrganiserID, &e.Capacity, &e.Status, &e.CreatedAt)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *dbRepository) Update(e *Event) error {
	query := `UPDATE events SET title = $1, description = $2, location = $3, start_time = $4, end_time = $5, capacity = $6, status = $7 WHERE id = $8`
	_, err := r.db.Exec(query, e.Title, e.Description, e.Location, e.StartTime, e.EndTime, e.Capacity, e.Status, e.ID)
	return err
}

func (r *dbRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
