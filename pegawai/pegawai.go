package pegawai

import (
	"TokoBE14/config"
	"TokoBE14/items"
	"TokoBE14/transaksi"
	"fmt"
	"strings"
)

func MenuPegawai(nama string) {
	cfg := config.ReadConfig()
	conn := config.ConnectSQL(*cfg)
	itemMenu := items.ItemMenu{DB: conn}
	transMenu := transaksi.TransMenu{DB: conn}

	inputLogin := "A"
	for inputLogin != "0" {
		fmt.Println("===============| Employee  Menu |================") // MENU PEGAWAI------------
		fmt.Println("Welcome", strings.ToUpper(nama))
		fmt.Println("  1. Add New Item")
		fmt.Println("  2. Edit Item Name & Stock")
		fmt.Println("  3. Add New Customer")
		fmt.Println("  4. Add New Transaction")
		fmt.Println("  5. Show Transactions")
		fmt.Println("  6. Show Item Table")
		fmt.Println("  0. Logout")
		fmt.Print("Enter your option: ")
		fmt.Scanln(&inputLogin)
		fmt.Println("")
		// ==============================================================================================
		if inputLogin == "1" {
			fmt.Println("=========== Add New Item ===========")
			var newItem items.Items
			fmt.Print("Enter item name : ")
			fmt.Scanln(&newItem.Nama)
			fmt.Print("Enter item stock : ")
			fmt.Scanln(&newItem.Stock)
			itemBaru := newItem.Nama
			if itemBaru != "" {
				res, err := itemMenu.TambahItem(nama, newItem)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("*** Product successfully added ***")
				} else {
					fmt.Println("Failed to add product")
				}
			} else {
				fmt.Println("Item Name is not allowed empty, Failed to add product")
			}
			// ==============================================================================================
		} else if inputLogin == "2" {
			inputEdit := "A"
			for inputEdit != "0" {
				fmt.Println("=================== Edit Item ===================")
				itemMenu.TampilkanItem()
				fmt.Println("1. Edit Item's Name")
				fmt.Println("2. Update Item's Stock")
				fmt.Println("0. Back to previous menu")
				fmt.Print("Input menu (0-2) : ")
				fmt.Scanln(&inputEdit)
				fmt.Println("")
				if inputEdit == "1" {
					var namaLama, namaBaru string
					fmt.Print("Input item name that you want to change : ")
					fmt.Scanln(&namaLama)
					fmt.Print("Change old name ", namaLama, " to : ")
					fmt.Scanln(&namaBaru)
					res, err := itemMenu.UbahNamaItem(namaLama, namaBaru)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res {
						fmt.Println("*** Successfully changed name ***")
					} else {
						fmt.Println("Failed to change name")
					}

				} else if inputEdit == "2" {
					var editStock items.Items
					fmt.Print("Input item name : ")
					fmt.Scanln(&editStock.Nama)
					fmt.Print("Update item stock : ")
					fmt.Scanln(&editStock.Stock)
					res, err := itemMenu.UpdateStock(editStock)
					if err != nil {
						fmt.Println(err.Error())
					}
					if res {
						fmt.Println("*** Successfully updated item's stock ***")
					} else {
						fmt.Println("Failed to update item's stock")
					}
				}
			}
			// ==============================================================================================
		} else if inputLogin == "3" {
			fmt.Println("================= Add Customers =================")
			var namaCus string
			fmt.Print("Enter customer name : ")
			fmt.Scanln(&namaCus)
			if namaCus != "" {
				res, err := itemMenu.RegisterCustomer(namaCus)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("*** Successfully added customer ***")
				} else {
					fmt.Println("Failed to add customer")
				}
			} else {
				fmt.Println("Customer Name is not allowed empty, Failed to add customer")
			}
			// ==============================================================================================
		} else if inputLogin == "4" {
			fmt.Println("============== Create Transaction ==============")
			itemMenu.TampilkanItem()
			var namaBarang, namaPembeli string
			fmt.Print("Enter customer's name : ")
			fmt.Scanln(&namaPembeli)
			if namaPembeli != "" && namaBarang != "0" {
				res, err := transMenu.BuatTransaksi(nama, namaPembeli)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("Transaction with customer name ", namaPembeli, " success create ")
					for namaBarang != "0" {
						fmt.Print("Input item Name (0 to cancel) : ")
						fmt.Scan(&namaBarang)
						if namaBarang != "0" {
							res, err := transMenu.BuatTransaksiItems(namaBarang)
							if err != nil {
								fmt.Println(err.Error())
							}
							if res {
								fmt.Println(" OK ✓")
							} else {
								fmt.Println("Failed to input item")
							}
						}

					}
				} else {
					fmt.Println("Failed to create transaction")
				}
				// ----------------CEK APAKAH ADA TRANSAKSI TAPI TIDAK ADA BARANG YANG DIBELI-----------
				if namaBarang == "0" {
					res, err := transMenu.CekTransaksiItems()
					if err != nil {
						fmt.Println(err.Error())
					}
					if !res {
						fmt.Println("** Transaksi Dibatalkan **")
						transMenu.BatalDanHapusTransaksi()
						namaBarang = "0"
					}

				}
				inputLogin = "A"
			} else {
				fmt.Println("Customer Name must be fill, Failed to create transaction")
			}

		} else if inputLogin == "5" {
			transMenu.TampilTransaksiModif()
			var inputIDTrans int
			fmt.Println("0. Cancel or Exit View")
			fmt.Print("Select and input transaction ID to see more: ")
			fmt.Scanln(&inputIDTrans)
			if inputIDTrans != 0 {
				res, err := transMenu.TranksaksiItem(inputIDTrans)
				if err != nil {
					fmt.Println(err.Error())
				}
				if res {
					fmt.Println("Here is the transaction table based on the transaction ID", inputIDTrans)
					transMenu.ViewTransaksiItem(inputIDTrans)
					var ExitView int
					fmt.Print("Press enter to exit ")
					fmt.Scanln(&ExitView)
				} else {
					fmt.Println("Failed to display transaction")
				}

			}

		} else if inputLogin == "6" {
			itemMenu.TampilkanItemFull()
			var ExitView int
			fmt.Print("Press enter to exit ")
			fmt.Scanln(&ExitView)
		}
		if inputLogin > "6" && inputLogin != "0" && inputLogin != "A" {
			fmt.Println("*** Incorrect input. Please input accordingly***")
		}
	}

}
