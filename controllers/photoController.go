package controllers

import (
	"finalproject/params"
	"finalproject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoService services.PhotoService
}

func NewPhotoController(service *services.PhotoService) *PhotoController {
	return &PhotoController{
		photoService: *service,
	}
}

func (pc *PhotoController) SavePhoto(c *gin.Context) {
	var req params.SavePhotoRq
	idToken := c.MustGet("id").(float64)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := pc.photoService.SavePhoto(req, uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (pc *PhotoController) GetPhotos(c *gin.Context) {

	responseSuccess, responseError, err := pc.photoService.GetPhotos()
	if err != nil {
		c.AbortWithStatusJSON(responseError.Status, responseError)
		return
	}

	c.JSON(http.StatusOK, responseSuccess)
}

func (pc *PhotoController) DeletePhoto(c *gin.Context) {
	idToken := c.MustGet("id").(float64)
	photoIdString := c.Param("photoId")

	photoId, err := strconv.Atoi(photoIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := pc.photoService.DeletePhoto(uint(idToken), uint(photoId))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (pc *PhotoController) UpdatePhoto(c *gin.Context) {
	var req params.SavePhotoRq
	idToken := c.MustGet("id").(float64)
	photoIdString := c.Param("photoId")

	photoId, err := strconv.Atoi(photoIdString)
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

	responseSuccess, responseError, err := pc.photoService.UpdatePhoto(req, uint(photoId), uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(responseError.Status, responseError)
		return
	}

	c.JSON(http.StatusOK, responseSuccess)
}
