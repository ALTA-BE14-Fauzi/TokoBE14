package pegawai

import (
	"fmt"
	"strings"
)

func MenuPegawai(nama string) {
	inputLogin := "A"
	for inputLogin != "0" {
		fmt.Println("================== PEGAWAI ====================") // MENU PEGAWAI------------
		fmt.Println("Selamat Datang ", strings.ToUpper(nama))
		fmt.Println("1. Tambahkan Barang Baru")
		fmt.Println("2. Update")
		fmt.Println("3. Tambahkan Customer baru")
		fmt.Println("4. Buat Transaksi")
		fmt.Println("5. Lihat Transaksi")
		fmt.Println("0. Logout")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&inputLogin)
		if inputLogin == "1" {
			fmt.Println("=========== Tambahkan Produk ===========")
		} else if inputLogin == "2" {
			fmt.Println("================= Edit =================")
		} else if inputLogin == "3" {
			fmt.Println("========== Tambahkan Customers =========")
		} else if inputLogin == "4" {
			fmt.Println("============ Buat Transaksi ============")
		} else if inputLogin == "5" {
			fmt.Println("=========== Lihat Transaksi ============")

		}
		if inputLogin > "5" && inputLogin != "0" {
			fmt.Println("Input yang anda masukan tidak cocok.")
		}
	}

}
