package db

import (
	"go-db/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// format koneksi
	// username:password@tcp(IP_Database:Port_Database)/Nama_Database
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3307)/db_be10_cli?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	// gor, "root:@tcp(127.0.0.1:3306)/db_be10_beaver")
	// defer db.Close()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entity.Batch{}, entity.Student{})
}
