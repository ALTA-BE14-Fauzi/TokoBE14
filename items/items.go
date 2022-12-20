package items

import (
	"database/sql"
	"errors"
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
