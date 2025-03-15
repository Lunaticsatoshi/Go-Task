package routes

import (
	"github.com/Lunaticsatoshi/go-task/app/api/v1/controllers"
	"github.com/Lunaticsatoshi/go-task/app/common/middlewares"
	"github.com/Lunaticsatoshi/go-task/app/services"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	UserController controllers.UserController
	JwtService     services.JWTService
}

func (urc *UserRouteController) UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/v1/users")
	router.POST("/register", urc.UserController.Register)
	router.POST("/login", urc.UserController.Login)

	authRouter := rg.Group("/v1/auth/users")
	authRouter.Use(middlewares.Authenticate(urc.JwtService, "user"))
	authRouter.GET("/", urc.UserController.GetAllUsers)
	authRouter.GET("/me", urc.UserController.Me)
	authRouter.GET("/:userId", urc.UserController.GetUserById)
	authRouter.PUT("/:userId", urc.UserController.UpdateUser)
}
