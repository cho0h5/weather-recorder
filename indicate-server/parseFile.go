package main

import (
        "os"
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
