package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func webSocket(w http.ResponseWriter, r *http.Request) {
	// ready websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	client, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// receive and send
	for {
		// receive
		messageType, message, err := client.ReadMessage()
		if err != nil {
			log.Println("somwone disconnected")
			break
		}

		var obj map[string]interface{}
		err = json.Unmarshal(message, &obj)
		if err != nil {
			log.Println(err)
		}

		switch obj["event"] {
		case "enter":
			log.Println("someone connected")

		case "recentData":
			log.Println("request recent Data", obj["n"])

			data := dm.getRecentDate(int(obj["n"].(float64)))

			b, _ := json.Marshal(data)
			err = client.WriteMessage(messageType, b)
			if err != nil {
				break
			}
		}
	}
}
