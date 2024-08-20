package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	Clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	rooms  = make(map[string]*Room)
	roomMu sync.Mutex
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Signaling server starting on :28080")
	log.Fatal(http.ListenAndServe(":28080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	log.Println("New WebSocket connection from:", remoteAddr)

	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		roomID = fmt.Sprintf("room_%d", len(rooms)+1)
		log.Printf("Created new room: %s\n", roomID)
	}

	roomMu.Lock()
	room, exists := rooms[roomID]
	if !exists {
		room = &Room{Clients: make(map[*websocket.Conn]bool)}
		rooms[roomID] = room
	}
	roomMu.Unlock()

	room.mu.Lock()
	room.Clients[conn] = true
	room.mu.Unlock()

	log.Printf("Client[%v] joined room %s\n", remoteAddr, roomID)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		var msg map[string]interface{}
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		msg["roomId"] = roomID
		updatedMessage, _ := json.Marshal(msg)

		room.mu.Lock()
		for client := range room.Clients {
			if client != conn {
				clientAddr := client.RemoteAddr().String()
				if err := client.WriteMessage(messageType, updatedMessage); err != nil {
					log.Println("Error writing message:", err)
				} else {
					log.Printf("writing message to client[%v] ok\n", clientAddr)
				}
			}
		}
		room.mu.Unlock()
	}

	room.mu.Lock()
	delete(room.Clients, conn)
	room.mu.Unlock()
	log.Printf("Client[%v] left room %s\n", remoteAddr, roomID)
}
