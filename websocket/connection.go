package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"strings"
)

const (
	WRITE_WAIT = 10 * time.Second      // Time allowed to write a message to the peer.
	PONG_WAIT = 60 * time.Second       // Time allowed to read the next pong message from the peer.
	PING_PERIOD = (PONG_WAIT * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	MAX_MESSAGE_SIZE = 512             // Maximum message size allowed from peer.
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// I'm betting we'll need to remove this at some point so that we only accept connections
	// from specific cross-origins
	// We don't want users connecting from other locations with (potentially) malintent
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
	WebSocket *websocket.Conn // The websocket connection.
	Send chan []byte // Buffered channel of outbound messages.
}

type Message struct {
	Payload []byte
	Connection *Connection
}

// ReadPump pumps messages from the websocket connection to the hub.
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

		if err != nil {
			break
		}
		h.Inbound <- &Message{message, c}
	}
}

// Write writes a message with the given message type and payload.
func (c *Connection) Write(mt int, payload []byte) error {
	c.WebSocket.SetWriteDeadline(time.Now().Add(WRITE_WAIT))
	return c.WebSocket.WriteMessage(mt, payload)
}

// WritePump pumps messages from the hub to the websocket connection.
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
