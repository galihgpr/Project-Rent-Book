package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pinjam struct {
	gorm.Model
	BukuID              uint
	UserID              uint
	nameBuku            string
	TanggalPinjam       time.Time
	TanggalPengembalian time.Time
}
