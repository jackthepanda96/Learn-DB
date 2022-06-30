package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	// format koneksi
	// username:password@tcp(IP_Database:Port_Database)/Nama_Database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_be10_beaver")
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
