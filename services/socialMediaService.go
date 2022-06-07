package services

import (
	"finalproject/models"
	"finalproject/params"
	"finalproject/repositories"
	"net/http"
	"time"
)

type SocialMediaService struct {
	socialMediaRepo repositories.SocialMediaRepository
	userRepo        repositories.UserRepository
}

func NewSocialMediaService(repo repositories.SocialMediaRepository, userRepo repositories.UserRepository) *SocialMediaService {
	return &SocialMediaService{
		socialMediaRepo: repo,
		userRepo:        userRepo,
	}
}

func (ss *SocialMediaService) CreateSocialMedia(request params.SocMedRq, idToken uint) (*params.CreateSocMedRs, error) {
	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.CreateSocMedRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	var socmed = models.SocialMedia{
		BaseModel:      models.BaseModel{CreatedAt: time.Now()},
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
		UserID:         idToken,
	}

	socmedModelResponse, err := ss.socialMediaRepo.CreateSocMed(&socmed)
	if err != nil {
		return &params.CreateSocMedRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	return &params.CreateSocMedRs{
		Id:             socmedModelResponse.ID,
		Name:           socmedModelResponse.Name,
		SocialMediaUrl: socmedModelResponse.SocialMediaUrl,
		UserId:         socmedModelResponse.UserID,
		CreatedAt:      &socmedModelResponse.CreatedAt,
	}, err
}

func (ss *SocialMediaService) DeleteSocmed(idToken uint, photoId uint) (*params.Response, error) {
	// call repository Delete sosmed by id
	err := ss.socialMediaRepo.DeleteSocMed(idToken, photoId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Message: "Your social media has been successfully deleted",
	}, err
}

func (ss *SocialMediaService) GetSocialMedias() (*params.GetSocMedRs, error) {

	var socmedItemsArrRs []params.GetSocMedItemRs
	var result params.GetSocMedRs

	socialMedias, err := ss.socialMediaRepo.GetSocialMedias()
	if err != nil {
		return &params.GetSocMedRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	for _, v := range *socialMedias {
		userInfo, err := ss.userRepo.GetUserById(v.UserID)
		if err != nil {
			continue
		}
		var temp = params.GetSocMedItemRs{
			Id:             v.ID,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         v.UserID,
			CreatedAt:      &v.CreatedAt,
			UpdatedAt:      &v.UpdatedAt,
			User:           params.UserSocMedRs{Id: userInfo.ID, Username: userInfo.Username},
		}

		socmedItemsArrRs = append(socmedItemsArrRs, temp)
	}

	result = params.GetSocMedRs{
		SocialMedias: socmedItemsArrRs,
	}

	return &result, err
}

func (ss *SocialMediaService) UpdateSocmed(request params.SocMedRq, socmedId, idToken uint) (*params.UpdateSocMedRs, error) {

	//validate struct request
	err := params.ValidateRq(request)
	if err != nil {
		return &params.UpdateSocMedRs{
			Response: params.Response{
				Status:  http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Error:   err.Error(),
			},
		}, err
	}

	var socmed = models.SocialMedia{
		BaseModel:      models.BaseModel{ID: socmedId},
		Name:           request.Name,
		SocialMediaUrl: request.SocialMediaUrl,
	}

	err = ss.socialMediaRepo.UpdateSocialMediaById(&socmed, idToken)
	if err != nil {
		return &params.UpdateSocMedRs{
			Response: params.Response{
				Status:  http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
				Error:   err.Error(),
			},
		}, err
	}

	// call to get socmed by id
	socmedModel, err := ss.socialMediaRepo.GetSocialMediaById(socmedId, idToken)
	if err != nil {
		return &params.UpdateSocMedRs{
			Response: params.Response{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
				Error:   "Not Authorized Update with given socmedId",
			},
		}, err
	}

	return &params.UpdateSocMedRs{
		Id:             socmedModel.ID,
		Name:           socmedModel.Name,
		SocialMediaUrl: socmedModel.SocialMediaUrl,
		UserId:         socmedModel.UserID,
		UpdatedAt:      &socmedModel.UpdatedAt,
	}, err
}
