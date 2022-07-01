package entity

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Batch struct {
	UID_Kelas string `gorm:"primaryKey;type:varchar(36);"`
	NamaKelas string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Students  []Student      `gorm:"foreignKey:ID_Kelas"`
}

type AksesKelas struct {
	DB *gorm.DB
}

func (ak *AksesKelas) GetAllKelas() []Batch {
	var daftarKelas = []Batch{}
	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
	err := ak.DB.Find(&daftarKelas)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return daftarKelas
}

func (ak *AksesKelas) TambahKelasBaru(newBatch Batch) Batch {

	uid := uuid.New()
	newBatch.UID_Kelas = uid.String()
	err := ak.DB.Create(&newBatch).Error
	if err != nil {
		log.Println(err)
		return Batch{}
	}

	return newBatch
}

// func (as *AksesStudent) GetSpecificKelas(UID int) Student {
// 	var daftarStudent = Student{}
// 	daftarStudent.ID = uint(UID)
// 	// err := as.DB.Raw("Select * from student").Scan(&daftarStudent)
// 	err := as.DB.First(&daftarStudent)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		return Student{}
// 	}

// 	return daftarStudent
// }

// func (as *AksesStudent) HapusKelas(IDStudent int) bool {
// 	postExc := as.DB.Where("ID = ?", IDStudent).Delete(&Student{})
// 	// ada masalah ga(?)
// 	if err := postExc.Error; err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	// berapa data yang berubah (?)
// 	if aff := postExc.RowsAffected; aff < 1 {
// 		log.Println("Tidak ada data yang dihapus")
// 		return false
// 	}

// 	return true

// }
