package user

import (
	// "intro-golang/entities"

	"intro-golang/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) Register(newUser entities.Users) (string, error) {
	user := []entities.Users{}

	if err := u.Db.Table("users").Where("email = ?", newUser.Email).Find(&user).Error; err != nil {
		return "Akses Database User Gagal", err
	}
	if len(user) != 0 {
		return "Registrasi Gagal, email " + newUser.Email + " telah terdaftar", nil
	} else {
		if err := u.Db.Create(&newUser).Error; err != nil {
			return "Pendaftaran Gagal", err
		}
	}
	return "Pendaftaran Berhasil", nil
}
