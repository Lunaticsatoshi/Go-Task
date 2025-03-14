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

type UserController struct {
	DB          *gorm.DB
	UserService services.UserService
	JwtService  services.JWTService
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, serviceErr := uc.UserService.GetAllUsers(ctx)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.UserFetchedSuccessfully, http.StatusCreated, users))
}

func (uc *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("userId")
	user, serviceErr := uc.UserService.GetUserById(ctx, utils.ConvertStringToUInt(id))
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.UserFetchedSuccessfully, http.StatusCreated, user))
}

func (uc *UserController) Me(ctx *gin.Context) {
	id := ctx.MustGet("UserID").(string)
	user, serviceErr := uc.UserService.GetUserById(ctx, utils.ConvertStringToUInt(id))
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.UserFetchedSuccessfully, http.StatusCreated, user))
}

func (uc *UserController) Register(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(constants.ErrorReadingRequestBody, err.Error())
		ctx.JSON(http.StatusBadRequest, interfaces.CreateFailResponse(constants.ErrorReadingRequestBody, err.Error(), http.StatusBadRequest))
		return
	}

	var rawPayload json.RawMessage = requestBody

	user, serviceErr := uc.UserService.CreateNewUser(ctx, rawPayload)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	uc.UserService.CreateNewUser(ctx, rawPayload)
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.UserCreatedSuccessfully, http.StatusCreated, user))
}

func (uc *UserController) Login(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(constants.ErrorReadingRequestBody, err.Error())
		ctx.JSON(http.StatusBadRequest, interfaces.CreateFailResponse(constants.ErrorReadingRequestBody, err.Error(), http.StatusBadRequest))
		return
	}

	var rawPayload json.RawMessage = requestBody

	user, serviceErr := uc.UserService.VerifyLogin(ctx, rawPayload)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	token := uc.JwtService.GenerateToken(utils.ConvertIntToString(int(user.ID)), "user")
	authResp := interfaces.CreateAuthResponse(token, "user")
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse("User logged in successfully", http.StatusCreated, authResp))
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	requestBody, err := io.ReadAll(ctx.Request.Body)
	id := ctx.MustGet("UserID").(string)
	if err != nil {
		log.Println(constants.ErrorReadingRequestBody, err.Error())
		ctx.JSON(http.StatusBadRequest, interfaces.CreateFailResponse(constants.ErrorReadingRequestBody, err.Error(), http.StatusBadRequest))
		return
	}

	var rawPayload json.RawMessage = requestBody

	user, serviceErr := uc.UserService.UpdateUser(ctx, utils.ConvertStringToUInt(id), rawPayload)
	if serviceErr != nil {
		log.Println(serviceErr.LogMessage())
		ctx.JSON(serviceErr.Code, interfaces.CreateFailResponse(serviceErr.Message, serviceErr.InternalErrorMessage, uint(serviceErr.Code)))
		return
	}
	ctx.JSON(http.StatusCreated, interfaces.CreateSuccessResponse(constants.UserUpdatedSuccessfully, http.StatusCreated, user))
}
