package models

import "time"

type Task struct {
	ID          uint      `gorm:"column:id;primary_key" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status" json:"status"`
	UserId      uint      `gorm:"column:user_id" json:"user_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Task) TableName() string {
	return "tasks"
}
