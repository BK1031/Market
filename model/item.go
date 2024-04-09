package model

import "time"

type Item struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	StoreID     string    `json:"store_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
