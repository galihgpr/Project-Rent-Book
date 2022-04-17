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

//METHOD LIST BUKU
func (b *BukuDB) ListBuku() ([]entities.Buku, error) {
	buku := []entities.Buku{}

	if err := b.Db.Order("name_buku").Find(&buku).Error; err != nil {
		ErrCek := errors.New("AKSES KE DATABASE GAGAL")
		return []entities.Buku{}, ErrCek
	}

	return buku, nil
}

//METHOD GET BUKU
func (b *BukuDB) DetailBuku(buku string, user uint) (entities.DetailBuku, error) {
	result := entities.DetailBuku{}
	hasil := b.Db.Table("users u").Joins("LEFT JOIN bukus b on u.id=b.user_id").Where("u.id=? AND b.name_buku =?", user, buku).Scan(&result)
	if hasil.Error != nil {
		fmt.Println()
		ErrCek := errors.New("AKSES KE DATABASE GAGAL")
		return entities.DetailBuku{}, ErrCek
	}
	return result, nil
}

//METHOD UPDATE BUKU

func (b *BukuDB) UpdateBuku(namalama string, namabaru string) (string, error) {
	update := b.Db.Table("bukus").Where("name_buku =?", namalama).Update("name_buku", namabaru)
	if update.Error != nil {
		return namalama, update.Error
	}
	return namabaru, nil
}
