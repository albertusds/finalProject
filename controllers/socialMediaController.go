package controllers

import (
	"finalproject/params"
	"finalproject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	socialMediaService services.SocialMediaService
}

func NewSocialMediaController(service *services.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		socialMediaService: *service,
	}
}

func (sc *SocialMediaController) CreateSocialMedia(c *gin.Context) {
	var req params.SocMedRq
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

	response, err := sc.socialMediaService.CreateSocialMedia(req, uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (sc *SocialMediaController) DeleteSocialMedia(c *gin.Context) {
	idToken := c.MustGet("id").(float64)
	socmedIdString := c.Param("socialMediaId")

	socmedId, err := strconv.Atoi(socmedIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := sc.socialMediaService.DeleteSocmed(uint(idToken), uint(socmedId))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (sc *SocialMediaController) GetSocmed(c *gin.Context) {

	response, err := sc.socialMediaService.GetSocialMedias()
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (sc *SocialMediaController) UpdateSocmed(c *gin.Context) {
	var req params.SocMedRq
	idToken := c.MustGet("id").(float64)
	socmedIdString := c.Param("socialMediaId")

	socmedId, err := strconv.Atoi(socmedIdString)
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

	response, err := sc.socialMediaService.UpdateSocmed(req, uint(socmedId), uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
