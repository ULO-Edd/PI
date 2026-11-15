package handlers

import (
	"encoding/json"
	"event-api/internal/domain/user"
	"event-api/internal/platform/auth"
	"net/http"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{service: service}
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with a username, password, and role.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body      RegisterRequest  true  "User Registration"
// @Success 201   {object}  user.User
// @Router /register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.service.Register(req.Username, req.Password, req.Role)
	if err != nil {
		switch err {
		case user.ErrInvalidUsername, user.ErrInvalidPassword, user.ErrInvalidRole:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case user.ErrUsernameTaken:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user with username and password to receive a JWT.
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body      LoginRequest     true  "User Login"
// @Success 200   {object}  LoginResponse
// @Router /login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		if err == user.ErrUserNotFound || err == user.ErrInvalidCredentials {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := auth.GenerateJWT(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
