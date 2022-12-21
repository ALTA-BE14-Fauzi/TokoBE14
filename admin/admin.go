package admin

import (
	"TokoBE14/config"
	"TokoBE14/transaksi"
	"TokoBE14/user"
	"fmt"
	"strings"
)

func MenuAdmin(nama string) {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	authmenu := user.AuthMenu{DB: conn}
	transMenu := transaksi.TransMenu{DB: conn}

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
			// reader := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan nama : ")
			fmt.Scanln(&newUser.Nama)
			// text, _ := reader.ReadString('\n')
			// text = strings.TrimSuffix(text, "\n")
			// newUser.Nama = text

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
				fmt.Println("0. Batal")
				fmt.Print("Masukan Pilihan : ")
				fmt.Scanln(&inputHapus)
				fmt.Println(inputHapus)
				if inputHapus == "1" {
					fmt.Println("============== Hapus Pegawai ==============")
					var HapusPegawai int
					fmt.Println("List Pegawai Toko :")
					authmenu.TampilPegawai()
					fmt.Println("0. Batal")
					fmt.Println("Pilih Pegawai yang akan dihapus (0-9) : ")
					fmt.Scanln(&HapusPegawai)

					if HapusPegawai != 0 {
						res, err := authmenu.HapusPegawai(HapusPegawai)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus pegawai")
						} else {
							fmt.Println("Gagal menghapus pegawai")
						}
					}
				}
				if inputHapus == "2" {
					fmt.Println("============== Hapus Items ==============")
					var hapusItem user.Items
					fmt.Println("List Items :")
					authmenu.TampilItem()
					fmt.Println("0. Batal")
					fmt.Print("Pilih ID Item yang akan dihapus (0-9) : ")
					fmt.Scanln(&hapusItem.ID)

					if hapusItem.ID != 0 {
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

				}
				if inputHapus == "3" {
					fmt.Println("============== Hapus Transaksi ==============")
					var hapusTransaksi int
					fmt.Println("List Transaksi :")
					transMenu.TampilHapusTransaksi()
					fmt.Println("0. Batal")
					fmt.Print("Pilih Transaksi yang akan dihapus (0-9) : ")
					fmt.Scanln(&hapusTransaksi)

					if hapusTransaksi != 0 {
						res, err := transMenu.HapusTransaksi(hapusTransaksi)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus transaksi")
						} else {
							fmt.Println("Gagal menghapus transaksi")
						}
					}
				}
				if inputHapus == "4" {
					fmt.Println("============== Hapus Customers ==============")
					var hapusCust int
					fmt.Println("List Customers :")
					transMenu.TampilCustomer()
					fmt.Println("0. Batal")

					fmt.Print("Pilih Customers yang akan dihapus (0-9) : ")
					fmt.Scanln(&hapusCust)

					if hapusCust != 0 {
						res, err := transMenu.HapusCustomer(hapusCust)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus transaksi")
						} else {
							fmt.Println("Gagal menghapus transaksi")
						}
					}
				}

			}

		} else if inputLogin != "0" && inputLogin != "1" && inputLogin != "2" {
			fmt.Println("Input yang anda masukan tidak cocok, Silahkan sesuai pilihan menu")
		}
	}
	// fmt.Println(inputLogin)
}
