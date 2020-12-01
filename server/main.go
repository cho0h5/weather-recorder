package server

type Data struct {
	Dht22_Humi  float32 `json:"dht22_Humi"`
	Dht22_Temp  float32 `json:"dht22_Temp"`
	Bmp180_Temp float32 `json:"bmp180_Temp"`
	Bmp180_Pres float32 `json:"bmp180_Pres"`
}

func main() {

}
