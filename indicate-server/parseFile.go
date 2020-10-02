package main

import (
        "os"
        "time"
        "strings"
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
    loc, _ := time.LoadLocation("Asia/Seoul")

    prevDate, _ := time.ParseInLocation("2006-01-02 15:04:05", rawDate, loc)
    nowDate := time.Now()
    term := nowDate.Sub(prevDate)

    state := term < time.Second * 3

    return state
}

func parseData() Data {
    lastLine := strings.TrimSpace(readLastLine(dataFileName))

    rawDate := lastLine[:19]
    rawData := lastLine[20:]

    data := parseJSON(rawData)
    data.Date = rawDate
    data.IsWorking = isWorking(rawDate)

    return data
}
