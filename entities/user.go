package entities

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Nama     string
	Email    string   `gorm:"unique;not null"`
	HP       string   `gorm:"unique;not null"`
	Password string   `gorm:not null"`
	Buku     []Buku   `gorm:"foreignkey:UserID;references:id"`
	Pinjam   []Pinjam `gorm:"foreignkey:UserID;references:id"`
}
