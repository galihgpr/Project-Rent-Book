package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	NameBuku string `gorm:"unique;not null"`
	Author   string `gorm:not null"`
	Jumlah   int    `gorm:default:0`
	UserID   uint
	Pinjam   Pinjam `gorm:"foreignkey:BukuID;references:id`
}
