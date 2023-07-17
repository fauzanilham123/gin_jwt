package models

import "time"

type (
	// AgeRatingCategory
	Genre struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Genre     string    `json:"genre"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)