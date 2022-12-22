package main

import (
	"TokoBE14/admin"
	"TokoBE14/config"
	"TokoBE14/pegawai"
	"TokoBE14/user"
	"fmt"
)

type User struct {
	ID       int
	Nama     string
	Password string
	Role     int
}

func main() {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	authmenu := user.AuthMenu{DB: conn}

	var inputHome string
	for inputHome != "0" {
		// MENU AWAL
		fmt.Println("============= WELCOME TO TOKO BE 14 =============")
		fmt.Println("")
		fmt.Println("\t\t   Select Menu")
		fmt.Print("\t\t1. Login   ")
		fmt.Println("0. Exit")
		fmt.Println("")
		fmt.Print("Enter your menu: ")
		fmt.Scanln(&inputHome)
		fmt.Println("")
		// fmt.Println(string(inputHome))
		if inputHome == "1" {
			var inputNama, inputPassword string
			fmt.Println("-------------------| Log  In |-------------------")
			fmt.Print("Enter username : ")
			fmt.Scanln(&inputNama)
			fmt.Print("Enter Password : ")
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
			fmt.Println("Please input 1 for Login and 0 to Exit")
		}
	}
}
