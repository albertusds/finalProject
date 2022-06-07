package services

import (
	"finalproject/models"
	"finalproject/params"
	"finalproject/repositories"
	"net/http"
	"time"
)

type PhotoService struct {
	photoRepo repositories.PhotoRepository
	userRepo  repositories.UserRepository
}

func NewPhotoService(repo repositories.PhotoRepository, usrRepo repositories.UserRepository) *PhotoService {
	return &PhotoService{
		photoRepo: repo,
		userRepo:  usrRepo,
	}
}

func (ps *PhotoService) SavePhoto(request params.SavePhotoRq, userId uint) (*params.SavePhotoRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.SavePhotoRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	var photo = models.Photo{
		BaseModel: models.BaseModel{CreatedAt: time.Now()},
		Title:     request.Title,
		Caption:   request.Caption,
		PhotoUrl:  request.PhotoUrl,
		UserID:    userId,
	}

	photoModelResponse, err := ps.photoRepo.SavePhoto(&photo)
	if err != nil {
		return &params.SavePhotoRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	return &params.SavePhotoRs{
		Id:        photoModelResponse.ID,
		Title:     photoModelResponse.Title,
		Caption:   photoModelResponse.Caption,
		PhotoUrl:  photoModelResponse.PhotoUrl,
		UserId:    photoModelResponse.UserID,
		CreatedAt: &photoModelResponse.CreatedAt,
	}, err
}

func (ps *PhotoService) GetPhotos() (*[]params.GetPhotosRs, *params.Response, error) {

	photoModelArrRs, err := ps.photoRepo.GetPhotos()
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	var Result []params.GetPhotosRs

	for _, v := range *photoModelArrRs {
		userInfo, err := ps.userRepo.GetUserById(v.UserID)
		if err != nil {
			continue
		}
		var temp = params.GetPhotosRs{
			Id:        v.ID,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserId:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User:      params.UserPhotosRs{Email: userInfo.Email, Username: userInfo.Username},
		}

		Result = append(Result, temp)
	}

	return &Result, nil, err
}

func (ps *PhotoService) DeletePhoto(idToken uint, photoId uint) (*params.Response, error) {
	// call repository Delete photo by id
	err := ps.photoRepo.DeletePhoto(idToken, photoId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Message: "Your photo has been successfully deleted",
	}, err
}

func (ps *PhotoService) UpdatePhoto(request params.SavePhotoRq, photoId, idToken uint) (*params.UpdatePhotoRs, *params.Response, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Error:   err.Error(),
		}, err
	}

	var photo = models.Photo{
		BaseModel: models.BaseModel{ID: photoId},
		Title:     request.Title,
		Caption:   request.Caption,
		PhotoUrl:  request.PhotoUrl,
	}

	err = ps.photoRepo.UpdatePhotoById(&photo, idToken)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	// call to get photo by id
	photoModelResponse, err := ps.photoRepo.GetPhotoById(photoId, idToken)
	if err != nil {
		return nil, &params.Response{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Error:   "Not Authorized Update with given photoId",
		}, err
	}

	return &params.UpdatePhotoRs{
		Id:        photoModelResponse.ID,
		Title:     photoModelResponse.Title,
		Caption:   photoModelResponse.Caption,
		PhotoUrl:  photoModelResponse.PhotoUrl,
		UserId:    photoModelResponse.UserID,
		UpdatedAt: photoModelResponse.UpdatedAt,
	}, nil, err
}
