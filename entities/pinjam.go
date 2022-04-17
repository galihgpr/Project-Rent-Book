package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pinjam struct {
	gorm.Model
	BukuID              uint
	UserID              uint
<<<<<<< HEAD
=======
	nameBuku            string `gorm:"not null"`
>>>>>>> List-Buku
	TanggalPinjam       time.Time
	TanggalPengembalian time.Time
}
