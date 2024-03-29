package server

import (
	"github.com/gorilla/websocket"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"net/http"
	"time"
)

const (
	WriteWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	PongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	PingPeriod     = (PongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	MaxMessageSize = 512                 // Maximum message size allowed from peer.
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // @todo maybe make this a bit more strict?
	},
}

// connection is an middleman between the websocket connection and the hub.
type Connection struct {
	webSocket *websocket.Conn // The websocket connection.
	send      chan []byte     // Buffered channel of outbound messages.
	ID        string
	User      *core.User
}

type InboundMessage struct {
	Payload    []byte
	Connection *Connection
}

// ReadPump pumps messages from the websocket connection to the hub.
func (c *Connection) readPump() {
	defer func() {
		hub.unregister <- c
		c.webSocket.Close()
	}()
	c.webSocket.SetReadLimit(MaxMessageSize)
	c.webSocket.SetReadDeadline(time.Now().Add(PongWait))
	c.webSocket.SetPongHandler(func(string) error { c.webSocket.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, msg, err := c.webSocket.ReadMessage()

		if err != nil {
			break
		}
		hub.inbound <- &InboundMessage{msg, c}
	}
}

// Write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.webSocket.SetWriteDeadline(time.Now().Add(WriteWait))
	return c.webSocket.WriteMessage(mt, payload)
}

// WritePump pumps messages from the hub to the websocket connection.
func (c *Connection) writePump() {
	ticker := time.NewTicker(PingPeriod)
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

func (c *Connection) Send(m []byte) {
	defer func() {
		if r := recover(); r != nil {
			logger.Write.Error("Connection [%s] send channel is closed!", c.ID)
		}
	}()

	c.send <- m
}
