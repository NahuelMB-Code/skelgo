package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/NahuelMB-Code/skelgo/internal/application/user"
	"github.com/NahuelMB-Code/skelgo/internal/interfaces/http/dto"
)

type UserHandler struct {
	userService *user.Service
}

func NewUserHandler(us *user.Service) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

// CreateUser maneja POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	// Parsear JSON entrante al DTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Crear entidad User
	newUser := &user.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Llamar al caso de uso
	createdUser, err := h.userService.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el usuario creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
