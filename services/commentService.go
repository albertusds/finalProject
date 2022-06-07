package services

import (
	"finalproject/models"
	"finalproject/params"
	"finalproject/repositories"
	"net/http"
	"time"
)

type CommentService struct {
	commentRepo repositories.CommentRepository
	userRepo    repositories.UserRepository
	photoRepo   repositories.PhotoRepository
}

func NewCommentService(repo repositories.CommentRepository, userRepo repositories.UserRepository, photoRepo repositories.PhotoRepository) *CommentService {
	return &CommentService{
		commentRepo: repo,
		userRepo:    userRepo,
		photoRepo:   photoRepo,
	}
}

// Create comment
func (cs *CommentService) CreateComment(request params.CreateCommentRq, idToken uint) (*params.CreateCommentRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.CreateCommentRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	model := models.Comment{
		BaseModel: models.BaseModel{
			CreatedAt: time.Now(),
		},
		UserID:  idToken,
		PhotoID: request.PhotoId,
		Message: request.Message,
	}

	// call repository register user
	modelCommentResponse, err := cs.commentRepo.CreateComment(&model)
	if err != nil {
		return &params.CreateCommentRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	return &params.CreateCommentRs{
		Id:        modelCommentResponse.ID,
		PhotoId:   modelCommentResponse.PhotoID,
		UserId:    modelCommentResponse.UserID,
		CreatedAt: &modelCommentResponse.CreatedAt,
		Response:  params.Response{Message: modelCommentResponse.Message},
	}, err
}

func (cs *CommentService) DeleteComment(idToken, commentId uint) (*params.Response, error) {
	// call repository Delete comment by id
	err := cs.commentRepo.DeleteComment(idToken, commentId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Message: "Your comment has been successfully deleted",
	}, err
}

func (cs *CommentService) GetAllComments() (*[]params.GetCommentRs, error) {

	var result []params.GetCommentRs
	allCommentRs, err := cs.commentRepo.GetComment()
	if err != nil {
		errorOutput := params.GetCommentRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}

		result = append(result, errorOutput)

		return &result, err
	}

	for _, v := range *allCommentRs {
		//get user by user_idid
		userInfo, err := cs.userRepo.GetUserById(v.UserID)
		if err != nil {
			continue
		}

		// get photo by photo_id
		photoInfo, err := cs.photoRepo.GetPhotoById(v.PhotoID, v.UserID)
		if err != nil {
			continue
		}

		var temp = params.GetCommentRs{
			Id:        v.ID,
			PhotoId:   v.PhotoID,
			UserId:    v.UserID,
			CreatedAt: &v.CreatedAt,
			UpdatedAt: &v.UpdatedAt,
			User: params.UserCommentRs{
				Id:       userInfo.ID,
				Email:    userInfo.Email,
				Username: userInfo.Username},
			Photo: params.PhotoCommentRs{Id: photoInfo.ID,
				Title:    photoInfo.Title,
				Caption:  photoInfo.Caption,
				PhotoUrl: photoInfo.PhotoUrl,
				UserId:   photoInfo.UserID},
		}

		result = append(result, temp)
	}

	return &result, err
}

func (cs *CommentService) UpdateComment(request params.UpdateCommentRq, commentId, idToken uint) (*params.UpdateCommentRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.UpdateCommentRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	var comment = models.Comment{
		BaseModel: models.BaseModel{ID: commentId},
		Message:   request.Message,
	}

	err = cs.commentRepo.UpdateCommentById(&comment, idToken)
	if err != nil {
		return &params.UpdateCommentRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	commentModel, err := cs.commentRepo.GetCommentById(commentId, idToken)
	if err != nil {
		return &params.UpdateCommentRs{
			Response: params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Not Authorized Update",
			},
		}, err
	}

	photoModel, err := cs.photoRepo.GetPhotoById(commentModel.ID, idToken)
	if err != nil {
		return &params.UpdateCommentRs{
			Response: params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Not Authorized Update",
			},
		}, err
	}

	return &params.UpdateCommentRs{
		Id:        photoModel.ID,
		Title:     photoModel.Title,
		Caption:   photoModel.Caption,
		PhotoUrl:  photoModel.PhotoUrl,
		UserId:    commentModel.UserID,
		UpdatedAt: &commentModel.UpdatedAt,
	}, err
}
