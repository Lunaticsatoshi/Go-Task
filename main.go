package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Lunaticsatoshi/go-task/app/api/v1/controllers"
	"github.com/Lunaticsatoshi/go-task/app/api/v1/routes"
	"github.com/Lunaticsatoshi/go-task/app/common/middlewares"
	"github.com/Lunaticsatoshi/go-task/app/services"
	"github.com/Lunaticsatoshi/go-task/docs"
	"github.com/Lunaticsatoshi/go-task/initializers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	server              *gin.Engine
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	// Initialize the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initializers.MigrateDB()
	initializers.ConnectDB()

	if os.Getenv("ENV") != "LOCAL" {
		gin.SetMode(gin.ReleaseMode)
	}

	userService := services.NewUserService(initializers.DB)
	jwtService := services.NewJWTService()
	UserController = controllers.UserController{
		DB:          initializers.DB,
		UserService: userService,
		JwtService:  jwtService,
	}
	UserRouteController = routes.UserRouteController{
		UserController: UserController,
		JwtService:     jwtService,
	}

	server = gin.Default()
	server.Use(
		middlewares.CORSMiddleware(),
	)

}

func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "GO Task"
	docs.SwaggerInfo.Description = "An Go based Api for Task Management"
	docs.SwaggerInfo.Version = os.Getenv("SWAGGER_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_URL")
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := server.Group("/api")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	UserRouteController.UserRoutes(router)

	server.Run(":8080")
}
