package models

import (
	"time"
)

type User struct {
	ID       uint32 `gorm:"primary_key;column:id" json:"-"`
	Email    string `gorm:"column:email;unique" json:"email,omitempty"`
	Name     string `gorm:"column:name" json:"name,omitempty"`
	Username string `gorm:"column:username;unique" json:"username,omitempty"`
	Password string `gorm:"column:password" json:"-"`
	Address  string `gorm:"column:address;type:text" json:"address,omitempty"`

	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}
