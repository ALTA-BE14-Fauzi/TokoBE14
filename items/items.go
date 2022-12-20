package items

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Items struct {
	ID    int
	Nama  string
	Stock int
}

type ItemMenu struct {
	DB *sql.DB
}

func (im *ItemMenu) DuplicateItem(iNama string) bool {
	res := im.DB.QueryRow("SELECT id FROM items where nama = ?", iNama)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Product Baru Berhasil Dibuat", err.Error())
		return false
	}
	return true
}

func (im *ItemMenu) TambahItem(newItem Items) (bool, error) {
	itemQuery, err := im.DB.Prepare("INSERT INTO items(nama,stock) VALUES (?,?)")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare INSERT ke tabel items ERROR **")
	}
	if im.DuplicateItem(newItem.Nama) {
		log.Println("--- Duplicated information ---")
		return false, errors.New("nama sudah digunakan")
	}
	res, err := itemQuery.Exec(newItem.Nama, newItem.Stock)
	if err != nil {
		log.Println("Insert Items ", err.Error())
		return false, errors.New("** Error saat Insert Item **")
	}
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("Error setelah insert item ", err.Error())
		return false, errors.New("error setelah insert")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (im *ItemMenu) TampilkanItem() {
	resultRows, err := im.DB.Query("SELECT * FROM items ")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrItem := []Items{}
	for resultRows.Next() {
		tmp := Items{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Stock)
		arrItem = append(arrItem, tmp)
	}
	// id := arrItem[0].Nama
	// namar := arrItem[0].Password
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|  No  |\t Nama\t\t|\tStock   |")
	fmt.Println("|-----------------------------------------------|")
	for i := 0; i < len(arrItem); i++ {
		if len(arrItem[i].Nama) > 5 {
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t|\t", arrItem[i].Stock, "    |")
		} else {
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t\t|\t", arrItem[i].Stock, "    |")

		}
	}
	fmt.Println("|-----------------------------------------------|")
}
