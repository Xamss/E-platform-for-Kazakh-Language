package model

type Activity struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
	Type string `json:"type" db:"type" binding:"required"`
}
