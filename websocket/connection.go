package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"strings"
	"fmt"
)

const (
	// Time allowed to write a message to the peer.
	WRITE_WAIT = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	PONG_WAIT = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	PING_PERIOD = (PONG_WAIT * 9) / 10

	// Maximum message size allowed from peer.
	MAX_MESSAGE_SIZE = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// I'm betting we'll need to remove this at some point so that we only accept connections
	// from specific cross-origins
	// We don't want users connecting from other locations with (potentially) mal-intent
	CheckOrigin: func(r *http.Request) bool {
		if strings.Contains(r.Header.Get("Origin"), "9000") {
			return true
		} else {
			return false
		}
	},
}

// connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	WebSocket *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (c *Connection) ReadPump() {
	defer func() {
		h.Unregister <- c
		c.WebSocket.Close()
	}()
	c.WebSocket.SetReadLimit(MAX_MESSAGE_SIZE)
	c.WebSocket.SetReadDeadline(time.Now().Add(PONG_WAIT))
	c.WebSocket.SetPongHandler(func(string) error { c.WebSocket.SetReadDeadline(time.Now().Add(PONG_WAIT)); return nil })
	for {
		_, message, err := c.WebSocket.ReadMessage()

		fmt.Println(string(message))
		// this is where the incoming messages are passed to the hub. In our case, when a socket is connected it should be given a handler interface
		// this handler should be able to pull from its incoming messages outside of the websocket class

		if err != nil {
			break
		}
		h.Broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.WebSocket.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
	return c.WebSocket.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Connection) WritePump() {
	ticker := time.NewTicker(PING_PERIOD)
	defer func() {
		ticker.Stop()
		c.WebSocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.Write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serverWs handles websocket requests from the peer.
func ServeWS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &Connection{Send: make(chan []byte, 256), WebSocket: ws}
	h.Register <- c
	go c.WritePump()
	c.ReadPump()
}
