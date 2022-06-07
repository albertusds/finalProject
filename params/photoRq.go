package params

type SavePhotoRq struct {
	Title    string `json:"title" valid:"required~title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required~photo url is required"`
}
