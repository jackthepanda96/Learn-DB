package main

import (
	"fmt"
	"go-db/db"
	"go-db/entity"
	"log"
)

func main() {
	conn := db.InitDB()

	var studentData entity.Student
	var daftarStudent = []entity.Student{}

	rows, err := conn.Query("select * from student")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentData.ID, &studentData.Nama, &studentData.NomorHP, &studentData.Email, &studentData.Alamat, &studentData.ID_Kelas)
		if err != nil {
			log.Fatal(err)
		}
		daftarStudent = append(daftarStudent, studentData)
	}
	qry := "insert into student(ID, Nama, Nomor_HP, Email, Alamat, ID_Kelas) values (?,?,?,?,?,?)"
	stmt, err := conn.Prepare(qry)
	if err != nil {
		log.Fatal(err)
	}
	// res, err := stmt.Exec(20, "Doni", "012345678", "Doni@student.com", "Jakarta", 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := conn.Exec(qry, 20, "Doni", "012345678", "Doni@student.com", "Jakarta", 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res2, err := stmt.Exec(21, "Dono", "012345678", "Doni@student.com", "Jakarta", 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var nama, alamat, nomorhp, email string
	var id = 30
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan alamat: ")
	fmt.Scanln(&alamat)
	fmt.Print("Masukkan nomorhp: ")
	fmt.Scanln(&nomorhp)
	fmt.Print("Masukkan email: ")
	fmt.Scanln(&email)

	res, err := stmt.Exec(id, nama, nomorhp, email, alamat, 7)
	if err != nil {
		log.Fatal(err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	if affected > 0 {
		fmt.Println("Berhasil insert")
	}

}
