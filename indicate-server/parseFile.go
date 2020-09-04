package main

import (
        "os"
        "time"
        "encoding/json"
)

func readLastLine(fileName string) string {
    file, _ := os.Open(fileName)
    defer file.Close()

    buf := make([]byte, 104)
    stat, _ := os.Stat(fileName)
    start := stat.Size() - 104
    file.ReadAt(buf, start)

    return string(buf)
}

func parseJSON(rawData string) Data {
    var data Data
    json.Unmarshal([]byte(rawData), &data)

    return data
}

func isWorking(rawDate string) bool {
    prevDate, _ := time.Parse("2006-01-02 15:04:05", rawDate)
    nowDate := time.Now()
    term := nowDate.Sub(prevDate)

    state := term < time.Second * 3

    return state
}
