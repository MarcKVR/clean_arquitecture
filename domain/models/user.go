package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"id"`
	Names    string `gorm:"not null"`
	Paternal string `gorm:"not null"`
	Maternal string
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return
}
