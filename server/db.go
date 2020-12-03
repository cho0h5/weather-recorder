package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
		"VALUES (%f, %f, %f, %f, datetime('now', 'localtime'))", data.Dht22_Humi, data.Dht22_Temp, data.Bmp180_Temp, data.Bmp180_Pres)
	_, err := dm.db.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func (dm dbManager) getRecentDate(n int) (data []Data) {
	query := fmt.Sprintf("SELECT * FROM sensor_data ORDER BY id DESC LIMIT %d", n)

	rows, err := dm.db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}

	var tempData Data
	for rows.Next() {
		err := rows.Scan(&tempData.Id,
			&tempData.Dht22_Humi,
			&tempData.Dht22_Temp,
			&tempData.Bmp180_Temp,
			&tempData.Bmp180_Pres,
			&tempData.Datetime)
		if err != nil {
			log.Println(err)
		}

		recentTime, _ := time.Parse(time.RFC3339, tempData.Datetime)

		tempData.IsWorking = checkIsWorking(recentTime)

		data = append(data, tempData)
	}

	return
}

func checkIsWorking(recentTime time.Time) bool {
	currentTime := time.Now()
	deltaTime := currentTime.Sub(recentTime)

	if deltaTime < time.Second*3 {
		return true
	}
	return false
}
