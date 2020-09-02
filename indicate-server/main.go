package main

import (
        "strings"
        "net/http"
        "html/template"
       )

const dataFileName = "../collect-server/data.txt"

type Data struct {
    Dht22_Humi  float32 `json:"dht22_Humi"`
    Dht22_Temp  float32 `json:"dht22_Temp"`
    Bmp180_Temp float32 `json:"bmp180_Temp"`
    Bmp180_Pres float32 `json:"bmp180_Pres"`
    Date string
}

func dashboard(w http.ResponseWriter, r *http.Request) {
    lastLine := strings.TrimSpace(readLastLine(dataFileName))

    rawDate := lastLine[:19]
    rawData := lastLine[20:]

    data := parseJSON(rawData)
    data.Date = rawDate

    t, _ := template.ParseFiles("dashboard.html")
    t.Execute(w, data)
}

func main() {
    http.HandleFunc("/", dashboard)

    http.ListenAndServe(":80", nil)
}
