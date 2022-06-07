package services

import (
	"finalproject/helpers"
	"finalproject/models"
	"finalproject/params"
	"finalproject/repositories"
	"net/http"
	"time"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

// Register New User
func (us *UserService) RegisterUser(request params.RegisterUserRq) (*params.RegisterUserRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.RegisterUserRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	model := models.User{
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
		},
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
	}

	// call repository register user
	modelUserResponse, err := us.userRepo.RegisterUser(&model)
	if err != nil {
		return &params.RegisterUserRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	return &params.RegisterUserRs{
		Age:      modelUserResponse.Age,
		Email:    modelUserResponse.Email,
		ID:       modelUserResponse.ID,
		Username: modelUserResponse.Username,
	}, err
}

// Register New User
func (us *UserService) UserLogin(request params.UserLoginRq) (*params.LoginUserRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.LoginUserRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	// call repository register user
	modelUserResponse, err := us.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return &params.LoginUserRs{
			Response: params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Invalid Email / Password",
			},
		}, err
	}

	// validate password
	isPasswordOk := helpers.ComparePass(modelUserResponse.Password, request.Password)
	if !isPasswordOk {
		return &params.LoginUserRs{
			Response: params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Invalid Email / Password",
			},
		}, err
	}

	// generate token
	token := helpers.GenerateToken(modelUserResponse.ID, modelUserResponse.Email)

	return &params.LoginUserRs{
		Token: token,
	}, err
}

// Update User
func (us *UserService) UpdateUser(request params.UpdateUserRq, userId, idToken uint, emailToken string) (*params.UpdateUserRs, *params.Response, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		}, err
	}

	// validate token userid and userid param
	if userId != idToken {
		return nil, &params.Response{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Error:   "your token / user id invalid",
		}, err
	}

	userModelRq := models.User{
		BaseModel: models.BaseModel{ID: userId},
		Username:  request.Username,
		Email:     request.Email,
	}

	// call repository update user by id
	err = us.userRepo.UpdateUserById(&userModelRq)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Error:   err.Error(),
		}, err
	}

	// call to get user by id
	userUpdated, err := us.userRepo.GetUserById(userId)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Error:   "Not Authorized Update",
		}, err
	}

	return &params.UpdateUserRs{
		Age:       userUpdated.Age,
		Email:     userUpdated.Email,
		ID:        userUpdated.ID,
		Username:  userUpdated.Username,
		UpdatedAt: userUpdated.UpdatedAt,
	}, nil, err
}

// Delete User
func (us *UserService) DeleteUser(idToken uint) (*params.Response, error) {

	// call repository Delete user by id
	err := us.userRepo.DeleteUserById(idToken)
	if err != nil {
		return &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Message: "Your account has been successfully deleted",
	}, err
}
