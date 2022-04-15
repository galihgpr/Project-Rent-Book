package entities

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	NameBuku string
	Author   string
	Status   bool
	IDUser   uint
	Pinjam   []Pinjam `gorm:"foreignKey:IDBuku"`
}
