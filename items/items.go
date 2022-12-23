package items

import (
	"TokoBE14/user"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Items struct {
	ID      int
	Nama    string
	Stock   int
	user_id int
}

type ItemMenu struct {
	DB *sql.DB
}

//=========================================================================================Duplicate

func (im *ItemMenu) DuplicateItem(iNama string) bool {
	res := im.DB.QueryRow("SELECT id FROM items where nama = ?", iNama)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("New item has been made", err.Error())
		return false
	}
	return true
}

//======================================================================================TAMBAH BARANG

func (im *ItemMenu) TambahItem(nama string, newItem Items) (bool, error) {

	resultRows, err := im.DB.Query("SELECT id FROM users WHERE nama = ?", nama)
	if err != nil {
		fmt.Println("Error Reading Data from Database", err.Error())
	}
	arrUser := []user.User{}
	for resultRows.Next() {
		tmp := user.User{}
		resultRows.Scan(&tmp.ID)
		arrUser = append(arrUser, tmp)
	}
	userID := arrUser[0].ID

	itemQuery, err := im.DB.Prepare("INSERT INTO items(nama,stock,user_id) VALUES (?,?,?)")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare INSERT to items table ERROR **")
	}
	if im.DuplicateItem(newItem.Nama) {
		log.Println("--- Duplicated information ---")
		return false, errors.New("name already exists")
	}
	res, err := itemQuery.Exec(newItem.Nama, newItem.Stock, userID)
	if err != nil {
		log.Println("Insert Items ", err.Error())
		return false, errors.New("** Error when inserting Item **")
	}
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("Error after inserting item ", err.Error())
		return false, errors.New("error after insert")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

//====================================================================================== UBAH NAMA BARANG

func (im *ItemMenu) UbahNamaItem(namaLama string, namaBaru string) (bool, error) {
	itemQuery, err := im.DB.Prepare("UPDATE items SET nama = ? WHERE nama = ?")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare Update to item's table ERROR **")
	}
	res, err := itemQuery.Exec(namaBaru, namaLama)
	if err != nil {
		log.Println("Update Items ", err.Error())
		return false, errors.New("** Error while Updating Item **")
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error after update item ", err.Error())
		return false, errors.New("error after update")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

//====================================================================================== TAMBAH STOK

func (im *ItemMenu) UpdateStock(editStock Items) (bool, error) {
	// resultRows, err := im.DB.Query("SELECT stock FROM items WHERE nama=? ", editStock.Nama)
	// if err != nil {
	// 	fmt.Println("Ambil Data dari Database Error", err.Error())
	// }
	// arrItem := []Items{}
	// for resultRows.Next() {
	// 	tmp := Items{}
	// 	resultRows.Scan(&tmp.Stock)
	// 	arrItem = append(arrItem, tmp)
	// }
	// fmt.Println(arrItem[0].Stock)
	// newStock := (arrItem[0].Stock) + editStock.Stock

	itemQuery, err := im.DB.Prepare("UPDATE items SET stock = ? WHERE nama = ?")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare Update Stock to items table ERROR **")
	}
	res, err := itemQuery.Exec(editStock.Stock, editStock.Nama)
	if err != nil {
		log.Println("Update Stock Items ", err.Error())
		return false, errors.New("** Error while Updating Stock Item **")
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error after updating Stock item ", err.Error())
		return false, errors.New("error after updating Stock")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

//========================================================================================NEW CUSTOMER

func (im *ItemMenu) DuplicateCustomer(cNama string) bool {
	res := im.DB.QueryRow("SELECT id FROM customers where nama = ?", cNama)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("New customer has been successfully registered", err.Error())
		return false
	}
	return true
}

func (im *ItemMenu) RegisterCustomer(nama string) (bool, error) {
	itemQuery, err := im.DB.Prepare("INSERT INTO customers(nama) VALUES (?)")
	if err != nil {
		log.Println("prepare insert customer ", err.Error())
		return false, errors.New("** Prepare INSERT to customer table ERROR **")
	}
	if im.DuplicateCustomer(nama) {
		log.Println("--- Duplicated information ---")
		return false, errors.New("name already exists")
	}
	res, err := itemQuery.Exec(nama)
	if err != nil {
		log.Println("Insert customer ", err.Error())
		return false, errors.New("** Error while Inserting customer **")
	}
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("Error after inserting item ", err.Error())
		return false, errors.New("error after inserting")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

//=======================================================================================TAMPILKAN SEMUA BARANG

func (im *ItemMenu) TampilkanItem() {
	resultRows, err := im.DB.Query("SELECT * FROM items ")
	if err != nil {
		fmt.Println("Read Data from Database Error", err.Error())
	}
	arrItem := []Items{}
	for resultRows.Next() {
		tmp := Items{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Stock, &tmp.user_id)
		arrItem = append(arrItem, tmp)
	}
	// id := arrItem[0].Nama
	// namar := arrItem[0].Password
	fmt.Println("|------------------------------------------|")
	fmt.Println("|  No  |\t Name\t\t|  Stock   |")
	fmt.Println("|------------------------------------------|")
	for i := 0; i < len(arrItem); i++ {
		if len(arrItem[i].Nama) > 5 {
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t| ", arrItem[i].Stock, "\t   |")
		} else {
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t\t| ", arrItem[i].Stock, "\t   |")

		}
	}
	fmt.Println("|------------------------------------------|")
}

type ItemModif struct {
	ID       int
	Nama     string
	Stock    int
	UserNama string
}

func (im *ItemMenu) TampilkanItemFull() {
	resultRows, err := im.DB.Query("SELECT i.id , i.nama,i.stock, u.nama FROM items i LEFT JOIN users u ON u.id = i.user_id;")
	if err != nil {
		fmt.Println("Read Data from Database Error", err.Error())
	}
	arrItem := []ItemModif{}
	for resultRows.Next() {
		tmp := ItemModif{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Stock, &tmp.UserNama)
		arrItem = append(arrItem, tmp)
	}

	fmt.Println("|---------------------------------------------------------------|")
	fmt.Println("|  No\t|\t Name\t\t| Stock\t | Data Entered By\t|")
	fmt.Println("|---------------------------------------------------------------|")
	for i := 0; i < len(arrItem); i++ {
		if len(arrItem[i].Nama) > 5 {
			fmt.Println("| ", i+1, "\t|\t", arrItem[i].Nama, "\t| ", arrItem[i].Stock, "\t | ", arrItem[i].UserNama, "\t\t|")
		} else {
			fmt.Println("| ", i+1, "\t|\t", arrItem[i].Nama, "\t\t| ", arrItem[i].Stock, "\t | ", arrItem[i].UserNama, "\t\t|")

		}
	}
	fmt.Println("|---------------------------------------------------------------|")
}
