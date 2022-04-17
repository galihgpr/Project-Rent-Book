package datastore

import (
	"intro-golang/entities"
	"time"

	"gorm.io/gorm"
)

func Pinjam(db *gorm.DB, buku entities.Buku, id uint) (string, error) {
	pjm := entities.Pinjam{}
	Buku := buku
	pjm.UserID = id
	pjm.BukuID = Buku.ID
	pjm.NameBuku = Buku.NameBuku
	pjm.TanggalPinjam = time.Now()
	pjm.TanggalPengembalian = time.Now().AddDate(0, 0, 10)

	res := db.Create(&pjm)
	if res.Error != nil {
		return "Buku gagal di pinjam", res.Error
	}
	var tambah int = 1
	Buku.Jumlah = Buku.Jumlah + tambah
	Buku.Status = false
	db.Save(&Buku)

	return "berhasil di tambahkan ke list pinjam anda", nil
}

func GetAllPinjam(db *gorm.DB) ([]entities.Pinjam, error) {
	res := []entities.Pinjam{}

	qry := db.Find(&res)

	if qry.Error != nil {
		return nil, qry.Error
	}

	return res, nil
}

func Kembalikan(db *gorm.DB, id uint, idbuku uint) (string, error) {
	hps := db.Delete(&entities.Pinjam{}, id)
	if hps.Error != nil {
		return "gagal mengembalikan buku", hps.Error
	}
	db.Model(&entities.Buku{}).Where("id", idbuku).Update("status", true)

	return "berhasil mengembalikan buku", nil
}
