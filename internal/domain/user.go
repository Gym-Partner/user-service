package domain

import (
	"github.com/Gym-Partner/api-common/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"reflect"
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
	data := gin.H{
		"id":        u.ID,
		"firstName": u.FirstName,
		"lastName":  u.LastName,
		"username":  u.Username,
		"email":     u.Email,
	}
	addIfNotEmpty(data, "phone", u.Phone)
	addIfNotEmpty(data, "followers", u.Followers)
	addIfNotEmpty(data, "following", u.Following)
	addIfNotEmpty(data, "image", u.Image)

	return gin.H{
		"data": data,
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

func (u *User) HashPassword(hashFunc func(string) (string, *errs.Error)) *errs.Error {
	hashed, err := hashFunc(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

func addIfNotEmpty(m gin.H, key string, value any) {
	if value == nil {
		return
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Ptr, reflect.Interface:
		if val.IsNil() {
			return
		}
		val = val.Elem()
	default:
	}

	switch val.Kind() {
	case reflect.String:
		if val.Len() == 0 {
			return
		}
	case reflect.Slice, reflect.Map:
		if val.Len() == 0 {
			return
		}
	case reflect.Invalid:
		return
	default:
	}

	m[key] = value
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
