package transaksi

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type Transaksi struct {
	ID          int
	User_ID     int
	Item_ID     int
	Customer_ID int
	Create_Date string
}

type ModifTransaksi struct {
	ID           int
	NamaKasir    string
	NamaBarang   string
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
		resultRows.Scan(&tmp.ID, &tmp.User_ID, &tmp.Item_ID, &tmp.Customer_ID, &tmp.Create_Date)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|                           TABEL TRANSAKSI                             |")
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|  No  | Nama Pegawai\t| Nama Barang\t| Nama Customer\t| Tgl Pembelian |")
	fmt.Println("|-----------------------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("|  ", i+1, " | ", arrTrans[i].User_ID, "\t\t|\t", arrTrans[i].Item_ID, "\t| ", arrTrans[i].Customer_ID, "\t\t|", arrTrans[i].Create_Date, "\t|")

	}
	fmt.Println("|-----------------------------------------------------------------------|")
}

func (tm *TransMenu) TampilHapusTransaksi() {
	resultRows, err := tm.DB.Query("SELECT t.id, u.nama ,i.nama,c.nama,create_date FROM transaksis t JOIN users u ON u.id = user_id JOIN items i ON i.id = item_id JOIN customers c ON c.id = customer_id;")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []ModifTransaksi{}
	for resultRows.Next() {
		tmp := ModifTransaksi{}
		resultRows.Scan(&tmp.ID, &tmp.NamaKasir, &tmp.NamaBarang, &tmp.NamaCustomer, &tmp.CreateDate)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|                           TABEL TRANSAKSI                             |")
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|  ID  | Nama Pegawai\t| Nama Barang\t| Nama Customer\t| Tgl Pembelian |")
	fmt.Println("|-----------------------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("|  ", arrTrans[i].ID, " | ", arrTrans[i].NamaKasir, "\t| ", arrTrans[i].NamaBarang, "\t| ", arrTrans[i].NamaCustomer, "\t|", arrTrans[i].CreateDate, "\t|")

	}
	fmt.Println("|-----------------------------------------------------------------------|")
}

func (tm *TransMenu) HapusTransaksi(hapusTransaksi int) (bool, error) {

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
	resultRows, err := tm.DB.Query("SELECT u.nama ,i.nama,c.nama,create_date FROM transaksis t JOIN users u ON u.id = user_id JOIN items i ON i.id = item_id JOIN customers c ON c.id = customer_id;")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrTrans := []ModifTransaksi{}
	for resultRows.Next() {
		tmp := ModifTransaksi{}
		resultRows.Scan(&tmp.NamaKasir, &tmp.NamaBarang, &tmp.NamaCustomer, &tmp.CreateDate)
		arrTrans = append(arrTrans, tmp)
	}
	// id := arrTrans[0].Nama
	// namar := arrTrans[0].Password
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|                           TABEL TRANSAKSI                             |")
	fmt.Println("|-----------------------------------------------------------------------|")
	fmt.Println("|  No  | Nama Pegawai\t| Nama Barang\t| Nama Customer\t| Tgl Pembelian |")
	fmt.Println("|-----------------------------------------------------------------------|")
	for i := 0; i < len(arrTrans); i++ {

		fmt.Println("|  ", i+1, " | ", arrTrans[i].NamaKasir, "\t| ", arrTrans[i].NamaBarang, "\t| ", arrTrans[i].NamaCustomer, "\t|", arrTrans[i].CreateDate, "\t|")

	}
	fmt.Println("|-----------------------------------------------------------------------|")
}
