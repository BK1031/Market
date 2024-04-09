package model

import "time"

type Store struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	LogoURL     string    `json:"logo_url"`
	BannerURL   string    `json:"banner_url"`
	Items       []Item    `gorm:"-" json:"items"`
	Admins      []User    `gorm:"-;" json:"admins"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
