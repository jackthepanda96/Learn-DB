package entity

import (
	"database/sql"
	"log"
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
	DB *sql.DB
}

func (as *AksesStudent) GetAllData() []Student {
	var studentData Student
	var daftarStudent = []Student{}

	rows, err := as.DB.Query("select * from student")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentData.ID, &studentData.Nama, &studentData.NomorHP, &studentData.Email, &studentData.Alamat, &studentData.ID_Kelas)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		daftarStudent = append(daftarStudent, studentData)
	}

	return daftarStudent
}

func (as *AksesStudent) TambahMuridBaru(newStudent Student) Student {
	qry := "insert into student(Nama, Nomor_HP, Email, Alamat, ID_Kelas) values (?,?,?,?,?);"
	res, err := as.DB.Exec(qry, newStudent.ID, newStudent.Nama, newStudent.NomorHP, newStudent.Email, newStudent.Alamat, newStudent.ID_Kelas)
	if err != nil {
		log.Fatal(err)
		return Student{}
	}
	affected, _ := res.RowsAffected()
	if affected < 1 {
		log.Fatal("Ada masalah -", err)
		return Student{}
	}

	return newStudent

}
