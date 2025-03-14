package models

import (
	"time"

	"github.com/Lunaticsatoshi/go-task/app/common/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"column:id;primary_key" json:"id"`
	Name      string    `gorm:"column:name" json:"name,omitempty"`
	Email     string    `gorm:"column:email" json:"email,omitempty"`
	Phone     string    `gorm:"column:phone" json:"phone,omitempty"`
	Password  string    `gorm:"column:password" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = utils.PasswordHash(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" {
		var err error
		u.Password, err = utils.PasswordHash(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}

func (User) TableName() string {
	return "users"
}
