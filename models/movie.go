package models

import "time"

// Movie model
type Movie struct {
	Id                  int               `json:"id" gorm:"primary_key"`
	Title               string            `json:"title"`
	Year                int               `json:"year"`
	AgeRatingCategoryId int               `json:"age_rating_category_id"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	AgeRatingCategory   AgeRatingCategory `json:"-"`
}
