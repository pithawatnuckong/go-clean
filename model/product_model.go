package model

import "time"

// Requests

type ProductCreateOrUpdateModel struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int32   `json:"quantity"`
	OwnerID  int     `json:"ownerId"`
}

// Responses

type ProductModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Quantity  int32     `json:"quantity"`
	OwnerID   int       `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
