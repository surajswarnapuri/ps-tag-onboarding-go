// Package user contains the api layer for the user domain.
package user

import (
	"context"
	"encoding/json"
	"net/http"

	userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"

	"github.com/gorilla/mux"
	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/interface/shared"
)

const (
	findRoute = "/find/{id}"
	saveRoute = "/save"
)

type userApplicationService interface {
	Find(ctx context.Context, id string) (*userDomain.User, error)
	Save(ctx context.Context, user *userDomain.User) (*userDomain.User, error)
}

type Handler struct {
	userService userApplicationService
}

func (h Handler) Find() shared.Handler {
	return shared.Handler{
		Route: func(r *mux.Route) {
			r.Path(findRoute).Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			id := mux.Vars(r)["id"]
			user, err := h.userService.Find(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var userDTO UserDTO
			userDTO.FromDomain(user)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userDTO)
		},
	}
}

func (h Handler) Save() shared.Handler {
	return shared.Handler{
		Route: func(r *mux.Route) {
			r.Path(saveRoute).Methods("POST")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			var userRequest UserDTO
			json.NewDecoder(r.Body).Decode(&userRequest)
			user, err := h.userService.Save(r.Context(), userRequest.ToDomain())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			var userResponse UserDTO
			userResponse.FromDomain(user)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(userResponse)
		},
	}
}
