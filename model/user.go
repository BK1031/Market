package model

import "time"

type User struct {
	ID                string    `gorm:"primaryKey" json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `gorm:"unique" json:"email"`
	Password          string    `json:"password"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}
