package main

// import (
// 	"bufio"
// 	"net"
// 	"net/http"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestHandleWebSocket(t *testing.T) {
// 	e := echo.New()
// 	e.GET("/ws", handleWebSocket)

// 	go e.Start(":8080")
// 	defer e.Close()

// 	time.Sleep(1 * time.Second)

// 	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws?room=room1", nil)
// 	if err != nil {
// 		t.Fatalf("failed to dial: %v", err)
// 	}
// 	defer conn.Close()

// 	// Send a message from the first client
// 	err = conn.WriteMessage(websocket.TextMessage, []byte("hello from client 1"))
// 	if err != nil {
// 		t.Fatalf("failed to write message: %v", err)
// 	}

// 	// Connect a second client to the same room
// 	conn2, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws?room=room1", nil)
// 	if err != nil {
// 		t.Fatalf("failed to dial: %v", err)
// 	}
// 	defer conn2.Close()

// 	// Send a message from the second client
// 	err = conn2.WriteMessage(websocket.TextMessage, []byte("hello from client 2"))
// 	if err != nil {
// 		t.Fatalf("failed to write message: %v", err)
// 	}

// 	// Read messages from the first client and check if they match the expected values
// 	expectedMessages := []string{"hello from client 2"}
// 	for _, expectedMessage := range expectedMessages {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			t.Fatalf("failed to read message: %v", err)
// 		}
// 		assert.Equal(t, expectedMessage, string(message))
// 	}

// 	// Read messages from the second client and check if they match the expected values
// 	expectedMessages = []string{"hello from client 1"}
// 	for _, expectedMessage := range expectedMessages {
// 		_, message, err := conn2.ReadMessage()
// 		if err != nil {
// 			t.Fatalf("failed to read message: %v", err)
// 		}
// 		assert.Equal(t, expectedMessage, string(message))
// 	}
// }

// func TestHandleWebSocketWithoutRoom(t *testing.T) {
// 	e := echo.New()
// 	e.GET("/ws", handleWebSocket)

// 	go e.Start(":8080")
// 	defer e.Close()

// 	time.Sleep(1 * time.Second)

// 	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
// 	if err != nil {
