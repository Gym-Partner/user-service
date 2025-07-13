package domain

import (
	"github.com/Gym-Partner/api_common/serviceError"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string     `json:"id" swaggerignore:"true"`
	FirstName string     `json:"first_name" example:"Oscar"`
	LastName  string     `json:"last_name" example:"Escorneboueu"`
	Username  string     `json:"username" example:"don_oscar_anton"`
	Email     string     `json:"email" example:"titouan.esc@icloud.com"`
	Phone     string     `json:"phone" example:"+33672135172"`
	Password  string     `json:"password" example:"aaaAAA111"`
	Followers []string   `json:"followers,omitempty" gorm:"-"`
	Following []string   `json:"following,omitempty" gorm:"-"`
	Image     string     `json:"image,omitempty" gorm:"-"`
	CreatedAt *time.Time `json:"created_at"`
}
type Users []User

func (u *User) Response() gin.H {
	return gin.H{
		"data": gin.H{
			"id":         u.ID,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"username":   u.Username,
			"email":      u.Email,
			"phone":      u.Phone,
			"followers":  u.Followers,
			"following":  u.Following,
			"image":      u.Image,
		},
	}
}

func (u *Users) Response() gin.H {
	var result []gin.H

	for _, user := range *u {
		result = append(result, user.Response())
	}

	return gin.H{
		"data": result,
	}
}

func (u *User) GenerateId() {
	u.ID = uuid.New().String()
}

func (u *User) HashPassword(hashFunc func(string) (string, *serviceError.Error)) *serviceError.Error {
	hashed, err := hashFunc(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

// ----------------------------- SQL -----------------------------

type MigrateUser struct {
	ID        string `gorm:"primaryKey; not null"`
	FirstName string
	LastName  string
	Username  string `gorm:"column:username; not null"`
	Email     string `gorm:"not null"`
	Phone     string
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (MigrateUser) TableName() string { return "users" }
