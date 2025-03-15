package dto

import (
	"github.com/Lunaticsatoshi/go-task/app/common/interfaces"
	"github.com/Lunaticsatoshi/go-task/app/models"
)

type (
	TaskCreateRequest struct {
		Title       string `json:"title" form:"title" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
		Status      string `json:"status" form:"status" binding:"required"`
		UserId      uint   `json:"user_id" form:"user_id" binding:"required"`
	}

	TaskResponse struct {
		ID          string `json:"id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Status      string `json:"status,omitempty"`
		UserId      uint   `json:"user_id,omitempty"`
	}

	TaskListResponse struct {
		Meta  interfaces.PaginationMeta `json:"meta"`
		Tasks []models.Task             `json:"tasks"`
	}

	TaskUpdateRequest struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Status      string `json:"status,omitempty"`
	}
)
