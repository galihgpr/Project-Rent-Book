package main

import (
	"fmt"
	"intro-golang/entities"
)

// //Menu Login
// func Login() {

// }

// //Menu Register
// func Register() {
// 	// koneksi ke database
// 	conn := config.ConnectDB()
// 	// migration table
// 	conn.AutoMigrate(entities.Users{}, entities.Buku{}, entities.Pinjam{})
// 	fmt.Println("migrasi berhasil")

// 	user := user.UserDB{Db: conn}
// 	var nama, email, password, hp string
// 	fmt.Print("Masukkan Nama : ")
// 	fmt.Scanf("%s", &nama)
// 	fmt.Print("Masukkan Email : ")
// 	fmt.Scanf("%s", &email)
// 	fmt.Print("Masukkan Password : ")
// 	fmt.Scanf("%s", &password)
// 	fmt.Print("Masukkan No HP : ")
// 	fmt.Scanf("%s", &hp)
// 	daftar, err := user.Register(entities.Users{Nama: nama, Email: email, Password: password, Hp: hp})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(daftar)
// }

// //Tampilkan List Buku
// func ListBuku() {

// }

// //Menu Apps

// func MenuApps() {
// 	Mymenu := []string{"1. Lihat Profil", "2. Buat Buku", "3. List Buku", "4. Buku yang dipinjam"}
// 	fmt.Println("===== Menu Apps =====")
// 	for _, v := range Mymenu {
// 		fmt.Println(v)
// 	}
// 	var nomor int
// 	fmt.Print("Pilih Menu : ")
// 	fmt.Scanf("%d", &nomor)
// 	fmt.Println("\n")
// 	switch nomor {
// 	case 1:
// 		fmt.Println("===== My Profil =====")
// 		// Login()
// 	case 2:
// 		fmt.Println("===== Tambahkan Buku =====")
// 		// Register()
// 	case 3:
// 		fmt.Println("===== List Buku =====")
// 		// ListBuku()
// 	default:
// 		fmt.Println("===== Buku yang Dipinjam =====")
// 	}
// }

func main() {
	DBcon := config.connectDB()
	DBcon.AutoMigrate(entities.Users{}, entities.Buku{}, entities.Pinjam{})
	fmt.Println("migrasi berhasil")

	// MENU UTAMA
	Mymenu := []string{"1. Login", "2. Register", "3. Lis Buku", "4. Keluar"}
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
		fmt.Println("===== Login Akun =====")
		Login()
	case 2:
		fmt.Println("===== Registrasi Akun =====")
		Register()
	case 3:
		fmt.Println("===== List Buku =====")
		ListBuku()
	default:
		fmt.Println("===== Terimakasih Sudah Berkunjung =====")
	}
}
