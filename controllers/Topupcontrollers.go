package controllers

import (
	"be15/gp4/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetAllTopup(db *sql.DB) {
	// SELECT DATA
	// select * from guru
	rows, errSelect := db.Query("SELECT ID, , AccountID, Jumlah_Topup, Tanggal_Topup, FROM Topup")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allTopup []entities.Topup
	for rows.Next() {
		var datarow entities.Topup
		errScan := rows.Scan(&datarow.ID, &datarow.AccountID, &datarow.Jumlah_Topup, &datarow.Tanggal_Topup,)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allTopup = append(allTopup, datarow)
		// fmt.Println(datarow)
	}

	// fmt.Println(allGuru)
	for i, v := range allTopup {
		fmt.Println("i:", i, "v:", v.ID, v.AccountID, v.Jumlah_Topup, v.Tanggal_Topup)
	}
}

func InsertTopup(db *sql.DB, newTopup entities.Topup) {
	var query = "INSERT INTO guru(ID, , AccountID, Jumlah_Topup, Tanggal_Topup) VALUES(?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newTopup.ID, newTopup.Jumlah_Topup, newTopup.AccountID, newTopup.Tanggal_Topup )
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