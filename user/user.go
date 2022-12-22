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
	ID     int
	Nama   string
	Stock  int
	userID int
}
type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Login(nama string, password string) (User, error) {
	userQuery, err := am.DB.Prepare("SELECT role FROM users WHERE nama=? AND password=?")
	if err != nil {
		log.Println("Read Data with Prepare function Error", err.Error())
		return User{}, errors.New("SELECT prepare statement Error")
	}
	row := userQuery.QueryRow(nama, password)
	if row.Err() != nil {

		log.Println("User Query data not found", row.Err().Error())
		return User{}, errors.New("Login Failed, Name and/or Password not found")
	}
	dataRes := User{}
	err = row.Scan((&dataRes.Role)) // <-- Data Yang akan menjadi Return Value

	if err != nil {
		log.Println("Data not found", err.Error())
		return User{}, errors.New("** Wrong Name and/or Password, Please try again **")
	}
	dataRes.Nama = nama
	return dataRes, nil

}

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id FROM users where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("New account successfully created", err.Error())
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
		return false, errors.New("name already exists")
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
		return false, errors.New("error after insert")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}

func (am *AuthMenu) TampilItem() {
	resultRows, err := am.DB.Query("SELECT * FROM items")
	if err != nil {
		fmt.Println("Error Reading Data from Database", err.Error())
	}
	arrItem := []Items{}
	for resultRows.Next() {
		tmp := Items{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Stock, &tmp.userID)
		arrItem = append(arrItem, tmp)
	}
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|  ID  |\t Name\t\t|\tStock   |")
	fmt.Println("|-----------------------------------------------|")
	for i := 0; i < len(arrItem); i++ {
		if len(arrItem[i].Nama) > 5 {
			fmt.Println("|  ", arrItem[i].ID, " |\t", arrItem[i].Nama, "\t|\t", arrItem[i].Stock, "\t|")
		} else {
			fmt.Println("|  ", arrItem[i].ID, " |\t", arrItem[i].Nama, "\t\t|\t", arrItem[i].Stock, "\t|")

		}
	}
	fmt.Println("|-----------------------------------------------|")
}

func (am *AuthMenu) TampilPegawai() {
	resultRows, err := am.DB.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error Reading Data from Database", err.Error())
	}
	arrUser := []User{}
	for resultRows.Next() {
		tmp := User{}
		resultRows.Scan(&tmp.ID, &tmp.Nama, &tmp.Password, &tmp.Role)
		arrUser = append(arrUser, tmp)
	}
	fmt.Println("|-------------------------------|")
	fmt.Println("|         EMPLOYEE TABLE       	|")
	fmt.Println("|-------------------------------|")
	fmt.Println("|  ID\t|      Employee Name\t|")
	fmt.Println("|-------------------------------|")
	for i := 0; i < len(arrUser); i++ {
		if arrUser[i].Role > 1 {
			if len(arrUser[i].Nama) <= 12 {
				fmt.Println("| ", arrUser[i].ID, "\t| ", arrUser[i].Nama, "\t\t| ")
			} else {
				fmt.Println("| ", arrUser[i].ID, "\t| ", arrUser[i].Nama, "\t| ")
			}

		}

	}
	fmt.Println("|-------------------------------|")
}

// -------------------------------------------------------------------------------------------------
func (am *AuthMenu) CekUser(userID int) bool {
	res := am.DB.QueryRow("SELECT id FROM users where id = ?", userID)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil { // err hanya bernilai nil & bukan nil
		log.Println("Wrong ID, Please input the correct ID", err.Error())
		return true
	}
	return false
}

func (am *AuthMenu) HapusPegawai(hapusPegawai int) (bool, error) {
	//--------------------------------Cek ID -----------------------------------
	if am.CekUser(hapusPegawai) {
		log.Println("ID not registered")
		return false, errors.New("-Please input the correct ID-")
	}
	//--------------------------------Cek Role----------------------------------
	resultRows, err := am.DB.Query("SELECT role FROM users WHERE id = ?", hapusPegawai)
	if err != nil {
		fmt.Println("Error Reading Data from Database", err.Error())
	}
	arrUser := []User{}
	for resultRows.Next() {
		tmp := User{}
		resultRows.Scan(&tmp.Role)
		arrUser = append(arrUser, tmp)
	}
	role := arrUser[0].Role
	// fmt.Println(arrUser[0].Role)

	if role == 2 {
		delQry, err := am.DB.Prepare("DELETE FROM users WHERE id = ?")
		if err != nil {
			log.Println("prepare delete employee ", err.Error())
			return false, errors.New("prepare statement delete employee error")
		}
		// menjalankan query dengan parameter tertentu
		res, err := delQry.Exec(hapusPegawai)
		if err != nil {
			log.Println("delete employee", err.Error())
			return false, errors.New("delete employee error")
		}
		// Cek berapa baris yang terpengaruh query diatas
		affRows, err := res.RowsAffected()

		if err != nil {
			log.Println("after delete employee ", err.Error())
			return false, errors.New("error after delete")
		}
		if affRows <= 0 {
			log.Println("no record affected")
			return false, errors.New("no record")
		}
		return true, nil
	} else {
		return false, errors.New("*** Please input ID according to the table ***")
	}

}

// id : namar := arrUser[0].Password

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
		return false, errors.New("error after delete")
	}
	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}
	return true, nil
}
