package model

import "time"

type Tradition struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Date        time.Time `json:"date" db:"date" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
}
