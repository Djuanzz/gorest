package models

import "time"

// AgeRatingCategory model
type AgeRatingCategory struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Movies      []Movie   `json:"-"`
}
