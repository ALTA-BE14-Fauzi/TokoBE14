package main

import (
	"TokoBE14/admin"
	"TokoBE14/config"
	"TokoBE14/pegawai"
	"TokoBE14/user"
	"fmt"
	"reflect"
)

func main() {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	authmenu := user.AuthMenu{DB: conn}
	var inputHome string
	fmt.Println(reflect.ValueOf(authmenu).Kind())
	for inputHome != "0" {
		// MENU AWAL
		fmt.Println("============= Welcome to TOKO BE 14 =============")
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Print("Masukan Pilihan : ")
		fmt.Scanln(&inputHome)
		// fmt.Println(string(inputHome))
		if inputHome == "1" {
			var inputNama, inputPassword string
			fmt.Println("============= Log In =============")
			fmt.Print("Masukkan Nama : ")
			fmt.Scanln(&inputNama)
			fmt.Print("Masukkan Password : ")
			fmt.Scanln(&inputPassword)
			res, err := authmenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}
			// fmt.Println(res.Role)
			if res.Role == 1 {
				admin.MenuAdmin(res.Nama)
			}
			if res.Role == 2 {
				pegawai.MenuPegawai(res.Nama)
			}
		} else if inputHome != "0" {
			fmt.Println("Anda Hanya dapat Menginput 1. Login dan 0. Exit")
		}
	}
}
