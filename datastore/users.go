package datastore

import (
	"fmt"
	"intro-golang/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) GetAllDataUser() (string, error) {
	res := []entities.Users{}

	if err := u.Db.Table("users").Find(&res).Error; err != nil {
		fmt.Println("Terjadi kesalahan memasukkan user", err)
		return newUser.Email, err
	}
	return "Pendaftaran Berhasil", nil
}
