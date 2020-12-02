package main

import (
	"fmt"
)

type Data struct {
	Id          int     `json:"id"`
	Dht22_Humi  float32 `json:"dht22_Humi"`
	Dht22_Temp  float32 `json:"dht22_Temp"`
	Bmp180_Temp float32 `json:"bmp180_Temp"`
	Bmp180_Pres float32 `json:"bmp180_Pres"`
	Datetime    string  `json:"datetime"`
}

func main() {
	dm := initializeDB()

	data := dm.getRecentDate(1)
	fmt.Println(data)

	data = Data{-1, 1.1, 2.2, 3.3, 4.4, ""}
	dm.addData(data)
	//
	//data = dm.getRecentDate(1)
	//fmt.Println(data)
}
