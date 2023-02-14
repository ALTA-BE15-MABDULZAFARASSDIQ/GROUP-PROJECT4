package controllers

import (
	"be15/gp4/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetAllTransfer(db *sql.DB) {
	// SELECT DATA
	// select * from guru
	rows, errSelect := db.Query("SELECT ID, Akun_ID_Penerima, Nomor_Rekening, Jumlah_Transfer, Nama_Bank FROM Topup")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allTransfer []entities.Transfer
	for rows.Next() {
		var datarow entities.Transfer
		errScan := rows.Scan(&datarow.ID, &datarow.Akun_ID_Penerima , &datarow.Nomor_Rekening, &datarow.Jumlah_Transfer, &datarow.Nama_Bank,)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allTransfer = append(allTransfer, datarow)
		// fmt.Println(datarow)
	}

	// fmt.Println(allGuru)
	for i, v := range allTransfer {
		fmt.Println("i:", i, "v:", v.ID, v.Akun_ID_Penerima, v.Nomor_Rekening, v.Jumlah_Transfer, v.Nama_Bank)
	}
}

func InsertTransfer(db *sql.DB, newTransfer entities.Transfer) {
	var query = "INSERT INTO guru(ID, Akun_ID_Penerima, Nomor_Rekening, Jumlah_Transfer, Nama_Bank) VALUES(?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newTransfer.ID, newTransfer.Akun_ID_Penerima, newTransfer.Nomor_Rekening, newTransfer.Jumlah_Transfer, newTransfer.Nama_Bank )
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")
		} else {
			fmt.Println("proses gagal")
		}
	}
}