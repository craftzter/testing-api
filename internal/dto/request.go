package dto

import db "monly-login-api/internal/generate"

// disini terdiam struct bisa utnuk req atau response berguna jika project growt besar

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

type UserResponse struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetUser(u db.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func GetUserList(users []db.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, GetUser(u))
	}
	return res
}
