package models

import "time"

type Photo struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `gorm:"type:varchar(255);not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `gorm:"type:varchar(255);not null" json:"photo_url"`

	// Photo Has One User
	UserID uint `gorm:"type:int;not null;" json:"user_id"`
	// Photo Has Many Comments
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
