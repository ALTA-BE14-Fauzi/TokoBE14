package items

import (
	"TokoBE14/user"
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

//=========================================================================================Duplicate

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

//======================================================================================TAMBAH BARANG

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

//====================================================================================== UBAH NAMA BARANG

func (im *ItemMenu) UbahNamaItem(namaLama string, namaBaru string) (bool, error) {
	itemQuery, err := im.DB.Prepare("UPDATE items SET nama = ? WHERE nama = ?")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare Update ke tabel items ERROR **")
	}
	res, err := itemQuery.Exec(namaBaru, namaLama)
	if err != nil {
		log.Println("Update Items ", err.Error())
		return false, errors.New("** Error saat Update Item **")
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error setelah update item ", err.Error())
		return false, errors.New("error setelah update")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

//====================================================================================== TAMBAH STOK

func (im *ItemMenu) UpdateStock(editStock Items) (bool, error) {
	resultRows, err := im.DB.Query("SELECT stock FROM items WHERE nama=? ", editStock.Nama)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrItem := []Items{}
	for resultRows.Next() {
		tmp := Items{}
		resultRows.Scan(&tmp.Stock)
		arrItem = append(arrItem, tmp)
	}
	fmt.Println(arrItem[0].Stock)
	newStock := (arrItem[0].Stock) + editStock.Stock

	itemQuery, err := im.DB.Prepare("UPDATE items SET stock = ? WHERE nama = ?")
	if err != nil {
		log.Println("prepare insert items ", err.Error())
		return false, errors.New("** Prepare Update Stock ke tabel items ERROR **")
	}
	res, err := itemQuery.Exec(newStock, editStock.Nama)
	if err != nil {
		log.Println("Update Stock Items ", err.Error())
		return false, errors.New("** Error saat Update Stock Item **")
	}
	affRows, err := res.RowsAffected()
	if err != nil {
		log.Println("Error setelah update Stock item ", err.Error())
		return false, errors.New("error setelah update Stock")
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
		log.Println("Pelanggan/Customer Baru Berhasil Didaftarkan", err.Error())
		return false
	}
	return true
}

func (im *ItemMenu) RegisterCustomer(nama string) (bool, error) {
	itemQuery, err := im.DB.Prepare("INSERT INTO customers(nama) VALUES (?)")
	if err != nil {
		log.Println("prepare insert customer ", err.Error())
		return false, errors.New("** Prepare INSERT ke tabel customer ERROR **")
	}
	if im.DuplicateCustomer(nama) {
		log.Println("--- Duplicated information ---")
		return false, errors.New("nama sudah digunakan")
	}
	res, err := itemQuery.Exec(nama)
	if err != nil {
		log.Println("Insert customer ", err.Error())
		return false, errors.New("** Error saat Insert customer **")
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

// ===========================================================================================BUAT TRANSAKSI
func (im *ItemMenu) CekBarang(namaBarang string) bool {
	res := im.DB.QueryRow("SELECT id FROM items where nama = ?", namaBarang)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Barang Tidak Tersedia Ditoko ini", err.Error())
		return true
	}
	return false
}
func (im *ItemMenu) CekCustomer(namaCustomer string) bool {
	res := im.DB.QueryRow("SELECT id FROM customers where nama = ?", namaCustomer)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Customer Belum Terdaftar", err.Error())
		return true
	}
	return false
}

type Customer struct {
	ID   int
	Nama string
}

func (im *ItemMenu) BuatTransaksi(nama string, namaBarang string, namaCustomer string) (bool, error) {
	//-------------------------CARI ID USER---------------------
	resultRows, err := im.DB.Query("SELECT id FROM users WHERE nama =?", nama)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrUser := []user.User{}
	for resultRows.Next() {
		tmp := user.User{}
		resultRows.Scan(&tmp.ID)
		arrUser = append(arrUser, tmp)
	}
	userID := arrUser[0].ID

	//----------------------Cek Customer----------------------
	if im.CekCustomer(namaCustomer) {
		log.Println("--- Empty Data ---")
		return false, errors.New("--Customer Belum Terdaftar--")
	}

	//-----------------------Cari ID Customer--------------------
	resIDCust, err := im.DB.Query("SELECT id FROM customers WHERE nama =?", namaCustomer)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrCust := []Customer{}
	for resIDCust.Next() {
		tmpCust := Customer{}
		resIDCust.Scan(&tmpCust.ID)
		arrCust = append(arrCust, tmpCust)
	}
	custID := arrCust[0].ID
	//------------------------Cek Barang-----------------------
	if im.CekBarang(namaBarang) {
		log.Println("--- Empty Item ---")
		return false, errors.New("--Barang Tidak Tersedia--")
	}

	//---------------------Cek Stok & ID----------------------
	ResStockRows, err := im.DB.Query("SELECT id,stock FROM items WHERE nama =?", namaBarang)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrStockItem := []Items{}
	for ResStockRows.Next() {
		tmpStock := Items{}
		ResStockRows.Scan(&tmpStock.ID, &tmpStock.Stock)
		arrStockItem = append(arrStockItem, tmpStock)
	}
	StockBarang := arrStockItem[0].Stock
	idBarang := arrStockItem[0].ID
	if StockBarang > 0 {

		//------------------Tidak Ada Error Lanjut Execute----------------

		itemQuery, err := im.DB.Prepare("INSERT INTO transaksis(user_id,item_id,customer_id,create_date) VALUES (?,?,?,now())")
		if err != nil {
			log.Println("prepare insert Transaksi ", err.Error())
			return false, errors.New("** Prepare INSERT ke tabel Transaksi ERROR **")
		}
		res, err := itemQuery.Exec(userID, idBarang, custID)

		if err != nil {
			log.Println("Insert Transaksi ", err.Error())
			return false, errors.New("** Error saat Insert Transaksi **")
		}
		affRows, err := res.RowsAffected()
		if err != nil {
			log.Println("Error setelah insert Transaksi ", err.Error())
			return false, errors.New("error setelah insert")
		}
		if affRows <= 0 {
			log.Println("no record affected")
			return false, errors.New("no record")
		}
		//---------------------------Stock Item Kurangi Satu-----------------------------------
		newStock := StockBarang - 1
		rdcQuery, err := im.DB.Prepare("UPDATE items SET stock = ? WHERE id = ?")
		if err != nil {
			log.Println("prepare insert items ", err.Error())
			return false, errors.New("** Prepare Update Stock ke tabel items ERROR **")
		}
		resReduce, err := rdcQuery.Exec(newStock, idBarang)
		if err != nil {
			log.Println("Update Stock Items ", err.Error())
			return false, errors.New("** Error saat Update Stock Item **")
		}
		affRowsRdcQuery, err := resReduce.RowsAffected()
		if err != nil {
			log.Println("Error setelah update Stock item ", err.Error())
			return false, errors.New("error setelah update Stock")
		}
		if affRowsRdcQuery <= 0 {
			log.Println("no record affected")
			return false, errors.New("no record")
		}
		// fmt.Println(userID, StockBarang, idBarang, custID, "Aman")
		return true, nil
	} else {
		return false, errors.New("** Maaf Stok Barang Habis, Tranksaksi Gagal **")
	}

}

//=======================================================================================TAMPILKAN SEMUA BARANG

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
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t|\t", arrItem[i].Stock, "\t|")
		} else {
			fmt.Println("|  ", i+1, " |\t", arrItem[i].Nama, "\t\t|\t", arrItem[i].Stock, "\t|")

		}
	}
	fmt.Println("|-----------------------------------------------|")
}
