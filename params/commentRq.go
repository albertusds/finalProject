package params

type CreateCommentRq struct {
	PhotoId uint   `json:"photo_id" valid:"required~photo_id is required"`
	Message string `json:"message" valid:"required~message is required"`
}

type UpdateCommentRq struct {
	Message string `json:"message" valid:"required~message is required"`
}
