package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Handler untuk WebSocket
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		for {
			// Generate random values for water and wind
			water := rand.Intn(100) + 1
			wind := rand.Intn(100) + 1

			// Create status object
			status := Status{Water: water, Wind: wind}

			// Encode status to JSON
			statusJSON, err := json.Marshal(status)
			if err != nil {
				log.Println(err)
				return
			}

			// Write status JSON to WebSocket connection
			if err := conn.WriteMessage(websocket.TextMessage, statusJSON); err != nil {
				log.Println(err)
				return
			}

			// Wait for 15 seconds before next update
			time.Sleep(15 * time.Second)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
