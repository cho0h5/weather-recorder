package main

import (
        "net/http"
        "html/template"
       )

type Data struct {
    Dht22_Humi  float32 `json:"dht22_Humi"`
    Dht22_Temp  float32 `json:"dht22_Temp"`
    Bmp180_Temp float32 `json:"bmp180_Temp"`
    Bmp180_Pres float32 `json:"bmp180_Pres"`
}

func dashboard(w http.ResponseWriter, r *http.Request) {
    data := Data {66.60, 31.10, 31.80, 101180.00}

    t, _ := template.ParseFiles("dashboard.html")
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", dashboard)

    http.ListenAndServe(":80", nil)
}
