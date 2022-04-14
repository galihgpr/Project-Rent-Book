package user

import (
	"fmt"
	"intro-golang/entities"

	"gorm.io/gorm"
)

type UserDB struct {
	Db *gorm.DB
}

func (u *UserDB) Register(newUser entities.Users) (entities.Users, error) {
	u.Db.AutoMigrate(&entities.Users{})
	if err := u.Db.Create(&newUser).Error; err != nil {
		fmt.Println("Terjadi kesalahan saat insert user", err)
		return newUser, err
	}
	return newUser, nil
}
