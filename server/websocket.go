package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"log"
	"net/http"
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
			log.Println(err)
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
			receiveData := cast.ToStringMapInt(obj["data"])

			data := dm.getRecentDate(receiveData["n"])

			b, _ := json.Marshal(data)
			err = client.WriteMessage(messageType, b)
			if err != nil {
				break
			}
		}
	}
}
