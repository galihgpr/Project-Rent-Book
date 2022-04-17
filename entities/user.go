package entities

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
<<<<<<< HEAD

=======
	Nama     string   `gorm:"not null"`
	Email    string   `gorm:"not null"`
	HP       string   `gorm:"not null"`
	Password string   `gorm:"not null"`
>>>>>>> List-Buku
	Buku     []Buku   `gorm:"foreignkey:UserID;references:id"`
	Pinjam   []Pinjam `gorm:"foreignkey:UserID;references:id"`
}
