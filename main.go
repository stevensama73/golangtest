package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var clients = make(map[*client]struct{})

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &client{conn: conn}

	clients[client] = struct{}{}

	defer func() {
		client.mu.Lock()
		delete(clients, client)
		client.mu.Unlock()
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		broadcast(p, messageType)
	}
}

func broadcast(message []byte, messageType int) {
	for client := range clients {
		client.mu.Lock()
		err := client.conn.WriteMessage(messageType, message)
		client.mu.Unlock()

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	port := 5000
	fmt.Printf("Server is listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
