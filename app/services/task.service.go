package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lunaticsatoshi/go-task/app/common/constants"
	"github.com/Lunaticsatoshi/go-task/app/common/utils"
	"github.com/Lunaticsatoshi/go-task/app/dto"
	"github.com/Lunaticsatoshi/go-task/app/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type taskService struct {
	db *gorm.DB
}

type TaskService interface {
	GetAllTasks(ctx *gin.Context) (dto.TaskListResponse, *utils.ServiceError)
	GetAllUserTasks(ctx *gin.Context) (dto.TaskListResponse, *utils.ServiceError)
	GetTaskById(ctx *gin.Context, id uint) (dto.TaskResponse, *utils.ServiceError)
	CreateNewTask(ctx *gin.Context, rawPayload json.RawMessage) (dto.TaskResponse, *utils.ServiceError)
	UpdateTask(ctx *gin.Context, id uint, rawPayload json.RawMessage) (dto.TaskResponse, *utils.ServiceError)
	DeleteTask(ctx *gin.Context, id uint) *utils.ServiceError
}

func NewTaskService(db *gorm.DB) TaskService {
	return &taskService{db: db}
}

func (ts *taskService) GetAllTasks(ctx *gin.Context) (dto.TaskListResponse, *utils.ServiceError) {
	var tasks []models.Task
	intPage, intLimit, offset, sortKey, sortOrder := utils.GetRequestPaginationData(ctx)
	query, args := utils.DynamicFilterTasks(ctx)

	if len(query) == 0 {
		query = "1 = 1"
		args = []interface{}{}
	}

	// Finds all users
	findError := ts.db.Where(query, args...).Limit(intLimit).Offset(offset).Order(fmt.Sprintf("%s %s", sortKey, sortOrder)).Find(&tasks)
	if findError.Error != nil {
		return dto.TaskListResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindTaskError, InternalErrorMessage: findError.Error.Error()}
	}

	meta := utils.GeneratePaginationMeta(ts.db, intPage, intLimit, &models.Task{})
	return dto.TaskListResponse{
		Meta:  meta,
		Tasks: tasks,
	}, nil
}

func (ts *taskService) GetAllUserTasks(ctx *gin.Context) (dto.TaskListResponse, *utils.ServiceError) {
	var tasks []models.Task
	userId := ctx.MustGet("UserID").(string)
	intPage, intLimit, offset, sortKey, sortOrder := utils.GetRequestPaginationData(ctx)
	query, args := utils.DynamicFilterTasks(ctx)

	if len(query) == 0 {
		query = "1 = 1"
		args = []interface{}{}
	}

	// Finds all users
	findError := ts.db.Where(fmt.Sprintf("%s AND user_id = ?", query), append(args, userId)...).Limit(intLimit).Offset(offset).Order(fmt.Sprintf("%s %s", sortKey, sortOrder)).Find(&tasks)
	if findError.Error != nil {
		return dto.TaskListResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindTaskError, InternalErrorMessage: findError.Error.Error()}
	}

	meta := utils.GeneratePaginationMeta(ts.db, intPage, intLimit, &models.Task{})
	return dto.TaskListResponse{
		Meta:  meta,
		Tasks: tasks,
	}, nil
}

func (ts *taskService) GetTaskById(ctx *gin.Context, id uint) (dto.TaskResponse, *utils.ServiceError) {
	var task *models.Task
	userId := ctx.MustGet("UserID").(string)
	// Get a Task by Id
	findError := ts.db.Where(models.Task{ID: id, UserId: utils.ConvertStringToUInt(userId)}).First(&task)
	if findError.Error != nil {
		return dto.TaskResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindTaskError, InternalErrorMessage: findError.Error.Error(), Payload: utils.ConvertIntToString(int(id))}
	}

	return dto.TaskResponse{
		ID:          utils.ConvertIntToString(int(task.ID)),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}

func (ts *taskService) CreateNewTask(ctx *gin.Context, rawPayload json.RawMessage) (dto.TaskResponse, *utils.ServiceError) {
	var td dto.TaskCreateRequest
	userId := ctx.MustGet("UserID").(string)
	if err := json.Unmarshal(rawPayload, &td); err != nil {
		return dto.TaskResponse{}, &utils.ServiceError{Code: http.StatusBadRequest, Message: constants.InputError, InternalErrorMessage: err.Error(), Payload: string(rawPayload)}
	}

	createdTask := models.Task{
		Title:       td.Title,
		Description: td.Description,
		Status:      td.Status,
		UserId:      utils.ConvertStringToUInt(userId),
	}
	// create new task
	newUser := ts.db.Where(models.Task{Title: createdTask.Title}).Assign(createdTask).FirstOrCreate(&createdTask)
	if newUser.Error != nil {
		return dto.TaskResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.CreateTaskError, InternalErrorMessage: newUser.Error.Error(), Payload: string(td.Title)}
	}

	return dto.TaskResponse{
		ID:          utils.ConvertIntToString(int(createdTask.ID)),
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		UserId:      createdTask.UserId,
	}, nil
}

func (ts *taskService) UpdateTask(ctx *gin.Context, id uint, rawPayload json.RawMessage) (dto.TaskResponse, *utils.ServiceError) {
	var td dto.TaskUpdateRequest
	userId := ctx.MustGet("UserID").(string)
	if err := json.Unmarshal(rawPayload, &td); err != nil {
		return dto.TaskResponse{}, &utils.ServiceError{Code: http.StatusBadRequest, Message: constants.InputError, InternalErrorMessage: err.Error(), Payload: string(rawPayload)}
	}

	task := models.Task{
		Title:       td.Title,
		Description: td.Description,
		Status:      td.Status,
	}

	// Update a task
	updatedUser := ts.db.Where(models.Task{ID: id, UserId: utils.ConvertStringToUInt(userId)}).Assign(task).Updates(&task)
	if updatedUser.Error != nil {
		return dto.TaskResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.UpdateTaskError, InternalErrorMessage: updatedUser.Error.Error(), Payload: string(td.Title)}
	}

	return dto.TaskResponse{
		ID:          utils.ConvertIntToString(int(task.ID)),
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}, nil
}

func (ts *taskService) DeleteTask(ctx *gin.Context, id uint) *utils.ServiceError {
	userId := ctx.MustGet("UserID").(string)
	// Delete a task
	deleteError := ts.db.Where(models.Task{ID: id, UserId: utils.ConvertStringToUInt(userId)}).Delete(&models.Task{})
	if deleteError.Error != nil {
		return &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.DeleteTaskError, InternalErrorMessage: deleteError.Error.Error(), Payload: utils.ConvertIntToString(int(id))}
	}
	return nil
}
