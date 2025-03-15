package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Lunaticsatoshi/go-task/app/common/constants"
	"github.com/Lunaticsatoshi/go-task/app/common/interfaces"
	"github.com/Lunaticsatoshi/go-task/app/common/utils"
	"github.com/Lunaticsatoshi/go-task/app/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController struct {
	DB          *gorm.DB
	TaskService services.TaskService
}

func (tc *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, serviceErr := tc.TaskService.GetAllUserTasks(ctx)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.TaskFetchedSuccessfully, http.StatusCreated, tasks))
}

func (tc *TaskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("taskId")
	task, serviceErr := tc.TaskService.GetTaskById(ctx, utils.ConvertStringToUInt(id))
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.TaskFetchedSuccessfully, http.StatusCreated, task))
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(constants.ErrorReadingRequestBody, err.Error())
		ctx.JSON(http.StatusBadRequest, interfaces.CreateFailResponse(constants.ErrorReadingRequestBody, err.Error(), http.StatusBadRequest))
		return
	}

	var rawPayload json.RawMessage = requestBody

	task, serviceErr := tc.TaskService.CreateNewTask(ctx, rawPayload)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.TaskCreatedSuccessfully, http.StatusCreated, task))
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	id := ctx.Param("taskId")
	if err != nil {
		log.Println(constants.ErrorReadingRequestBody, err.Error())
		ctx.JSON(http.StatusBadRequest, interfaces.CreateFailResponse(constants.ErrorReadingRequestBody, err.Error(), http.StatusBadRequest))
		return
	}

	var rawPayload json.RawMessage = requestBody

	task, serviceErr := tc.TaskService.UpdateTask(ctx, utils.ConvertStringToUInt(id), rawPayload)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.TaskUpdatedSuccessfully, http.StatusCreated, task))
}
