package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

type dbManager struct {
	db *sql.DB
}

func initializeDB() (db *sql.DB) {
	// connect (or create) db
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// create table
	db.Exec("CREATE TABLE sensor_data (" +
		"id INT PRIMARY KEY AUTO_INCREMENT," +
		"dht22_Humi DECIMAL(4,2) NOT NULL," +
		"dht22_Temp DECIMAL(4,2) NOT NULL," +
		"bmp180_Temp DECIMAL(4,2) NOT NULL," +
		"bmp180_Pres DECIMAL(6,2) NOT NULL," +
		"datetime DATETIME NOT NULL)")

	return
}

func (dm dbManager) addData(data Data) {
	query := fmt.Sprintf("INSERT INTO sensor_data (dht22_Humi, dht22_Temp, bmp180_Temp, bmp180_Pres)"+
		"VALUES (%f, %f, %f, %f, NOW())", data.Dht22_Humi, data.Dht22_Temp, data.Bmp180_Temp, data.Bmp180_Pres)
	_, err := dm.db.Exec(query)
	if err != nil {
		log.Print(err)
		log.Printf("\n")
	}
}

func (dm dbManager) getRecentDate(n int) {

}
