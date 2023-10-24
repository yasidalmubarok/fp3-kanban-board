package dto

import "time"

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty, length(6|255)~Minimum password is 6 length"`
	Role     string `json:"role" valid:"required~role cannot be empty"`
}

type NewUserResponse struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type RoleUser struct {
	Role string `json:"role"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateResponse struct {
	Id       uint       `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type AdminResponse struct {
	Message string `json:"message"`
}

type GetUsers struct {
	Id       uint       `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
}
