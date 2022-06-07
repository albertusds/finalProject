package controllers

import (
	"finalproject/params"
	"finalproject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

// controller for register user
func (uc *UserController) RegisterUser(c *gin.Context) {
	var req params.RegisterUserRq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := uc.userService.RegisterUser(req)
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// controller for user login
func (uc *UserController) UserLogin(c *gin.Context) {
	var req params.UserLoginRq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := uc.userService.UserLogin(req)
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var req params.UpdateUserRq
	idToken := c.MustGet("id").(float64)
	emailToken := c.MustGet("email").(string)
	userIdString := c.Param("userId")

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	responseSuccess, responseError, err := uc.userService.UpdateUser(req, uint(userId), uint(idToken), emailToken)
	if err != nil {
		c.AbortWithStatusJSON(responseError.Status, responseError)
		return
	}

	c.JSON(http.StatusOK, responseSuccess)
}

// Delete user
func (uc *UserController) DeleteUser(c *gin.Context) {
	idToken := c.MustGet("id").(float64)

	response, err := uc.userService.DeleteUser(uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
