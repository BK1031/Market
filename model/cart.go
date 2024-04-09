package model

import "time"

type CartItem struct {
	ItemID    string    `gorm:"primaryKey" json:"item_id"`
	UserID    string    `gorm:"primaryKey" json:"user_id"`
	Item      Item      `gorm:"-" json:"item"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
