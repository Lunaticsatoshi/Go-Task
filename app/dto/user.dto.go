package dto

import "github.com/Lunaticsatoshi/go-task/app/models"

type (
	UserRegisterRequest struct {
		Name     string `json:"name" form:"name" binding:"required"`
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserResponse struct {
		ID    string `json:"id"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	}

	UserListResponse struct {
		Users []models.User `json:"users"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserUpdateRequest struct {
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	}
)
