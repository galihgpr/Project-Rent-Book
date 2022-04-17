package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Nama     string
	Email    string
	Password string
	Hp       string
}