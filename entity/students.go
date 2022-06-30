package entity

import (
	"log"

	"gorm.io/gorm"
)

type Student struct {
	ID       int
	Nama     string
	Email    string
	NomorHP  string
	Alamat   string
	ID_Kelas int
}

type AksesStudent struct {
	DB *gorm.DB
}

func (as *AksesStudent) GetAllData() []Student {
	var daftarStudent = []Student{}
	err := as.DB.Find(&daftarStudent).Error
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return daftarStudent
}

func (as *AksesStudent) TambahMuridBaru(newStudent Student) Student {
	err := as.DB.Create(newStudent).Error
	if err != nil {
		log.Fatal(err)
		return Student{}
	}

	return newStudent

}
