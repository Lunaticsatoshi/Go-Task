package models

type Task struct {
	ID          uint   `gorm:"column:id;primary_key" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Status      string `gorm:"column:status" json:"status"`
	UserId      uint   `gorm:"column:user_id" json:"user_id"`
}

func (Task) TableName() string {
	return "tasks"
}
