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
		fmt.Println("0. Logout")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&inputLogin)
		fmt.Println(inputLogin)
	}
}
