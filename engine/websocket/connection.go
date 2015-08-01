package websocket

import (
	"github.com/gorilla/websocket"
	"stablelib.com/v1/uniuri"
	"log"
	"net/http"
	"time"
	"strings"
	"fmt"
)

const (
	writeWait = 10 * time.Second     // Time allowed to write a message to the peer.
	pongWait = 60 * time.Second      // Time allowed to read the next pong message from the peer.
	pingPeriod = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512             // Maximum message size allowed from peer.
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
	webSocket *websocket.Conn // The websocket connection.
	send chan []byte // Buffered channel of outbound messages.
	ID   string
}

type Message struct {
	Payload []byte
	Connection *Connection
}

// ReadPump pumps messages from the websocket connection to the hub.
func (c *Connection) readPump() {
	defer func() {
		h.unregister <- c
		c.webSocket.Close()
	}()
	c.webSocket.SetReadLimit(maxMessageSize)
	c.webSocket.SetReadDeadline(time.Now().Add(pongWait))
	c.webSocket.SetPongHandler(func(string) error { c.webSocket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.webSocket.ReadMessage()

		if err != nil {
			break
		}
		h.inbound <- &Message{message, c}
	}
}

// Write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.webSocket.SetWriteDeadline(time.Now().Add(writeWait))
	return c.webSocket.WriteMessage(mt, payload)
}

// WritePump pumps messages from the hub to the websocket connection.
func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.webSocket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New()}
	fmt.Println(c.ID)
	h.register <- c
	go c.writePump()
	c.readPump()
}