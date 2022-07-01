package entity

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	NIM      string
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
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.Find(&daftarStudent)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarStudent
}

func (as *AksesStudent) TambahMuridBaru(newStudent Student) Student {
	if newStudent.Nama == "Jerry" {
		newStudent.ID = uint(1)
	}
	uid := uuid.New()
	newStudent.NIM = uid.String()
	err := as.DB.Create(&newStudent).Error
	if err != nil {
		log.Println(err)
		return Student{}
	}

	return newStudent
}

func (as *AksesStudent) GetSpecificUser(UID int) Student {
	var daftarStudent = Student{}
	daftarStudent.ID = uint(UID)
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := as.DB.First(&daftarStudent)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return Student{}
	}

	return daftarStudent
}

func (as *AksesStudent) HapusMurid(IDStudent int) bool {
	postExc := as.DB.Where("ID = ?", IDStudent).Delete(&Student{})
	// ada masalah ga(?)
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	// berapa data yang berubah (?)
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}

	return true

}
