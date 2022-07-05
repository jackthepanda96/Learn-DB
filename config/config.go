package config

import (
	"fmt"
	"go-db/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(s Setting) *gorm.DB {
	// format koneksi
	// username:password@tcp(IP_Database:Port_Database)/Nama_Database
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", s.DBUser, s.DBPassword)
	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entity.Batch{}, entity.Student{})
}

type Setting struct {
	DBUser     string
	DBPassword string
}

func ReadFromEnv() Setting {
	res := Setting{}
	err := godotenv.Load("local.env")
	if err != nil {
		return Setting{}
	}
	res.DBPassword = os.Getenv("DBPass")
	res.DBUser = os.Getenv("DBUser")
	return res
}
