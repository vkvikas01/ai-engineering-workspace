package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Name      string `gorm:"size:100;not null" json:"name"`

	Email     string `gorm:"size:255;uniqueIndex;not null" json:"email"`

	Password  string `gorm:"size:255;not null" json:"-"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}