package routes

import (
	"github.com/Lunaticsatoshi/go-task/app/api/v1/controllers"
	"github.com/Lunaticsatoshi/go-task/app/common/middlewares"
	"github.com/Lunaticsatoshi/go-task/app/services"
	"github.com/gin-gonic/gin"
)

type TaskRouteController struct {
	TaskController controllers.TaskController
	JwtService     services.JWTService
}

func (trc *TaskRouteController) TaskRoutes(rg *gin.RouterGroup) {
	authRouter := rg.Group("/v1/auth/tasks")
	authRouter.Use(middlewares.Authenticate(trc.JwtService, "user"))
	authRouter.GET("", trc.TaskController.GetAllTasks)
	authRouter.GET("/:taskId", trc.TaskController.GetTaskById)
	authRouter.POST("/", trc.TaskController.CreateTask)
	authRouter.PUT("/:taskId", trc.TaskController.UpdateTask)
	authRouter.DELETE("/:taskId", trc.TaskController.DeleteTask)
}
