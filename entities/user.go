package entities

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	HP       string
	NamaUser string
	Password string
	Buku     []Buku   `gorm:"foreignKey:IDUser"`
	Pinjam   []Pinjam `gorm:"foreignKey:IDUser"`
}
