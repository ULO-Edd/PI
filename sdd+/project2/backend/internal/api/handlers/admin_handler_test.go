package handlers_test

import (
	"bytes"
	"encoding/json"
	"event-api/internal/api/handlers"
	"event-api/internal/domain/user"
	"event-api/internal/domain/user/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAdminHandler_ListOrganizers(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	userService := user.NewService(mockRepo)
	adminHandler := handlers.NewAdminHandler(nil, userService)

	t.Run("successfully list organizers", func(t *testing.T) {
		orgs := []*user.User{
			{
				ID:       uuid.New(),
				Username: "org1",
				Role:     "organiser",
				Status:   "Pending",
			},
		}

		mockRepo.On("FindAllByRole", "organiser").Return(orgs, nil).Once()

		req, err := http.NewRequest("GET", "/api/admin/organizers", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(adminHandler.ListOrganizers)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response []*user.User
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Len(t, response, 1)
		assert.Equal(t, "org1", response[0].Username)
	})
}

func TestAdminHandler_UpdateOrganizerStatus(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	userService := user.NewService(mockRepo)
	adminHandler := handlers.NewAdminHandler(nil, userService)

	t.Run("successfully approve organizer", func(t *testing.T) {
		id := uuid.New()
		mockRepo.On("UpdateStatus", id, "Approved").Return(nil).Once()

		body := map[string]string{"status": "Approved"}
		jsonBody, _ := json.Marshal(body)

		req, err := http.NewRequest("PUT", "/api/admin/organizers/"+id.String()+"/status", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)

		// Set path value Mock for Go 1.22 standard mux variables
		req.SetPathValue("id", id.String())

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(adminHandler.UpdateOrganizerStatus)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}
