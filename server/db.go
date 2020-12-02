package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

type dbManager struct {
	db *sql.DB
}

func initializeDB() (dm dbManager) {
	// connect (or create) db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// create table
	_, err = db.Exec("CREATE TABLE sensor_data (" +
		"id INTEGER PRIMARY KEY," +
		"dht22_Humi DECIMAL(4,2) NOT NULL," +
		"dht22_Temp DECIMAL(4,2) NOT NULL," +
		"bmp180_Temp DECIMAL(4,2) NOT NULL," +
		"bmp180_Pres DECIMAL(6,2) NOT NULL," +
		"datetime DATETIME NOT NULL)")
	if err != nil {
		log.Print(err)
		log.Printf("\n")
	}

	dm = dbManager{db}

	return
}

func (dm dbManager) addData(data Data) {
	query := fmt.Sprintf("INSERT INTO sensor_data (dht22_Humi, dht22_Temp, bmp180_Temp, bmp180_Pres, datetime)"+
		"VALUES (%f, %f, %f, %f, datetime('now'))", data.Dht22_Humi, data.Dht22_Temp, data.Bmp180_Temp, data.Bmp180_Pres)
	_, err := dm.db.Exec(query)
	if err != nil {
		log.Print(err)
		log.Printf("\n")
	}
}

func (dm dbManager) getRecentDate(n int) (data Data) {
	query := fmt.Sprintf("SELECT * FROM sensor_data ORDER BY id DESC LIMIT %d", n)

	rows, err := dm.db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Print(err)
		log.Printf("\n")
	}

	for rows.Next() {
		err := rows.Scan(&data.Id, &data.Dht22_Humi, &data.Dht22_Temp, &data.Bmp180_Temp, &data.Bmp180_Pres, &data.Datetime)
		if err != nil {
			fmt.Println(err)
		}
	}

	return
}
