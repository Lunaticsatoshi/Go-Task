package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lunaticsatoshi/go-task/app/common/constants"
	"github.com/Lunaticsatoshi/go-task/app/common/utils"
	"github.com/Lunaticsatoshi/go-task/app/dto"
	"github.com/Lunaticsatoshi/go-task/app/models"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

type UserService interface {
	GetAllUsers(ctx context.Context) (dto.UserListResponse, *utils.ServiceError)
	GetUserById(ctx context.Context, id uint) (dto.UserResponse, *utils.ServiceError)
	VerifyLogin(ctx context.Context, rawPayload json.RawMessage) (*models.User, *utils.ServiceError)
	CreateNewUser(ctx context.Context, rawPayload json.RawMessage) (dto.UserResponse, *utils.ServiceError)
	UpdateUser(ctx context.Context, id uint, rawPayload json.RawMessage) (dto.UserResponse, *utils.ServiceError)
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (us *userService) GetAllUsers(ctx context.Context) (dto.UserListResponse, *utils.ServiceError) {
	var users []models.User

	// Finds all users
	findError := us.db.Find(&users)
	if findError.Error != nil {
		return dto.UserListResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindUserError, InternalErrorMessage: findError.Error.Error()}
	}

	return dto.UserListResponse{
		Users: users,
	}, nil
}

func (us *userService) GetUserById(ctx context.Context, id uint) (dto.UserResponse, *utils.ServiceError) {
	var user *models.User

	// Update a user
	findError := us.db.Where(models.User{ID: id}).First(&user)
	if findError.Error != nil {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindUserError, InternalErrorMessage: findError.Error.Error(), Payload: utils.ConvertIntToString(int(id))}
	}

	return dto.UserResponse{
		ID:    utils.ConvertIntToString(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (us *userService) VerifyLogin(ctx context.Context, rawPayload json.RawMessage) (*models.User, *utils.ServiceError) {
	var ld dto.UserLoginRequest
	if err := json.Unmarshal(rawPayload, &ld); err != nil {
		return nil, &utils.ServiceError{Code: http.StatusBadRequest, Message: constants.InputError, InternalErrorMessage: err.Error(), Payload: string(rawPayload)}
	}
	var user models.User
	userCheck := us.db.Where(models.User{Email: ld.Email}).First(&user)
	if userCheck.Error != nil {
		return nil, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.FindUserError, InternalErrorMessage: userCheck.Error.Error()}
	}
	passwordCheck, err := utils.PasswordCompare(user.Password, []byte(ld.Password))
	if err != nil {
		return nil, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.InvalidPassword, InternalErrorMessage: err.Error()}
	}

	if user.Email != ld.Email || !passwordCheck {
		return nil, &utils.ServiceError{Code: http.StatusUnauthorized, Message: constants.InvalidLoginDetails, InternalErrorMessage: userCheck.Error.Error()}
	}
	return &user, nil
}

func (us *userService) CreateNewUser(ctx context.Context, rawPayload json.RawMessage) (dto.UserResponse, *utils.ServiceError) {
	var ud dto.UserRegisterRequest
	if err := json.Unmarshal(rawPayload, &ud); err != nil {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusBadRequest, Message: constants.InputError, InternalErrorMessage: err.Error(), Payload: string(rawPayload)}
	}
	user := models.User{}

	userCheck := us.db.Where(models.User{Email: ud.Email}).First(&user)
	if userCheck.Error != nil && !(errors.Is(userCheck.Error, gorm.ErrRecordNotFound)) {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.CreateUserError, InternalErrorMessage: userCheck.Error.Error(), Payload: string(ud.Email)}
	}

	if user.Email == ud.Email {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.EmailExistsError, Payload: string(ud.Email)}
	}

	createdUser := models.User{
		Name:     ud.Name,
		Email:    ud.Email,
		Password: ud.Password,
	}
	// create new user
	newUser := us.db.Where(models.User{Email: user.Email}).Assign(user).Create(&createdUser)
	if newUser.Error != nil {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.CreateUserError, InternalErrorMessage: newUser.Error.Error(), Payload: string(ud.Email)}
	}

	return dto.UserResponse{
		ID:    utils.ConvertIntToString(int(createdUser.ID)),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}, nil
}

func (us *userService) UpdateUser(ctx context.Context, id uint, rawPayload json.RawMessage) (dto.UserResponse, *utils.ServiceError) {
	var ud dto.UserUpdateRequest
	if err := json.Unmarshal(rawPayload, &ud); err != nil {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusBadRequest, Message: constants.InputError, InternalErrorMessage: err.Error(), Payload: string(rawPayload)}
	}

	user := models.User{
		Name:  ud.Name,
		Email: ud.Email,
		Phone: ud.Phone,
	}

	// Update a user
	updatedUser := us.db.Where(models.User{ID: id}).Assign(user).Updates(&user)
	if updatedUser.Error != nil {
		return dto.UserResponse{}, &utils.ServiceError{Code: http.StatusInternalServerError, Message: constants.UpdateUserError, InternalErrorMessage: updatedUser.Error.Error(), Payload: string(ud.Email)}
	}

	return dto.UserResponse{
		ID:    utils.ConvertIntToString(int(user.ID)),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
