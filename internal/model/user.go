package model

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"username" db:"username" binding:"required"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password" db:"hashed_password" binding:"required"`
}
