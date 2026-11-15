package api

import (
	"database/sql"
	"event-api/internal/api/handlers"
	"event-api/internal/api/middleware"
	"event-api/internal/domain/event"
	"event-api/internal/domain/registration"
	"event-api/internal/domain/user"
	"net/http"
	_ "event-api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(db *sql.DB) http.Handler {
	// User domain
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Event domain
	eventRepo := event.NewRepository(db)
	eventService := event.NewService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	// Registration domain
	regRepo := registration.NewRepository(db)
	regService := registration.NewService(regRepo, eventRepo)
	regHandler := handlers.NewRegistrationHandler(regService)

	// Admin domain
	adminHandler := handlers.NewAdminHandler(db, userService)

	// Attendance domain
	attendanceHandler := handlers.NewAttendanceHandler(db)

	// Csv domain
	csvHandler := handlers.NewCsvHandler(db)

	// Router
	mux := http.NewServeMux()

	// Swagger
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Public routes
	mux.HandleFunc("POST /register", userHandler.Register)
	mux.HandleFunc("POST /login", userHandler.Login)
	mux.HandleFunc("GET /events", eventHandler.GetEvents)
	mux.HandleFunc("GET /events/{eventId}", eventHandler.GetEventByID)
	mux.HandleFunc("GET /categories", adminHandler.ListCategories)

	// Protected routes (Admin-only)
	adminRoutes := http.NewServeMux()
	adminRoutes.HandleFunc("GET /admin/organizers", adminHandler.ListOrganizers)
	adminRoutes.HandleFunc("PUT /admin/organizers/{id}/status", adminHandler.UpdateOrganizerStatus)
	adminRoutes.HandleFunc("POST /categories", adminHandler.CreateCategory)
	mux.Handle("/admin/", middleware.JWTAuth(middleware.HasRole("administrator", adminRoutes)))
	mux.Handle("/categories", middleware.JWTAuth(middleware.HasRole("administrator", adminRoutes)))


	// Protected routes (Attendees)
	attendeeRoutes := http.NewServeMux()
	attendeeRoutes.HandleFunc("POST /events/{eventId}/register", regHandler.RegisterForEvent)
	attendeeRoutes.HandleFunc("DELETE /events/{eventId}/register", regHandler.UnregisterFromEvent)
	mux.Handle("/events/{eventId}/register", middleware.JWTAuth(middleware.HasRole("attendee", attendeeRoutes)))


	// Protected routes (Organisers)
	organiserRoutes := http.NewServeMux()
	organiserRoutes.HandleFunc("POST /events", eventHandler.CreateEvent)
	organiserRoutes.HandleFunc("PUT /events/{eventId}", eventHandler.UpdateEvent)
	organiserRoutes.HandleFunc("DELETE /events/{eventId}", eventHandler.DeleteEvent)
	organiserRoutes.HandleFunc("GET /events/{eventId}/attendees", attendanceHandler.GetAttendeeRoster)
	organiserRoutes.HandleFunc("PUT /events/{eventId}/attendees/{registrationId}/presence", attendanceHandler.UpdatePresence)
	organiserRoutes.HandleFunc("GET /events/{eventId}/attendees/export", csvHandler.ExportEventAttendees)
	mux.Handle("/events", middleware.JWTAuth(middleware.HasRole("organiser", organiserRoutes)))
	mux.Handle("/events/", middleware.JWTAuth(middleware.HasRole("organiser", organiserRoutes))) // For paths with eventId


	return middleware.CORS(mux)
}
