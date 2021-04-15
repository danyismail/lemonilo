package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"email"`
	Address  string `json:"address,omitempty"`
	Password string `json:"password" validate:"required"`
}
