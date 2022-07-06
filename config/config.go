package config

import (
	"fmt"
	"go-db/entity"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Setting struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int16
	DBName     string
}

func ReadFromEnv() Setting {
	res := Setting{}
	err := godotenv.Load("local.env")
	if err != nil {
		return Setting{}
	}
	res.DBUser = os.Getenv("DBUser")
	res.DBPassword = os.Getenv("DBPass")
	res.DBHost = os.Getenv("DBHost")
	portConv, _ := strconv.Atoi(os.Getenv("DBPort"))
	res.DBPort = int16(portConv)
	res.DBName = os.Getenv("DBName")
	return res
}

func InitDB(s Setting) *gorm.DB {
	// format koneksi
	// username:password@tcp(IP_Database:Port_Database)/Nama_Database
	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		s.DBUser,
		s.DBPassword,
		s.DBHost,
		s.DBPort,
		s.DBName)
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
