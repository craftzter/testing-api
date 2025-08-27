package service

import (
	"context"
	"errors"
	"fmt"
	"monly-login-api/internal/dto"
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

func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (db.User, error) {
	// validate the input using utils in utils folder
	if err := utils.ValidateRegisterInput(req); err != nil {
		return db.User{}, err
	}
	// hashing password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return db.User{}, err
	}

	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Created:  time.Now(),
		Updated:  time.Now(),
	})
	if err != nil {
		return db.User{}, err
	}
	return user, nil
}

func (s *UserService) LoginUser(ctx context.Context, req dto.LoginUserRequest) (*db.User, error) {
	// collect user from db by email
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password: %w", err)
	}
	// compare password
	if !utils.ComparePassword(user.Password, req.Password) {
		return nil, fmt.Errorf("invalid email or password: %w", errors.New("password does not match"))
	}
	return &user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req dto.UpdateUserRequest) (db.User, error) {
	// Ambil user lama dari DB
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, err
	}

	// Update hanya field yang dikirim
	if req.Username != nil {
		user.Username = *req.Username
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		hashed, err := utils.HashPassword(*req.Password)
		if err != nil {
			return db.User{}, err
		}
		user.Password = hashed
	}
	// Update updated_at
	user.Updated = time.Now()

	// Jalankan update ke DB (bisa pakai sqlc custom query, atau update semua field)
	updatedUser, err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Updated:  user.Updated,
	})
	return updatedUser, err
}
