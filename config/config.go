package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBPort int
}

func ReadConfig() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Error saat baca file envv", err.Error())
		return nil
	}
	var read Config
	read.DBUser = os.Getenv("DBUSER")
	read.DBPass = os.Getenv("DBPASS")
	read.DBHost = os.Getenv("DBHOST")
	read.DBName = os.Getenv("DBNAME")
	convData := os.Getenv("DBPORT")
	read.DBPort, err = strconv.Atoi(convData)
	if err != nil {
		fmt.Println("Error Saat Conversi Port String ke Int", err.Error())
	}
	return &read
}

func ConnectSQL(c Config) *sql.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		fmt.Println("Error Saat Connect Ke SQL", err.Error())
	}
	return db
}
