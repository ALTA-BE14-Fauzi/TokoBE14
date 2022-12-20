package pegawai

import (
	"TokoBE14/config"
	"TokoBE14/items"
	"fmt"
	"strings"
)

func MenuPegawai(nama string) {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	itemMenu := items.ItemMenu{DB: conn}

	inputLogin := "A"
	for inputLogin != "0" {
		fmt.Println("================== PEGAWAI ====================") // MENU PEGAWAI------------
		fmt.Println("Selamat Datang ", strings.ToUpper(nama))
		fmt.Println("1. Tambahkan Barang Baru")
		fmt.Println("2. Edit Nama Barang & Stok")
		fmt.Println("3. Tambahkan Customer baru")
		fmt.Println("4. Buat Transaksi")
		fmt.Println("5. Lihat Transaksi")
		fmt.Println("0. Logout")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&inputLogin)
		// ==============================================================================================
		if inputLogin == "1" {
			fmt.Println("=========== Tambahkan Produk ===========")
			var newItem items.Items
			fmt.Print("Masukan Nama Barang : ")
			fmt.Scanln(&newItem.Nama)
			fmt.Print("Masukan Stock Barang : ")
			fmt.Scanln(&newItem.Stock)
			res, err := itemMenu.TambahItem(newItem)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("*** Sukses Menambahkan Barang ***")
			} else {
				fmt.Println("Gagal Menambahkan Barang")
			}
			// ==============================================================================================
		} else if inputLogin == "2" {
			inputEdit := "A"
			for inputEdit != "0" {
				fmt.Println("================= Edit =================")
				itemMenu.TampilkanItem()
				fmt.Println("1. Edit Nama Barang")
				fmt.Println("2. Edit Stok Barang")
				fmt.Println("0. Kembali ke menu sebelumnya")
				fmt.Print("Pilih Menu (0-2) : ")
				fmt.Scanln(&inputEdit)
				if inputEdit == "1" {
					var namaLama, namaBaru string
					fmt.Print("Masukan Nama Barang yang akan Di Edit : ")
					fmt.Scanln(&namaLama)
					fmt.Print("Ubah Nama Barang ", namaLama, " Menjadi : ")
					fmt.Scanln(&namaBaru)
					res, err := itemMenu.UbahNamaItem(namaLama, namaBaru)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res {
						fmt.Println("*** Sukses Mengganti Nama Barang ***")
					} else {
						fmt.Println("Gagal Mengganti Nama Barang")
					}

				} else if inputEdit == "2" {
					var editStock items.Items
					fmt.Print("Masukan Nama Barang : ")
					fmt.Scanln(&editStock.Nama)
					fmt.Print("Tambahkan Jumlah Stock Barang : ")
					fmt.Scanln(&editStock.Stock)
					res, err := itemMenu.UpdateStock(editStock)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res {
						fmt.Println("*** Sukses Menambah Stock Nama Barang ***")
					} else {
						fmt.Println("Gagal Menambah Stock Nama Barang")
					}
				}
			}
			// ==============================================================================================
		} else if inputLogin == "3" {
			fmt.Println("========== Tambahkan Customers =========")
			// ==============================================================================================
		} else if inputLogin == "4" {
			fmt.Println("============ Buat Transaksi ============")
			// ==============================================================================================
		} else if inputLogin == "5" {
			fmt.Println("============ Lihat Transaksi ============")

		}
		if inputLogin > "5" && inputLogin != "0" {
			fmt.Println("*** Input yang anda masukan tidak cocok.***")
		}
	}

}
