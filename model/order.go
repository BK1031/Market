package model

import "time"

type Order struct {
	ID        string        `gorm:"primaryKey" json:"id"`
	UserID    string        `json:"user_id"`
	Items     []OrderItem   `gorm:"-" json:"items"`
	Price     int           `json:"price"`
	Shipping  OrderShipping `gorm:"-" json:"shipping"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
}

type OrderItem struct {
	OrderID  string `gorm:"primaryKey" json:"order_id"`
	ItemID   string `gorm:"primaryKey" json:"item_id"`
	Item     Item   `gorm:"-" json:"item"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type OrderShipping struct {
	OrderID      string `gorm:"primaryKey" json:"order_id"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
}
