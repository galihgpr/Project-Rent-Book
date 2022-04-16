package datastore

import (
	"errors"
	"fmt"
	"intro-golang/entities"
	"strings"

	"gorm.io/gorm"
)

type BukuDB struct {
	Db *gorm.DB
}

func (b *BukuDB) TambahBuku(NewBuku entities.Buku, user entities.Users) (entities.Buku, error) {
	buku := entities.Buku{}
	NewBuku.Author = strings.ToUpper(user.Nama[:2])
	NewBuku.UserID = user.ID

	if err := b.Db.Where("name_buku = ?", NewBuku.NameBuku).Find(&buku).Error; err != nil {
		fmt.Println()
		ErrCek := errors.New("AKSES KE DATABASE GAGAL")
		return entities.Buku{}, ErrCek
	}
	if buku.NameBuku != "" {
		fmt.Println()
		err := errors.New("BUKU GAGAL DI BUAT, BUKU" + buku.NameBuku + " SUDAH ADA")
		return entities.Buku{}, err
	} else {
		if err := b.Db.Create(&NewBuku).Error; err != nil {
			fmt.Println()
			ErrCek := errors.New("Pendaftaran Gagal")
			return entities.Buku{}, ErrCek
		}
	}
	return NewBuku, nil
}

func (b *BukuDB) ListBuku(user entities.Users) (entities.Buku, error) {
	listbuku := entities.Buku{}
	buku := []entities.Buku{}

	if err := b.Db.Find(&listbuku).Error; err != nil {
		ErrCek := errors.New("AKSES KE DATABASE GAGAL")
		return entities.Buku{}, ErrCek
	}
	if err := b.Db.Select("name_buku").Order("name_buku").Find(&buku).Error; err != nil {
		ErrCek := errors.New("AKSES KE DATABASE GAGAL")
		return entities.Buku{}, ErrCek
	}
	for i, v := range buku {
		fmt.Println(i+1, " ", strings.ToUpper(v.NameBuku))
	}
	fmt.Println("99 Kembali Ke Menu Sebelumnya")
	return listbuku, nil
}
