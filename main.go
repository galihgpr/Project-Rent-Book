package main

import (
	"fmt"
	"intro-golang/config"
	"intro-golang/datastore"
	"intro-golang/entities"
	"strings"
)

func TampilanMenuUtama() {
	// MENU UTAMA
	Mymenu := []string{"1. Login", "2. Register", "3. Lis Buku", "99. Keluar"}
	fmt.Println("===== Menu Utama =====")
	for _, v := range Mymenu {
		fmt.Println(v)
	}
	var nomor int
	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%d", &nomor)
	fmt.Println("\n")
	// PILIH MENU
	switch nomor {
	case 1:
		Login()
	case 2:
		Register()
	case 3:
		// ListBuku()
	case 99:
		fmt.Println("===== Terimakasih Sudah Berkunjung =====")
		break
	default:
		TampilanMenuUtama()
	}
}

//Menu Login
func Login() {
	fmt.Println("===== Login Akun =====")
	conn := config.ConnectDB()
	user := datastore.UserDB{Db: conn}
	var email, password string
	fmt.Print("Masukkan Email : ")
	fmt.Scanf("%s", &email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanf("%s", &password)
	masuk, err := user.LoginAkun(entities.Users{Email: email, Password: password})
	if err != nil {
		fmt.Println(err)
		TampilanMenuUtama()
	} else {
		fmt.Println()
		MenuApps(masuk)
	}
}

//Menu Register
func Register() {
	fmt.Println("===== Registrasi Akun =====")
	conn := config.ConnectDB()
	user := datastore.UserDB{Db: conn}
	var nama, email, password, hp string
	fmt.Print("Masukkan Nama : ")
	fmt.Scanf("%s", &nama)
	fmt.Print("Masukkan Email : ")
	fmt.Scanf("%s", &email)
	fmt.Print("Masukkan Password : ")
	fmt.Scanf("%s", &password)
	fmt.Print("Masukkan No HP : ")
	fmt.Scanf("%s", &hp)
	daftar, err := user.Register(entities.Users{Nama: nama, Email: email, Password: password, HP: hp})
	if err != nil {
		fmt.Println(err)
		TampilanMenuUtama()
	}
	fmt.Println()
	MenuApps(daftar)

}

//Menu Apps

func MenuApps(user entities.Users) {
	Mymenu := []string{"1. Lihat Profil", "2. Buat Buku", "3. List Buku", "4. Buku yang dipinjam", "5. Logout"}
	fmt.Println()
	fmt.Println("SELAMAT DATANG", strings.ToUpper(user.Nama))
	fmt.Println("===== Menu Apps =====")
	for _, v := range Mymenu {
		fmt.Println(v)
	}
	var nomor int
	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%d", &nomor)
	fmt.Println()
	switch nomor {
	case 1:
		fmt.Println("===== My Profil =====")
		// Login()
	case 2:
		// BuatBuku(user)
	case 3:
		// ListBuku(user)
	case 4:
		fmt.Println("===== Buku yang Dipinjam =====")
	case 5:
		TampilanMenuUtama()
	default:
		MenuApps(user)
	}
}

func main() {
	// DBcon := config.ConnectDB()
	// DBcon.AutoMigrate(entities.Buku{}, entities.Pinjam{})
	TampilanMenuUtama()
}
