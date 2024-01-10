package model

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"firstname" db:"firstname" binding:"required"`
	LastName  string `json:"lastname" db:"lastname" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
}
