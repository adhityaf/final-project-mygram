package params

type CreateComment struct {
	Message string `json:"message" valid:"required~Field Message is required"`
	PhotoID int `json:"photo_id" valid:"required~Field Photo Id is required"`
	UserID  uint   `json:"user_id" valid:"required~Field User Id is required"`
}

type UpdateComment struct {
	ID      uint   `json:"id" valid:"required~Field ID is required"`
	Message string `json:"message" valid:"required~Field Message is required"`
	UserID  uint   `json:"user_id" valid:"required~Field User Id is required"`
}
