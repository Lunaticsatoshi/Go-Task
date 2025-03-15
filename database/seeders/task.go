package seeders

import (
	"errors"

	"github.com/Lunaticsatoshi/go-task/app/common/constants"
	"github.com/Lunaticsatoshi/go-task/app/models"
	"gorm.io/gorm"
)

func TaskSeeder(db *gorm.DB) error {
	var dummyTasks = []models.Task{
		{
			Title:       "Test Task 1",
			Description: "This is a test task 1",
			Status:      constants.EnumStatusStarted,
			UserId:      1,
		},
		{
			Title:       "Test Task 2",
			Description: "This is a test task 2",
			Status:      constants.EnumStatusPending,
			UserId:      1,
		},
		{
			Title:       "Test Task 3",
			Description: "This is a test task 3",
			Status:      constants.EnumStatusInReview,
			UserId:      1,
		},
		{
			Title:       "Test Task 4",
			Description: "This is a test task 4",
			Status:      constants.EnumStatusInReview,
			UserId:      1,
		},
		{
			Title:       "Test Task 5",
			Description: "This is a test task 5",
			Status:      constants.EnumStatusCompleted,
			UserId:      1,
		},
		{
			Title:       "Test Task 6",
			Description: "This is a test task 6",
			Status:      constants.EnumStatusPending,
			UserId:      1,
		},
		{
			Title:       "Test Task 7",
			Description: "This is a test task 7",
			Status:      constants.EnumStatusCompleted,
			UserId:      1,
		},
		{
			Title:       "Test Task 8",
			Description: "This is a test task 8",
			Status:      constants.EnumStatusCompleted,
			UserId:      1,
		},
		{
			Title:       "Test Task 1",
			Description: "This is a test task 1",
			Status:      constants.EnumStatusStarted,
			UserId:      3,
		},
		{
			Title:       "Test Task 3",
			Description: "This is a test task 1",
			Status:      constants.EnumStatusStarted,
			UserId:      3,
		},
		{
			Title:       "Test Task 3",
			Description: "This is a test task 1",
			Status:      constants.EnumStatusInReview,
			UserId:      3,
		},
		{
			Title:       "Test Task 1",
			Description: "This is a test task 1",
			Status:      constants.EnumStatusStarted,
			UserId:      2,
		},
	}

	hasTable := db.Migrator().HasTable(&models.Task{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&models.Task{}); err != nil {
			return err
		}
	}

	for _, data := range dummyTasks {
		var task models.Task
		err := db.Where(&models.Task{Title: data.Title}).First(&task).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&task, "title = ?", data.Title).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
