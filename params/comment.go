package params

type CreateComment struct {
	Message string `json:"message" valid:"required~Field Title is required"`
	PhotoID uint   `json:"photo_id" valid:"required~Field Photo Id is required"`
}

type UpdateComment struct {
	ID      uint   `json:"id" valid:"required~Field Title is required"`
	Message string `json:"message" valid:"required~Field Title is required"`
	PhotoID uint   `json:"photo_id" valid:"required~Field Photo Id is required"`
}
