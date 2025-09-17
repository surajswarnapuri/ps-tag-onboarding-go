package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type mockUserApplicationService struct {
	FindFunc func(ctx context.Context, id string) (*userDomain.User, error)
}

func (m *mockUserApplicationService) Find(ctx context.Context, id string) (*userDomain.User, error) {
	return m.FindFunc(ctx, id)
}

func TestFind_HappyPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	serviceUser := &userDomain.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25}
	handler := Handler{
		userService: &mockUserApplicationService{
			FindFunc: func(ctx context.Context, id string) (*userDomain.User, error) {
				return serviceUser, nil
			},
		},
	}
	handler.Find().AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/find/1", nil))
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
	var userDTO UserDTO
	json.Unmarshal(w.Body.Bytes(), &userDTO)
	if userDTO.ID != serviceUser.ID {
		t.Errorf("expected user ID %s, got %s", serviceUser.ID, userDTO.ID)
	}
	if userDTO.FirstName != serviceUser.FirstName {
		t.Errorf("expected user FirstName %s, got %s", serviceUser.FirstName, userDTO.FirstName)
	}
	if userDTO.LastName != serviceUser.LastName {
		t.Errorf("expected user LastName %s, got %s", serviceUser.LastName, userDTO.LastName)
	}
	if userDTO.Email != serviceUser.Email {
		t.Errorf("expected user Email %s, got %s", serviceUser.Email, userDTO.Email)
	}
	if userDTO.Age != serviceUser.Age {
		t.Errorf("expected user Age %d, got %d", serviceUser.Age, userDTO.Age)
	}
}

func TestFind_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	handler := Handler{
		userService: &mockUserApplicationService{
			FindFunc: func(ctx context.Context, id string) (*userDomain.User, error) {
				return nil, fmt.Errorf("user not found")
			},
		},
	}
	handler.Find().AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/find/1", nil))
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}
