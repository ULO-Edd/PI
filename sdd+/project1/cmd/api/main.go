// @title Event Management API
// @version 1.0
// @description This is the API for managing events at ESTG/IPLeiria.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @title Event Management API
// @version 1.0
// @description This is a sample server for an event management system.
// @host localhost:8080
// @BasePath /

package main

import (
	"event-api/internal/api"
	"event-api/internal/config"
	"event-api/internal/platform/database"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "event-api/docs"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	m, err := migrate.New(
		"file://db/migrations",
		cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("Migrations ran successfully")

	router := api.NewRouter(db)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
