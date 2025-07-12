package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Users []User

func (u *User) Response() gin.H {
	return gin.H{
		"data": gin.H{
			"username": u.Username,
			"email":    u.Email,
		},
	}
}

func (u *User) GenerateId() {
	u.ID = uuid.New().String()
}

type MigrateUser struct {
	ID       string `gorm:"primaryKey; not null"`
	Username string `gorm:"column:username; not null"`
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func (MigrateUser) TableName() string { return "users" }
