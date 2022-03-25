package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.Id = uuid.NewV4()
	return nil
}
