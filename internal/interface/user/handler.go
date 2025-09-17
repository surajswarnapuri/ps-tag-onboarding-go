package user

import (
	"context"
	"encoding/json"
	"net/http"

	userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type userService interface {
	Find(ctx context.Context, id string) (*userDomain.User, error)
	Save(ctx context.Context, user *userDomain.User) (*userDomain.User, error)
}

type Handler struct {
	userService userService
}

func NewHandler(userService userService) *Handler {
	return &Handler{userService: userService}
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	user, err := h.userService.Find(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var userDTO UserDTO
	userDTO.FromDomain(user)
	json.NewEncoder(w).Encode(userDTO)
}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request) {
	var userDTORequest UserDTO
	json.NewDecoder(r.Body).Decode(&userDTORequest)
	user, err := h.userService.Save(r.Context(), userDTORequest.ToDomain())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var userDTOResponse UserDTO
	userDTOResponse.FromDomain(user)
	json.NewEncoder(w).Encode(userDTOResponse)
}
