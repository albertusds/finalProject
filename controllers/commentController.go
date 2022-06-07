package controllers

import (
	"finalproject/params"
	"finalproject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(service *services.CommentService) *CommentController {
	return &CommentController{
		commentService: *service,
	}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var req params.CreateCommentRq
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

	response, err := cc.commentService.CreateComment(req, uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	idToken := c.MustGet("id").(float64)
	commentIdString := c.Param("commentId")

	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		})
		return
	}

	response, err := cc.commentService.DeleteComment(uint(idToken), uint(commentId))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cc *CommentController) GetComments(c *gin.Context) {

	response, err := cc.commentService.GetAllComments()
	if err != nil {
		c.AbortWithStatusJSON((*response)[0].Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var req params.UpdateCommentRq
	idToken := c.MustGet("id").(float64)
	commentIdString := c.Param("commentId")

	commentId, err := strconv.Atoi(commentIdString)
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

	response, err := cc.commentService.UpdateComment(req, uint(commentId), uint(idToken))
	if err != nil {
		c.AbortWithStatusJSON(response.Status, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
