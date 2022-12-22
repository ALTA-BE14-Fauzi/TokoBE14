package transaksi

import (
	"TokoBE14/items"
	"TokoBE14/user"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Transaksi struct {
	ID          int
	User_ID     int
	Customer_ID int
	Create_Date string
}

type ModifTransaksi struct {
	ID           int
	NamaKasir    string
	NamaCustomer string
	CreateDate   string
}

type Customer struct {
	ID   int
	Nama string
}

type TransMenu struct {
	DB *sql.DB
}

func (tm *TransMenu) TampilTransaksi() {
	resultRows, err := tm.DB.Query("SELECT * FROM transaksis ")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []Transaksi{}
	for resultRows.Next() {
		tmp := Transaksi{}
		resultRows.Scan(&tmp.ID, &tmp.User_ID, &tmp.Customer_ID, &tmp.Create_Date)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|------------------------------------------------------|")
	fmt.Println("|                     TRANSAKSI                        |")
	fmt.Println("|------------------------------------------------------|")
	fmt.Println("|  No  | Nama Pegawai\t| Nama Customer\t| Tgl Pembelian |")
	fmt.Println("|------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("|  ", i+1, " | ", arrTrans[i].User_ID, "\t\t| ", arrTrans[i].Customer_ID, "\t\t|", arrTrans[i].Create_Date, "\t|")

	}
	fmt.Println("|------------------------------------------------------|")
}

func (tm *TransMenu) TampilHapusTransaksi() {
	resultRows, err := tm.DB.Query("SELECT t.id, u.nama,c.nama,create_date FROM transaksis t JOIN users u ON u.id = user_id JOIN customers c ON c.id = customer_id;")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []ModifTransaksi{}
	for resultRows.Next() {
		tmp := ModifTransaksi{}
		resultRows.Scan(&tmp.ID, &tmp.NamaKasir, &tmp.NamaCustomer, &tmp.CreateDate)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("|                     TABEL TRANSAKSI                   |")
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("|  ID  | Nama Pegawai\t| Nama Customer\t| Tgl Pembelian |")
	fmt.Println("|-------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("|  ", arrTrans[i].ID, " | ", arrTrans[i].NamaKasir, "\t| ", arrTrans[i].NamaCustomer, "\t|", arrTrans[i].CreateDate, "\t|")

	}
	fmt.Println("|-------------------------------------------------------|")
}

func (tm *TransMenu) HapusTransaksi(hapusTransaksi int) (bool, error) {
	//Delete Transaksi_items
	delQryId, err := tm.DB.Prepare("DELETE FROM transaksi_items WHERE transaction_id = ?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}
	// menjalankan query dengan parameter tertentu
	resId, err := delQryId.Exec(hapusTransaksi)
	if err != nil {
		log.Println("delete transaksi", err.Error())
		return false, errors.New("delete transaksi error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRowsId, err := resId.RowsAffected()

	if err != nil {
		log.Println("after delete transaksi ", err.Error())
		return false, errors.New("error setelah delete")
	}
	if affRowsId <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	//Delete Transaksis
	delQry, err := tm.DB.Prepare("DELETE FROM transaksis WHERE id = ?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}
	// menjalankan query dengan parameter tertentu
	res, err := delQry.Exec(hapusTransaksi)
	if err != nil {
		log.Println("delete transaksi", err.Error())
		return false, errors.New("delete transaksi error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete transaksi ", err.Error())
		return false, errors.New("error setelah delete")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil

}

func (tm *TransMenu) TampilTransaksiModif() {
	resultRows, err := tm.DB.Query("SELECT t.id,u.nama ,c.nama,create_date FROM transaksis t JOIN users u ON u.id = user_id JOIN customers c ON c.id = customer_id;")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []ModifTransaksi{}
	for resultRows.Next() {
		tmp := ModifTransaksi{}
		resultRows.Scan(&tmp.ID, &tmp.NamaKasir, &tmp.NamaCustomer, &tmp.CreateDate)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("|                     TABEL TRANSAKSI                   |")
	fmt.Println("|-------------------------------------------------------|")
	fmt.Println("|  ID   | Nama Pegawai\t| Nama Customer\t| Tgl Transaksi |")
	fmt.Println("|-------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("| ", arrTrans[i].ID, "\t| ", arrTrans[i].NamaKasir, "\t| ", arrTrans[i].NamaCustomer, "\t|", arrTrans[i].CreateDate, "\t|")

	}
	fmt.Println("|-------------------------------------------------------|")
}

func (tm *TransMenu) TampilCustomer() {
	resultRows, err := tm.DB.Query("SELECT * FROM customers ")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrCust := []Customer{}
	for resultRows.Next() {
		tmp := Customer{}
		resultRows.Scan(&tmp.ID, &tmp.Nama)
		arrCust = append(arrCust, tmp)
	}
	// id := arrItem[0].Nama
	// namar := arrItem[0].Password
	fmt.Println("|-------------------------------|")
	fmt.Println("|  ID  |\t Nama\t\t|")
	fmt.Println("|-------------------------------|")
	for i := 0; i < len(arrCust); i++ {
		fmt.Println("|  ", arrCust[i].ID, " |\t", arrCust[i].Nama, "\t|")
	}
	fmt.Println("|-------------------------------|")
}

func (tm *TransMenu) HapusCustomer(HapusCustomer int) (bool, error) {

	delQry, err := tm.DB.Prepare("DELETE FROM customers WHERE id = ?")
	if err != nil {
		log.Println("prepare delete customer ", err.Error())
		return false, errors.New("prepare statement delete customer error")
	}
	// menjalankan query dengan parameter tertentu
	res, err := delQry.Exec(HapusCustomer)
	if err != nil {
		log.Println("delete customer", err.Error())
		return false, errors.New("delete customer error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete customer ", err.Error())
		return false, errors.New("error setelah delete")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil

}

// ===========================================================================================BUAT TRANSAKSI
func (tm *TransMenu) CekBarang(namaBarang string) bool {
	res := tm.DB.QueryRow("SELECT id FROM items where nama = ?", namaBarang)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Barang Tidak Tersedia Ditoko ini", err.Error())
		return true
	}
	return false
}
func (tm *TransMenu) CekCustomer(namaCustomer string) bool {
	res := tm.DB.QueryRow("SELECT id FROM customers where nama = ?", namaCustomer)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Customer Belum Terdaftar", err.Error())
		return true
	}
	return false
}

func (tm *TransMenu) BuatTransaksi(nama string, namaCustomer string) (bool, error) {
	//-------------------------CARI ID USER---------------------
	resultRows, err := tm.DB.Query("SELECT id FROM users WHERE nama =?", nama)
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
	if tm.CekCustomer(namaCustomer) {
		log.Println("--- Empty Data ---")
		return false, errors.New("--Customer Belum Terdaftar--")
	}
	//-----------------------Cari ID Customer--------------------
	resIDCust, err := tm.DB.Query("SELECT id FROM customers WHERE nama =?", namaCustomer)
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

	//------------------Tidak Ada Error Lanjut Execute----------------

	itemQuery, err := tm.DB.Prepare("INSERT INTO transaksis(user_id,customer_id,create_date) VALUES (?,?,now())")
	if err != nil {
		log.Println("prepare insert Transaksi ", err.Error())
		return false, errors.New("** Prepare INSERT ke tabel Transaksi ERROR **")
	}
	res, err := itemQuery.Exec(userID, custID)

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

	// fmt.Println(userID, StockBarang, idBarang, custID, "Aman")
	return true, nil
}

func (tm *TransMenu) BuatTransaksiItems(namaBarang string) (bool, error) {
	//-----------------------Ambil ID Transaksi-------------------
	TransRows, err := tm.DB.Query("SELECT id FROM transaksis ")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []Transaksi{}
	for TransRows.Next() {
		tmpStock := Transaksi{}
		TransRows.Scan(&tmpStock.ID)
		arrTrans = append(arrTrans, tmpStock)
	}
	transID := arrTrans[len(arrTrans)-1].ID
	//------------------------Cek Barang-----------------------
	if tm.CekBarang(namaBarang) {
		log.Println("--- Empty Item ---")
		return false, errors.New("--Barang Tidak Tersedia--")
	}

	//---------------------Cek Stok & ID----------------------
	ResStockRows, err := tm.DB.Query("SELECT id,stock FROM items WHERE nama =?", namaBarang)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrStockItem := []items.Items{}
	for ResStockRows.Next() {
		tmpStock := items.Items{}
		ResStockRows.Scan(&tmpStock.ID, &tmpStock.Stock)
		arrStockItem = append(arrStockItem, tmpStock)
	}
	StockBarang := arrStockItem[0].Stock
	idBarang := arrStockItem[0].ID
	if StockBarang > 0 {
		//--------------------------Input Data Ke Transaksi Items-----------------------------
		itemQuery, err := tm.DB.Prepare("INSERT INTO transaksi_items(transaction_id,item_id) VALUES (?,?)")
		if err != nil {
			log.Println("prepare insert Transaksi ", err.Error())
			return false, errors.New("** Prepare INSERT ke tabel Transaksi ERROR **")
		}
		res, err := itemQuery.Exec(transID, idBarang)

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
		rdcQuery, err := tm.DB.Prepare("UPDATE items SET stock = ? WHERE id = ?")
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
		//------------------------------Buat Transaksi Items-------------------------------

		return true, nil

	} else {
		return false, errors.New("** Maaf Stok Barang Habis, Tranksaksi Gagal **")
	}
}

type ViewTransaksiItemStruct struct {
	ID           int
	NamaPegawai  string
	NamaCustomer string
	NamaBarang   string
	Quantity     string
	CreateDate   string
}

func (tm *TransMenu) CekTransaksiID(id int) bool {
	res := tm.DB.QueryRow("SELECT id FROM transaksis where id = ?", id)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("ID Transaksi tidak ada", err.Error())
		return true
	}
	return false
}

func (tm *TransMenu) TranksaksiItem(id int) (bool, error) {
	if tm.CekTransaksiID(id) {
		log.Println("--- ID Transaksi tidak ada  ---")
		return false, errors.New("ID kosong")
	}
	return true, nil
}

func (tm *TransMenu) ViewTransaksiItem(id int) {
	resultRows, err := tm.DB.Query("SELECT t.id,u.nama, c.nama,i.nama,COUNT(*),t.create_date FROM transaksi_items JOIN transaksis t ON t.id = transaction_id JOIN customers c ON c.id = customer_id JOIN items i ON i.id = item_id 	jOIN users u ON u.id = t.user_id WHERE t.id = ? group by t.id,c.nama,i.nama,t.create_date having COUNT(*) >= 1;", id)
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []ViewTransaksiItemStruct{}
	for resultRows.Next() {
		tmp := ViewTransaksiItemStruct{}
		resultRows.Scan(&tmp.ID, &tmp.NamaPegawai, &tmp.NamaCustomer, &tmp.NamaBarang, &tmp.Quantity, &tmp.CreateDate)
		arrTrans = append(arrTrans, tmp)
	}
	// fmt.Println(arrTrans)
	fmt.Println("|-------------------------------------------------------------------------------|")
	fmt.Println("|                               TABEL TRANSAKSI ITEM                            |")
	fmt.Println("|-------------------------------------------------------------------------------|")
	fmt.Println("| ID\t| Nama Pegawai\t| Nama Customer\t| Nama Barang\t| Qty\t| Tgl Transaksi |")
	fmt.Println("|-------------------------------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {
		if i > 0 {
			if arrTrans[i].ID == arrTrans[i-1].ID {
				tmpStr := ""
				fmt.Println("| ", tmpStr, "\t| ", tmpStr, "\t\t| ", tmpStr, "\t\t| ", arrTrans[i].NamaBarang, "\t| ", arrTrans[i].Quantity, "\t| ", tmpStr, "\t\t|")
			}

		} else {
			fmt.Println("| ", arrTrans[i].ID, "\t| ", arrTrans[i].NamaPegawai, "\t| ", arrTrans[i].NamaCustomer, "\t| ", arrTrans[i].NamaBarang, "\t| ", arrTrans[i].Quantity, "\t| ", arrTrans[i].CreateDate, "  |")
		}

	}
	fmt.Println("|-------------------------------------------------------------------------------|")
}
