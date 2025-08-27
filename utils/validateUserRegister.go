package utils

import (
	"errors"
	"monly-login-api/internal/dto"
	"net/mail"
	"strings"
	"unicode"
)

func ValidateRegisterInput(req dto.CreateUserRequest) error {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return errors.New("all field are required")
	}
	// validate email format
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return errors.New("invalid email format")
	}
	// validate email should contain email prefix
	if !strings.Contains(req.Email, "@gmail") {
		return errors.New("email must contain '@gmail'")
	}
	// password len 8 to 14
	if len(req.Password) < 8 || len(req.Password) >= 14 {
		return errors.New("password must be at least 8 to 14 characters")
	}
	// validate passowrd input should have uppercase lowercase and number also spaces not allowed
	var hasUpper, hasLower, hasNumber bool
	for _, c := range req.Password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasNumber = true
		}
	}

	if !hasUpper {
		return errors.New("password must have at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must have at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must have at least one number")
	}

	if strings.Contains(req.Password, " ") {
		return errors.New("password must not contain spaces")
	}

	return nil
}
