package main

import (
	"fmt"
	"go-db/config"
	"go-db/entity"
)

func main() {
	param := config.ReadFromEnv()
	conn := config.InitDB(param)
	config.MigrateDB(conn)
	aksesStudent := entity.AksesStudent{DB: conn}
	aksesKelas := entity.AksesKelas{DB: conn}
	var input int = 0
	for input != 99 {
		fmt.Println("\tSistem pendataan siswa")
		fmt.Println("1. Tambah Data Siswa")
		fmt.Println("2. Lihat Data Siswa")
		fmt.Println("3. Update Data Siswa")
		fmt.Println("4. Hapus Data Siswa")
		fmt.Println("5. Lihat Siswa Tertentu")
		fmt.Println("6. Tambah Batch")
		fmt.Println("7. Lihat Data Batch")
		fmt.Println("99. Keluar")
		fmt.Print("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var newStudent entity.Student
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newStudent.Nama)
			fmt.Print("Masukkan alamat: ")
			fmt.Scanln(&newStudent.Alamat)
			fmt.Print("Masukkan nomorhp: ")
			fmt.Scanln(&newStudent.NomorHP)
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&newStudent.Email)
			fmt.Print("Masukkan ID_Kelas: ")
			fmt.Scanln(&newStudent.ID_Kelas)
			res := aksesStudent.TambahMuridBaru(newStudent)
			if res.ID == 0 {
				fmt.Println("Tidak bisa input siswa, ada error")
				break
			}
			fmt.Println("Berhasl input siswa")
		case 2:
			fmt.Println("Daftar Seluruh Murid")
			for _, val := range aksesStudent.GetAllData() {
				fmt.Println(val)
			}
		case 4:
			var ID int
			fmt.Print("Masukkan ID yang akan dihapus ")
			fmt.Scanln(&ID)
			fmt.Println(aksesStudent.HapusMurid(ID))
		case 5:
			var ID int
			fmt.Print("Masukkan ID yang akan dicari ")
			fmt.Scanln(&ID)
			fmt.Println(aksesStudent.GetSpecificUser(ID))
		case 6:
			var newStudent entity.Batch
			fmt.Print("Masukkan nama: ")
			fmt.Scanln(&newStudent.NamaKelas)

			aksesKelas.TambahKelasBaru(newStudent)

			fmt.Println("Berhasl input kelas")
		case 7:
			fmt.Println("Daftar Seluruh Kelas")
			fmt.Println(aksesKelas.GetAllKelas())
		default:
			continue
		}
	}
	fmt.Println("Terima kasih sudah mencoba program saya")

	// fmt.Println("Nilai dari ping", conn.Ping())
	// var studentData entity.Student
	// var daftarStudent = []entity.Student{}

	// rows, err := conn.Query("select * from student")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&studentData.ID, &studentData.Nama, &studentData.NomorHP, &studentData.Email, &studentData.Alamat, &studentData.ID_Kelas)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	daftarStudent = append(daftarStudent, studentData)
	// }
	// fmt.Println(daftarStudent)

	// qry := "insert into student(ID, Nama, Nomor_HP, Email, Alamat, ID_Kelas) values (?,?,?,?,?,?);"
	// res, err := conn.Exec(qry, 20, "Doni", "012345678", "Doni@student.com", "Jakarta", 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// affected, _ := res.RowsAffected()
	// if affected > 0 {
	// 	fmt.Println("Berhasil insert")
	// }else{
	// 	fmt.Println("Tidak ada yang diupdate")
	// }
	// stmt, err := conn.Prepare(qry)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // res, err := stmt.Exec(20, "Doni", "012345678", "Doni@student.com", "Jakarta", 7)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// res2, err := stmt.Exec(21, "Dono", "012345678", "Doni@student.com", "Jakarta", 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var nama, alamat, nomorhp, email string
	// var id = 30
	// fmt.Print("Masukkan nama: ")
	// fmt.Scanln(&nama)
	// fmt.Print("Masukkan alamat: ")
	// fmt.Scanln(&alamat)
	// fmt.Print("Masukkan nomorhp: ")
	// fmt.Scanln(&nomorhp)
	// fmt.Print("Masukkan email: ")
	// fmt.Scanln(&email)

	// res, err := stmt.Exec(id, nama, nomorhp, email, alamat, 7)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
