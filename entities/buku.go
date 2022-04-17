package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	NameBuku string
	Author   string `gorm:"not null"`
	Jumlah   int    `gorm:"default:0"`
	Status   bool
	Aktif    bool
	UserID   uint
	Pinjam   Pinjam `gorm:"foreignkey:BukuID;references:id"`
}

type DetailBuku struct {
	NameBuku string
	Nama     string
	UserID   uint
	Status   bool
	Jumlah   int
	Author   string
}
