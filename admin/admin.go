package admin

import (
	"TokoBE14/config"
	"TokoBE14/user"
	"fmt"
	"strings"
)

func MenuAdmin(nama string) {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	authmenu := user.AuthMenu{DB: conn}

	inputLogin := "A"
	for inputLogin != "0" {
		fmt.Println("=================== ADMIN =====================") // MENU ADMIN-------------
		fmt.Println("Selamat Datang ", strings.ToUpper(nama))
		fmt.Println("1. Tambahkan Pegawai Baru")
		fmt.Println("2. Hapus Data")
		fmt.Println("0. Logout")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&inputLogin)
		if inputLogin == "1" {
			fmt.Println("========== Tambahkan Pegawai ==========")
			var newUser user.User
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan password : ")
			fmt.Scanln(&newUser.Password)
			res, err := authmenu.Register(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("Sukses mendaftarkan Pegawai")
			} else {
				fmt.Println("Gagal mendaftarkann Pegawai")
			}
		} else if inputLogin == "2" {
			inputHapus := "A"
			for inputHapus != "0" {

				fmt.Println("========== Hapus ==========")
				fmt.Println("1. Hapus Pegawai")
				fmt.Println("2. Hapus Items")
				fmt.Println("3. Hapus Transaksi")
				fmt.Println("4. Hapus Customers")
				fmt.Print("Masukan Pilihan : ")
				fmt.Scanln(&inputHapus)
				fmt.Println(inputHapus)
				if inputHapus == "1" {
					fmt.Println("============== Hapus Pegawai ==============")
					fmt.Println("List Pegawai Toko :")
					fmt.Println("0. Batal")
					fmt.Println("1. Syahrini")
					fmt.Println("Pilih Pegawai yang akan dihapus (0-9) : ")

				}
				if inputHapus == "2" {
					fmt.Println("============== Hapus Items ==============")
					var hapusItem user.Items
					fmt.Println("List Items :")
					authmenu.Tampilkan("messi", "messi123")
					fmt.Println("0. Batal")
					fmt.Println("1. Syahrini")
					fmt.Print("Pilih Item yang akan dihapus (0-9) : ")
					fmt.Scanln(&hapusItem.ID)
					res, err := authmenu.HapusItem(hapusItem)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res {
						fmt.Println("Sukses menghapus item")
					} else {
						fmt.Println("Gagal menghapus item")
					}
				}
				if inputHapus == "3" {
					fmt.Println("============== Hapus Transaksi ==============")
					fmt.Println("List Transaksi :")
					fmt.Println("0. Batal")
					fmt.Println("1. Indomie")
					fmt.Println("Pilih Pegawai yang akan dihapus (0-9) : ")
				}
				if inputHapus == "4" {
					fmt.Println("============== Hapus Customers ==============")
					fmt.Println("List Customers :")
					fmt.Println("0. Batal")
					fmt.Println("1. Syahrini")
					fmt.Println("Pilih Customers yang akan dihapus (0-9) : ")
				}

			}

		} else if inputLogin != "0" && inputLogin != "1" && inputLogin != "2" {
			fmt.Println("Input yang anda masukan tidak cocok, Silahkan sesuai pilihan menu")
		}
	}
	// fmt.Println(inputLogin)
}
