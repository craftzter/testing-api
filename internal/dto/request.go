package dto

// disini terdiam struct bisa utnuk req atau response berguna jika project growt besar

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
