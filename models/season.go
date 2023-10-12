package models

type Season struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateSeasonInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}
