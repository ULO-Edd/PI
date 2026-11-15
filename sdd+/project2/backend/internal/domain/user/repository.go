package user

import (
	"database/sql"

	"github.com/google/uuid"
)

// dbRepository is the concrete implementation of the Repository interface for a SQL database.
type dbRepository struct {
	db *sql.DB
}

// NewRepository creates a new user repository with the given database connection.
func NewRepository(db *sql.DB) Repository {
	return &dbRepository{db: db}
}

func (r *dbRepository) Create(u *User) error {
	query := `INSERT INTO users (id, username, password_hash, role, status, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, u.ID, u.Username, u.PasswordHash, u.Role, u.Status, u.CreatedAt)
	return err
}

func (r *dbRepository) FindByUsername(username string) (*User, error) {
	u := &User{}
	query := `SELECT id, username, password_hash, role, status, created_at FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *dbRepository) FindAllByRole(role string) ([]*User, error) {
	query := `SELECT id, username, password_hash, role, status, created_at FROM users WHERE role = $1 ORDER BY username ASC`
	rows, err := r.db.Query(query, role)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		u := &User{}
		err := rows.Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role, &u.Status, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *dbRepository) UpdateStatus(id uuid.UUID, status string) error {
	query := `UPDATE users SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}
