package main

import (
        "net/http"
        "html/template"
        "encoding/json"
       )

const dataFileName = "../collect-server/data.txt"

type Data struct {
    Dht22_Humi  float32 `json:"dht22_Humi"`
    Dht22_Temp  float32 `json:"dht22_Temp"`
    Bmp180_Temp float32 `json:"bmp180_Temp"`
    Bmp180_Pres float32 `json:"bmp180_Pres"`
    Date        string  `json:"date"`
    IsWorking   bool    `json:"isWorking"`
}

func dashboard(w http.ResponseWriter, r *http.Request) {
    data := parseData()

    t, _ := template.ParseFiles("dashboard.html")
    t.Execute(w, data)
}

func json_data(w http.ResponseWriter, r *http.Request) {
    data := parseData()
    enc := json.NewEncoder(w)
    w.Header().Set("Content-Type", "application/json")
    enc.Encode(data)
}

func main() {
    http.HandleFunc("/", dashboard)
    http.HandleFunc("/json-data", json_data)

    http.ListenAndServe(":80", nil)
}
