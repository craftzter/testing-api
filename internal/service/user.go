package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"monly-login-api/internal/dto"
	db "monly-login-api/internal/generate"
	"monly-login-api/utils"
	"time"
)

type UserService struct {
	queries *db.Queries
	logger  *slog.Logger
}

func NewUserService(q *db.Queries, logger *slog.Logger) *UserService {
	return &UserService{queries: q, logger: logger}
}

func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (db.User, error) {
	s.logger.Debug("starting function CreateUser", "email", req.Email)
	// validate the input using utils in utils folder
	if err := utils.ValidateRegisterInput(req); err != nil {
		s.logger.Warn("invalid register input", "err", err, "email", req.Email)
		return db.User{}, err
	}
	// check if the email is exists
	_, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err == nil {
		s.logger.Info("email already registered", "email", req.Email)
		return db.User{}, utils.ConflictError{"email has already used please use another email!"}
	}
	if !errors.Is(err, sql.ErrNoRows) {
		s.logger.Error("database error on email checking", "err", err, "email", req.Email)
		return db.User{}, err
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("failed to hash password", "err", err, "email", req.Email)
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
		s.logger.Error("failed to create user", "err", err, "email", req.Email)
	}
	s.logger.Info("user registered", "userID", user.ID, "email", req.Email)
	return user, nil
}

func (s *UserService) LoginUser(ctx context.Context, req dto.LoginUserRequest) (*db.User, error) {
	s.logger.Debug("starting functions login user", "email", req.Email)
	// collect user from db by email
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Warn("login failed: user not found", "email", req.Email)
		return nil, fmt.Errorf("invalid email or password: %w", err)
	}
	// compare password
	if !utils.ComparePassword(user.Password, req.Password) {
		s.logger.Warn("login failed: password mismatch", "email", req.Email)
		return nil, fmt.Errorf("invalid email or password: %w", errors.New("password does not match"))
	}
	s.logger.Info("user login success", "userID", user.ID, "email", user.Email)
	return &user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req dto.UpdateUserRequest) (db.User, error) {
	// Ambil user lama dari DB
	s.logger.Debug("starting function UpdateUser", "userID", id)
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		s.logger.Warn("update failed: user not found", "userID", id, "err", err)
		return db.User{}, err
	}

	// Update hanya field yang dikirim
   if req.Username != nil {
        if *req.Username == "" {
            return db.User{}, utils.ValidationError{"username cannot be empty"}
        }
        user.Username = *req.Username
    }
    if req.Email != nil {
        if *req.Email == "" {
            return db.User{}, utils.ValidationError{"email cannot be empty"}
        }
        user.Email = *req.Email
    }
    if req.Password != nil {
        if *req.Password == "" {
            return db.User{}, utils.ValidationError{"password cannot be empty"}
        }
        hashed, err := utils.HashPassword(*req.Password)
        if err != nil {
            s.logger.Error("failed to hashing password on update", "err", err, "userID", id)
            return db.User{}, err
        }
        user.Password = hashed
    }
    user.Updated = time.Now()

    updatedUser, err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
        Password: user.Password,
        Updated:  user.Updated,
    })
    if err != nil {
        s.logger.Error("failed to update user", "err", err, "userID", id)
        return db.User{}, err
    }
    s.logger.Info("user updated", "userID", updatedUser.ID, "email", updatedUser.Email)
    return updatedUser, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (dto.UserResponse, error) {
	s.logger.Debug("starting GetUserByID", "userID", id)
	user, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.UserResponse{}, utils.NotFoundError{"user not found"}
		}
		return dto.UserResponse{}, err
	}
	userDto := dto.GetUser(user)
	s.logger.Info("success get a user", "userID", userDto.ID)
	return userDto, nil
}

func (s *UserService) GetListUser(ctx context.Context, id int32) ([]dto.UserResponse, error) {
	s.logger.Debug("starting GetListUser")
	users, err := s.queries.ListUsers(ctx)
	if err != nil {
		return nil, utils.NotFoundError{"not any user are exists"}
	}

	s.logger.Info("succes get users", "userID")
	usersDto := dto.GetUserList(users)
	return usersDto, nil
}
