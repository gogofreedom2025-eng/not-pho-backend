package models

import "time"

type Beer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Brewery     string    `gorm:"not null" json:"brewery"`
	Style       string    `gorm:"not null" json:"style"`
	ABV         float64   `gorm:"not null" json:"abv"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
