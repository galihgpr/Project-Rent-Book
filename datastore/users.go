package datastore

import (
	"errors"
	"intro-golang/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

//LOGIN AKUN
func (u *UserDB) LoginAkun(newLogin entities.Users) (entities.Users, error) {
	CekLogin := entities.Users{}

	if err := u.Db.Where("email = ? AND password = ?", newLogin.Email, newLogin.Password).Find(&CekLogin).Error; err != nil {
		ErrCek := errors.New("Error Akses Ke Database")
		return entities.Users{}, ErrCek
	}
	if CekLogin.Email == "" {
		err := errors.New("Email Atau Password Salah")
		return entities.Users{}, err
	}

	return CekLogin, nil
}

// REGISTER AKUN
func (u *UserDB) Register(newUser entities.Users) (entities.Users, error) {
	user := entities.Users{}

	if err := u.Db.Where("email = ?", newUser.Email).Find(&user).Error; err != nil {
		ErrCek := errors.New("Akses Database User Gagal")
		return entities.Users{}, ErrCek
	}
	if user.Email != "" {
		err := errors.New("Registrasi Gagal, email " + newUser.Email + " telah terdaftar")
		return entities.Users{}, err
	} else {
		if err := u.Db.Create(&newUser).Error; err != nil {
			ErrCek := errors.New("Pendaftaran Gagal")
			return newUser, ErrCek
		}
	}
	return newUser, nil
}

//EDIT PROFIL

func UpdateUser(db *gorm.DB, updatedUser entities.Users) (entities.Users, error, string) {
	qry := db.Save(&updatedUser)

	if qry.Error != nil {
		return entities.Users{}, qry.Error, "Edit Profil Gagal"
	}

	return updatedUser, nil, "Update Profil Berhasil"
}
