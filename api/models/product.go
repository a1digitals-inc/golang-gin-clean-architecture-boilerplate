package models

import "time"

type Product struct {
	ID           int       `json:"id"`
	ProductTitle string    `json:"title"`
	Price        string    `json:"price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Deleted_At   time.Time `json:"deleted_at"`
}

func (Product) TableName() string {
	return "product"
}
