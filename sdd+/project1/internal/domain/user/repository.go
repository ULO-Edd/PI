package user

import (
	"database/sql"
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
	query := `INSERT INTO users (id, username, password_hash, role, created_at)
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, u.ID, u.Username, u.PasswordHash, u.Role, u.CreatedAt)
	return err
}

func (r *dbRepository) FindByUsername(username string) (*User, error) {
	u := &User{}
	query := `SELECT id, username, password_hash, role, created_at FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
