package main

import "database/sql"

type Data struct {
	Id          int     `json:"id"`
	Dht22_Humi  float32 `json:"dht22_Humi"`
	Dht22_Temp  float32 `json:"dht22_Temp"`
	Bmp180_Temp float32 `json:"bmp180_Temp"`
	Bmp180_Pres float32 `json:"bmp180_Pres"`
	Datetime    string  `json:"datetime"`
}

type dbManager struct {
	db *sql.DB
}
