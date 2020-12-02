package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	// "encoding/json"
)

func input(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))

	log.Printf("%s\n", string(body))
	addData(string(body))

	/*
	   var data Data
	   json.Unmarshal(body, &data)

	   body2 := `{"dht22_Humi":"29.9"}`
	   json.Unmarshal([]byte(body2), &data)
	   log.Printf("%+v\n", data)
	*/
}

func addData(data string) {
	file, _ := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()

	fmt.Fprintln(file, time.Now().Format("2006-01-02 15:04:05"), data)
}

func main() {
	http.HandleFunc("/input", input)

	http.ListenAndServe(":8080", nil)
}
