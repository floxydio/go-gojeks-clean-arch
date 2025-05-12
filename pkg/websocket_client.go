package pkg

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	userSockets   = make(map[string]*websocket.Conn)
	driverSockets = make(map[string]*websocket.Conn)
	mu            sync.Mutex
)

func SaveUserSocket(userID string, conn *websocket.Conn) {
	mu.Lock()
	userSockets[userID] = conn
	mu.Unlock()
	log.Println("User WebSocket saved:", userID)
}

func RemoveUserSocket(userID string) {
	mu.Lock()
	delete(userSockets, userID)
	mu.Unlock()
}
func SaveDriverSocket(userID string, conn *websocket.Conn) {
	mu.Lock()
	driverSockets[userID] = conn
	mu.Unlock()
	log.Println("Driver WebSocket saved:", userID)
}

func RemoveDriverSocket(userID string) {
	mu.Lock()
	delete(userSockets, userID)
	mu.Unlock()
}

func SendToUser(userID string, message string) {
	mu.Lock()
	defer mu.Unlock()

	log.Println("Trying to send to userID:", userID)

	conn, ok := userSockets[userID]
	if !ok {
		log.Printf("User %s not connected via WebSocket\n", userID)
		return
	}

	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Printf("Failed to send message to user %s: %v\n", userID, err)
		conn.Close()
		delete(userSockets, userID)
	}
}

func SendDriver(message string) {
	mu.Lock()
	defer mu.Unlock()

	for driverId, conn := range driverSockets {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Failed to send message to Driver %s: %v\n", driverId, err)
			conn.Close()
			delete(driverSockets, driverId)
		}
	}
}
