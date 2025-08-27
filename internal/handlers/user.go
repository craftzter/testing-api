package handlers

import (
	"encoding/json"
	"monly-login-api/internal/dto"
	"monly-login-api/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}
		user, err := h.UserService.CreateUser(ctx, req) // return user, bukan hanya error
		if err != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.ResponseWithSuccess(w, http.StatusCreated, "user created success", user)
	}
}

func (h *Handler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// decode request
		var req dto.LoginUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}

		// panggil service
		user, err := h.UserService.LoginUser(ctx, req)
		if err != nil {
			utils.ResponseWithError(w, http.StatusUnauthorized, "email or password salah")
			return
		}

		// generate JWT
		token, err := utils.GenerateJWT(int64(user.ID), user.Username, utils.SecretKey)
		if err != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, "failed to generate token")
			return
		}

		// response sukses, bisa kirim token dan data user seperlunya
		resp := map[string]interface{}{
			"token": token,
			"user":  user.Username, // bisa custom
		}
		utils.ResponseWithSuccess(w, http.StatusOK, "login success", resp)
	}
}

func (h *Handler) UpdateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// update user telah login dan it tersompan ambil id dari path url
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr) // conversi string
		if err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid user id")
			return
		}
		ctx := r.Context()
		// decode request
		var req dto.UpdateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			utils.ResponseWithError(w, http.StatusBadRequest, "invalid request payload")
			return
		}
		// panggil service
		updatedUser, err := h.UserService.UpdateUser(ctx, int32(id), req)
		if err != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, "error updating user")
			return
		}
		// response sukses
		utils.ResponseWithSuccess(w, http.StatusOK, "user updated success", updatedUser)
	}
}
