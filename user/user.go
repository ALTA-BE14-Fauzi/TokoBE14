package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type User struct {
	ID       int
	Nama     string
	Password string
	Role     int
}

type Items struct {
	ID    int
	Nama  string
	Stock int
}
type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Login(nama string, password string) (User, error) {
	userQuery, err := am.DB.Prepare("SELECT role FROM users WHERE nama=? AND password=?")
	if err != nil {
		log.Println("Ambil Data dengan fungsi Prepare Error", err.Error())
		return User{}, errors.New("SELECT prepare statment Error")
	}
	row := userQuery.QueryRow(nama, password)
	if row.Err() != nil {

		log.Println("User Query data tidak ditemukan", row.Err().Error())
		return User{}, errors.New("Login Gagal, Nama dan Password tidak terdaftar")
	}
	dataRes := User{}
	err = row.Scan((&dataRes.Role)) // <-- Data Yang akan menjadi Return Value

	if err != nil {
		log.Println("Data Tidak Ditemukan", err.Error())
		return User{}, errors.New("** Nama dan Password Salah , Silahkan Coba kembali **")
	}
	dataRes.Nama = nama
	return dataRes, nil

}

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM users where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Akun Baru Berhasil Dibuat", err.Error())
		return false
	}
	return true
}

func (am *AuthMenu) Register(newUser User) (bool, error) {
	registerQry, err := am.DB.Prepare("INSERT INTO users (nama, password,role) values (?,?,?)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}
	if am.Duplicate(newUser.Nama) {
		log.Println("--- Duplicated information ---")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newUser.Nama, newUser.Password, 2)
	if err != nil {
		log.Println("insert user ", err.Error())
		return false, errors.New("insert user error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert user ", err.Error())
		return false, errors.New("error setelah insert")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

// func (am *AuthMenu) HapusPegawai(id int) (bool, error) {
// 	fmt.Println("DELETE FROM users WHERE id=?")
// 	registerQry, err := am.DB.Prepare("INSERT INTO users (nama, password,role) values (?,?,?)")
// 	if err != nil {
// 		log.Println("prepare insert user ", err.Error())
// 		return false, errors.New("prepare statement insert user error")
// 	}
// 	// menjalankan query dengan parameter tertentu
// 	res, err := registerQry.Exec(id)
// 	if err != nil {
// 		log.Println("insert user ", err.Error())
// 		return false, errors.New("insert user error")
// 	}
// 	// Cek berapa baris yang terpengaruh query diatas
// 	affRows, err := res.RowsAffected()

// 	if err != nil {
// 		log.Println("after insert user ", err.Error())
// 		return false, errors.New("error setelah insert")
// 	}
// 	if affRows <= 0 {
// 		log.Println("no record affected")
// 		return false, errors.New("no record")
// 	}
// 	return true, nil
// }

func (am *AuthMenu) Tampilkan(nama string, password string) {
	resultRows, err := am.DB.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println("Ambil Data dari Database Error", err.Error())
	}
	arrUser := []Items{}
	for resultRows.Next() {
		tmp := Items{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Stock)
		arrUser = append(arrUser, tmp)
	}
	// id := arrUser[0].Nama
	// namar := arrUser[0].Password
	fmt.Println(arrUser)

}

func (am *AuthMenu) HapusItem(hapusItem Items) (bool, error) {

	delQry, err := am.DB.Prepare("DELETE FROM items WHERE id = ?")
	if err != nil {
		log.Println("prepare delete item ", err.Error())
		return false, errors.New("prepare statement delete item error")
	}
	// menjalankan query dengan parameter tertentu
	res, err := delQry.Exec(hapusItem.ID)
	if err != nil {
		log.Println("delete item", err.Error())
		return false, errors.New("delete item error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete item ", err.Error())
		return false, errors.New("error setelah delete")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}
