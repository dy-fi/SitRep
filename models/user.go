package models

import (
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	ID			string
	username 	string
	rooms		[]Room
}