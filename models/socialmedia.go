package models

type SocialMedia struct {
	ID             uint   `gorm:"primary_key" json:"id"`
	Name           string `gorm:"type:varchar(255);not null" json:"name"`
	SocialMediaURL string `gorm:"type:varchar(255);not null" json:"social_media_url"`

	// SocialMedia Has One User
	UserID uint `gorm:"type:int;not null;" json:"user_id"`
}
