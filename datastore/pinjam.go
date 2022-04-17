package datastore

import (
	"errors"
	"fmt"
	"intro-golang/entities"
	"time"

	"gorm.io/gorm"
)

func Pinjam(db *gorm.DB, buku entities.DetailBuku, id uint) (string, error) {
	book := entities.Buku{}
	if buku.Status == true {
		err := errors.New("Masih di Pinjam Oleh Orang Lain")
		return "Buku Gagal di Pinjam", err
	}
	db.Table("bukus").Where("name_buku=?", buku.NameBuku).Updates(entities.Buku{Status: true, Jumlah: buku.Jumlah + 1})
	db.Where("name_buku=?", buku.NameBuku).Find(&book)
	fmt.Println(book.ID)
	Buku := buku
	pjm := entities.Pinjam{
		UserID:              id,
		BukuID:              book.ID,
		NameBuku:            Buku.NameBuku,
		TanggalPinjam:       time.Now(),
		TanggalPengembalian: time.Now().AddDate(0, 0, 10),
	}

	res := db.Create(&pjm)
	if res.Error != nil {
		return "Buku Gagal di Pinjam", res.Error
	}

	return "Berhasil di Tambahkan ke List Pinjam Anda", nil
}

func GetAllPinjam(db *gorm.DB, id uint) ([]entities.Pinjam, error) {
	res := []entities.Pinjam{}

	qry := db.Where("user_id=? and aktif=false", id).Find(&res)

	if qry.Error != nil {
		return nil, qry.Error
	}
	if len(res) == 0 {
		fmt.Println("TIDAK ADA BUKU YANG DIPINJAM")
	}
	return res, nil
}

func Kembalikan(db *gorm.DB, id uint, idbuku uint) (string, error) {
	if err := db.Table("pinjams").Where("buku_id=?", idbuku).Update("aktif", true).Error; err != nil {
		return "Pengembalian Buku Gagal", err
	}
	db.Model(&entities.Buku{}).Where("id", idbuku).Update("status", false)

	return "Berhasil Mengembalikan Buku", nil
}
