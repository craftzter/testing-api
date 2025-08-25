package handlers

import (
	"encoding/json"
	"monly-login-api/internal/dto"
	"monly-login-api/utils"
	"net/http"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// decode request
		var req dto.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		// panggil service
		err := h.UserService.CreateUser(ctx, req.Username, req.Email, req.Password)
		if err != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, "error creating user")
			return
		}

		// response sukses
		utils.ResponseWithSuccess(w, http.StatusCreated, "user created success", req.Username)
	}
}
