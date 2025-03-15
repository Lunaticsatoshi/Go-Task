package seeders

import (
	"errors"

	"github.com/Lunaticsatoshi/go-task/app/models"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	var dummyUsers = []models.User{
		{
			Name:     "Satoshi",
			Email:    "satoshi@gmail.com",
			Password: "satoshi@91",
		},
		{
			Name:     "User",
			Email:    "user@gmail.com",
			Password: "user1",
		},
		{
			Name:     "User2",
			Email:    "user2@gmail.com",
			Password: "user2",
		},
		{
			Name:     "User3",
			Email:    "user3@gmail.com",
			Password: "user3",
		},
		{
			Name:     "User4",
			Email:    "user4@gmail.com",
			Password: "user4",
		},
	}

	hasTable := db.Migrator().HasTable(&models.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&models.User{}); err != nil {
			return err
		}
	}

	for _, data := range dummyUsers {
		var user models.User
		err := db.Where(&models.User{Email: data.Email}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "email = ?", data.Email).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
