package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status" json:"status"`
}

func (Task) TableName() string {
	return "tasks"
}
