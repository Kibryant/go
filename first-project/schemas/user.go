package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Salary    int64     `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
