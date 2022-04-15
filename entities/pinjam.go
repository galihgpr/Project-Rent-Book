package entities

import (
	"time"

	"gorm.io/gorm"
)

type Pinjam struct {
	gorm.Model
	IDBuku              uint
	IDUser              uint
	nameBuku            string
	TanggalPinjam       time.Time
	TanggalPengembalian time.Time
}
