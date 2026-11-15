package handlers

import (
	"database/sql"
	"encoding/json"
	"event-api/internal/domain/user"
	"net/http"

	"github.com/google/uuid"
)

type AdminHandler struct {
	db          *sql.DB
	userService *user.Service
}

func NewAdminHandler(db *sql.DB, userService *user.Service) *AdminHandler {
	return &AdminHandler{
		db:          db,
		userService: userService,
	}
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (h *AdminHandler) ListOrganizers(w http.ResponseWriter, r *http.Request) {
	organizers, err := h.userService.ListOrganizers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(organizers)
}

func (h *AdminHandler) UpdateOrganizerStatus(w http.ResponseWriter, r *http.Request) {
	orgID, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid organizer ID", http.StatusBadRequest)
		return
	}

	var req UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.UpdateStatus(orgID, req.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Status updated successfully"})
}

func (h *AdminHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("SELECT id, name, description FROM categories ORDER BY name ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		var desc sql.NullString
		err := rows.Scan(&c.ID, &c.Name, &desc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if desc.Valid {
			c.Description = desc.String
		}
		categories = append(categories, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *AdminHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var c Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if c.Name == "" {
		http.Error(w, "Category name cannot be empty", http.StatusBadRequest)
		return
	}

	c.ID = uuid.New()
	_, err := h.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", c.ID, c.Name, c.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}
