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
		ErrCek := errors.New("ERROR AKSES DATABASE USER")
		return entities.Users{}, ErrCek
	}
	if CekLogin.Email == "" || CekLogin.StatusAkun == true {
		err := errors.New("EMAIL ATAU PASSWORD SALAH")
		return entities.Users{}, err
	}
	return CekLogin, nil
}

// REGISTER AKUN
func (u *UserDB) Register(newUser entities.Users) (entities.Users, error) {
	user := entities.Users{}

	if err := u.Db.Where("email = ?", newUser.Email).Find(&user).Error; err != nil {
		ErrCek := errors.New("ERROR AKSES DATABASE USER")
		return entities.Users{}, ErrCek
	}
	if user.Email != "" {
		err := errors.New("REGISTRASI GAGAL, EMAIL '" + newUser.Email + "' TELAH TERDAFTAR")
		return entities.Users{}, err
	} else {
		if err := u.Db.Create(&newUser).Error; err != nil {
			ErrCek := errors.New("REGISTRASI GAGAL")
			return newUser, ErrCek
		}
	}
	return newUser, nil
}

//EDIT PROFIL

func UpdateUser(db *gorm.DB, updatedUser entities.Users) (entities.Users, error, string) {
	qry := db.Save(&updatedUser)

	if qry.Error != nil {
		return entities.Users{}, qry.Error, "UPDATE PROFIL GAGAL"
	}

	return updatedUser, nil, "UPDATE PROFIT BERHASIL"
}
