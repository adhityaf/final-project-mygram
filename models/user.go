package models

import "time"

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Email    string `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	Username string `gorm:"type:varchar(255);unique_index;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Age      int    `gorm:"type:int;not null" json:"age"`

	// User Has Many Photos
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many SocialMedias
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User Has Many Comments
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` 

	// DROP TABLE users, comments, photos, social_media;
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

