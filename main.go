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

	Mymenu := []string{"1. Login", "2. Register", "3. List Buku", "4. Keluar"}
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

		ListBuku(entities.Users{})
	case 4:
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

//Tampilkan List Buku
func ListBuku(user entities.Users) {
	fmt.Println("===== List Buku =====")
	conn := config.ConnectDB()
	buku := datastore.BukuDB{Db: conn}
	list, err := buku.ListBuku()
	if err != nil {
		fmt.Println(err)
		MenuApps(user)
	}
	for i, v := range list {
		fmt.Println(i+1, " ", strings.ToUpper(v.NameBuku))
		fmt.Println("    Author : ", "(", v.Author, ")", " Dipinjam : ", "(", v.Jumlah, ")", " kali")
	}
	fmt.Println("99 Kembali Ke Menu Sebelumnya")
	fmt.Print("Masukkan Pilihan: ")
	var angka int
	fmt.Scanf("%d", &angka)
	for i, v := range list {
		if i+1 == angka {
			DetailBuku(v.NameBuku, v.UserID, user)
		} else if angka == 99 && user.ID != 0 {
			MenuApps(user)
		} else if angka == 99 && user.ID == 0 {
			TampilanMenuUtama()
		} else if angka > len(list) || angka == 0 {
			ListBuku(user)
		}
	}
}

//Detail Buku
func DetailBuku(Buku string, UserID uint, user entities.Users) {
	fmt.Println("===== Detail Buku =====")
	fmt.Println(user.ID == 0)
	conn := config.ConnectDB()
	buku := datastore.BukuDB{Db: conn}
	detail, err := buku.DetailBuku(Buku, UserID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Nama Buku 		: ", strings.ToUpper(detail.NameBuku))
	fmt.Println("Nama Penulis 		: ", strings.ToUpper(detail.Nama))
	fmt.Println("Di Pinjam Sebanyak 	: ", detail.Jumlah, " Kali")
	if detail.Status == true {
		fmt.Println("Status 			:  Bisa Dipinjam")
	} else {
		fmt.Println("Status 			:  Sedang Dipinjam")
	}
	fmt.Println("Pilihan : ")
	if user.ID == detail.UserID {
		fmt.Println("1. Edit Buku")
		fmt.Println("2. Delete Buku")
		fmt.Println("3. Pinjam Buku")
		fmt.Println("99. Menu Sebelumnya")
	} else if user.ID == 0 {
		fmt.Println("99. Menu Sebelumnya")
	} else {
		fmt.Println("1. Pinjam Buku")
		fmt.Println("99. Menu Sebelumnya")
	}
	pilihan := 0
	fmt.Print("Pilih Menu : ")
	fmt.Scanf("%d", &pilihan)
	if user.ID == detail.UserID {
		switch pilihan {
		case 1:
			fmt.Println("Buku diedit")
			UpdateBuku(Buku, UserID, user)
		case 2:
			// DeleteBuku
		case 3:
			fmt.Println("Buku di Pinjam")
			DetailBuku(Buku, UserID, user)
		case 99:
			ListBuku(user)
		default:
			DetailBuku(Buku, UserID, user)
		}
	} else if user.ID == 0 {
		switch pilihan {
		case 99:
			ListBuku(user)
		default:
			DetailBuku(Buku, UserID, user)
		}
	} else {
		switch pilihan {
		case 1:
			fmt.Println("Buku di Pinjam")
			DetailBuku(Buku, UserID, user)
		case 99:
			ListBuku(user)
		default:
			DetailBuku(Buku, UserID, user)
		}
	}
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

		BuatBuku(user)
	case 3:
		ListBuku(user)
	case 4:
		fmt.Println("===== Buku yang Dipinjam =====")
	case 5:
		TampilanMenuUtama()
	default:
		MenuApps(user)
	}
}

//BUAT BUKU
func BuatBuku(user entities.Users) {
	fmt.Println("===== Membuat Buku =====")
	conn := config.ConnectDB()
	buku := datastore.BukuDB{Db: conn}
	var NamaBuku string
	fmt.Print("Masukkan Nama Buku: ")
	fmt.Scanf("%s", &NamaBuku)
	TambahBuku, err := buku.TambahBuku(entities.Buku{NameBuku: NamaBuku}, user)
	if err != nil {
		fmt.Println(err)
		MenuApps(user)
	} else {
		fmt.Println()
		fmt.Println("BUKU ", strings.ToUpper(TambahBuku.NameBuku), " BERHASIL DI BUAT")
		MenuApps(user)
	}
}

func UpdateBuku(Buku string, UserID uint, user entities.Users) {
	fmt.Println("===== Update Buku =====")
	conn := config.ConnectDB()
	update := datastore.BukuDB{Db: conn}
	fmt.Println("Nama Buku Sebelumnya : ", Buku)
	fmt.Println("Ketik (batal) Untuk Membatalkan Perubahan")
	fmt.Print("Nama Buku Yang Baru : ")
	var namabaru string
	fmt.Scanf("%s", &namabaru)
	if namabaru == "batal" {
		DetailBuku(Buku, UserID, user)
	} else {
		_, err := update.UpdateBuku(Buku, namabaru)
		if err != nil {
			fmt.Println()
			fmt.Println("UPDATE NAMA BUKU GAGAL")
			fmt.Println()
			UpdateBuku(Buku, UserID, user)
		} else {
			fmt.Println()
			fmt.Println("UPDATE NAMA BUKU BERHASIL")
			fmt.Println()
			UpdateBuku(Buku, UserID, user)
		}
	}
}

func main() {
	// DBcon := config.ConnectDB()
	// DBcon.AutoMigrate(entities.Users{}, entities.Buku{}, entities.Pinjam{})
	TampilanMenuUtama()
}
