package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// --- Success & Error Response Structs ---

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// --- Error Types ---

type ValidationError struct {
	Msg string
}

func (e ValidationError) Error() string { return e.Msg }

type NotFoundError struct {
	Msg string
}

func (e NotFoundError) Error() string { return e.Msg }

type ConflictError struct {
	Msg string
}

func (e ConflictError) Error() string { return e.Msg }

type AuthError struct {
	Msg string
}

func (e AuthError) Error() string { return e.Msg }

// --- Success Response Helper ---

func ResponseWithSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(SuccessResponse{Message: message, Data: data})
}

// --- Error Response Helper ---

func ResponseWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// --- Centralized Error Handler ---

func ResponseWithAppropriateError(w http.ResponseWriter, err error) {
	log.Printf("error type: %T, value: %+v", err, err)
	switch {
	case errors.As(err, &ValidationError{}):
		var e ValidationError
		errors.As(err, &e)
		ResponseWithError(w, http.StatusBadRequest, e.Error())
	case errors.As(err, &NotFoundError{}):
		var e NotFoundError
		errors.As(err, &e)
		ResponseWithError(w, http.StatusNotFound, e.Error())
	case errors.As(err, &ConflictError{}):
		var e ConflictError
		errors.As(err, &e)
		ResponseWithError(w, http.StatusConflict, e.Error())
	case errors.As(err, &AuthError{}):
		var e AuthError
		errors.As(err, &e)
		ResponseWithError(w, http.StatusUnauthorized, e.Error())
	default:
		ResponseWithError(w, http.StatusInternalServerError, "internal server error")
	}
}
