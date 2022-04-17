package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pinjam struct {
	gorm.Model
	BukuID              uint
	UserID              uint
	NameBuku            string `gorm:"not null"`
	TanggalPinjam       time.Time
	TanggalPengembalian time.Time
}
