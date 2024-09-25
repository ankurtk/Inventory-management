package models

type Product struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Quantity    *int    `json:"quantity"`
	Description string `json:"description"`
}
