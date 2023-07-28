package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func receive(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
	}

	defer conn.Close()

	var count = 1
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Message %d", count)))
		if err != nil {
			log.Println("write:", err)
			break
		}

		count++

		time.Sleep(time.Second)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
	}

	defer conn.Close()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func getenv(key string, defaultValue string) string {
	if value := os.Getenv(key); value == "" {
		return defaultValue
	} else {
		return value
	}
}

func main() {
	addr := getenv("HOST", "127.0.0.1") + ":" + getenv("PORT", "8000")
	log.Println("addr:", addr)

	http.HandleFunc("/receive", receive)
	http.HandleFunc("/echo", echo)

	log.SetFlags(0)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}
