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
		fmt.Println("================= Admin  Menu ===================") // MENU ADMIN-------------
		fmt.Println("Welcome ", strings.ToUpper(nama))
		fmt.Println("1. Add New Employee")
		fmt.Println("2. Delete Data")
		fmt.Println("0. Logout")
		fmt.Print("Enter your option : ")
		fmt.Scanln(&inputLogin)
		fmt.Println("")
		if inputLogin == "1" {
			fmt.Println("========== Add New Employee ==========")
			var newUser user.User
			fmt.Print("Enter employee's name : ")
			fmt.Scanln(&newUser.Nama)
			// reader := bufio.NewReader(os.Stdin)
			// text, _ := reader.ReadString('\n')
			// text = strings.TrimSuffix(text, "\n")
			// newUser.Nama = text

			fmt.Print("Enter new employee's password : ")
			fmt.Scanln(&newUser.Password)
			namaBaru := newUser.Nama
			passBaru := newUser.Password
			if namaBaru != "" && passBaru != "" {
				res, err := authmenu.Register(newUser)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("Successfully added new employee")
				} else {
					fmt.Println("Failed to add new employee")
				}
			} else {
				fmt.Println("Input can't be empty, Failed to add new employee ")
			}
		} else if inputLogin == "2" {
			inputHapus := "A"
			for inputHapus != "0" {

				fmt.Println("========== Delete Menu ==========")
				fmt.Println("1. Delete Employee")
				fmt.Println("2. Delete Item")
				fmt.Println("3. Delete Transaction")
				fmt.Println("4. Delete Customer")
				fmt.Println("0. Cancel")
				fmt.Print("Enter your option : ")
				fmt.Scanln(&inputHapus)
				fmt.Println(inputHapus)
				if inputHapus == "1" {
					fmt.Println("============== Delete Employee ==============")
					var HapusPegawai int
					fmt.Println("List Store Employee :")
					authmenu.TampilPegawai()
					fmt.Println("0. Cancel")
					fmt.Println("Choose an employee you'd like to delete: ")
					fmt.Scanln(&HapusPegawai)

					if HapusPegawai != 0 {
						res, err := authmenu.HapusPegawai(HapusPegawai)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Successfully deleted an employee")
						} else {
							fmt.Println("Failed to delete an employee")
						}
					}
				}
				if inputHapus == "2" {
					fmt.Println("============== Delete Item ==============")
					var hapusItem user.Items
					fmt.Println("Item List :")
					authmenu.TampilItem()
					fmt.Println("0. Cancel")
					fmt.Print("Choose Item ID you'd like to delete : ")
					fmt.Scanln(&hapusItem.ID)

					if hapusItem.ID != 0 {
						res, err := authmenu.HapusItem(hapusItem)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Successfully deleted an item")
						} else {
							fmt.Println("Failed to delete an item")
						}
					}

				}
				if inputHapus == "3" {
					fmt.Println("============== Delete Transaction ==============")
					var hapusTransaksi int
					fmt.Println("Transaction List :")
					transMenu.TampilHapusTransaksi()
					fmt.Println("0. Cancel")
					fmt.Print("Choose a transaction you'd like to remove (0-9) : ")
					fmt.Scanln(&hapusTransaksi)

					if hapusTransaksi != 0 {
						res, err := transMenu.HapusTransaksi(hapusTransaksi)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Successfully deleted a transaction")
						} else {
							fmt.Println("Failed to delete a transaction")
						}
					}
				}
				if inputHapus == "4" {
					fmt.Println("============== Delete Customer ==============")
					var hapusCust int
					fmt.Println("Customers list :")
					transMenu.TampilCustomer()
					fmt.Println("0. Cancel")

					fmt.Print("Choose a customer you'd like to delete : ")
					fmt.Scanln(&hapusCust)

					if hapusCust != 0 {
						res, err := transMenu.HapusCustomer(hapusCust)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Successfully deleted a customer")
						} else {
							fmt.Println("Failed to delete a customer")
						}
					}
				}

			}

		} else if inputLogin != "0" && inputLogin != "1" && inputLogin != "2" {
			fmt.Println("Invalid input, please input according to the menu option")
		}
	}
	// fmt.Println(inputLogin)
}
