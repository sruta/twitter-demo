package dto

import (
	"time"
	"twitter-uala/internal/domain"
)

type UserCreate struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Username string `json:"username" validate:"required,min=3,max=32"`
}

type UserUpdate struct {
	ID       int64  `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,min=3,max=32"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUserCreateToUser(dto UserCreate) domain.User {
	return domain.User{
		Email:    dto.Email,
		Password: dto.Password,
		Username: dto.Username,
	}
}

func FromUserUpdateToUser(dto UserUpdate) domain.User {
	return domain.User{
		ID:       dto.ID,
		Username: dto.Username,
	}
}

func FromUserToUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
