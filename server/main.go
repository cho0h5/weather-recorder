package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var dm dbManager

func input(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	fmt.Fprintln(w, string(body))

	var data Data
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}
	log.Println(data)
	dm.addData(data)
}

func main() {
	dm = initializeDB()

	http.HandleFunc("/ws", webSocket)
	http.HandleFunc("/input", input)
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}
