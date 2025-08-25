package service

import (
	"context"
	db "monly-login-api/internal/generate"
	"monly-login-api/utils"
	"time"
)

type UserService struct {
	queries *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{queries: q}
}

func (s *UserService) CreateUser(ctx context.Context, username, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = s.queries.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Created:  time.Now(),
		Updated:  time.Now(),
	})
	return err
}
