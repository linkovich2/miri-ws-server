package websockets

import (
	"encoding/json"
  "log"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type ConnectionHub struct {
	Connections       map[*Connection]bool // Registered connections.
	connectionHandler ConnectionHandler
	inbound           chan *Message    // Inbound messages from the connections.
	register          chan *Connection // Register requests from the connections.
	unregister        chan *Connection // Unregister requests from connections.
}

var Hub = ConnectionHub{
	Connections: make(map[*Connection]bool),
  connectionHandler: &DefaultConnectionHandler{},
	inbound:     make(chan *Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
}

func (h *ConnectionHub) Run() {
  if h.connectionHandler == nil {
    log.Fatal("Connection Handler must be set for websocket hub to function.")
    return
  }

	for {
		select {
		case c := <-h.register:
			h.Connections[c] = true

			logger.Write.Info("New Connection [%s]", c.ID)
			h.connectionHandler.Connect(c)

		case c := <-h.unregister:
			if _, ok := h.Connections[c]; ok {
				// run any other logic on disconnect we need here
				logger.Write.Notice("Connection [%s] disconnected", c.ID)
        h.connectionHandler.Disconnect(c)

				delete(h.Connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			h.connectionHandler.Handle(m)
		}
	}
}

func (h *ConnectionHub) SetHandler(c ConnectionHandler) {
	h.connectionHandler = c
}

// Send a message to a lot of connections
func (h *ConnectionHub) Broadcast(m *Message, targets []*Connection) {
	msg, _ := json.Marshal(m)
	for _, c := range targets {
		c.send <- msg
	}
}

// Send a message to one connection
func (h *ConnectionHub) Send(m *Message, c *Connection) {
	msg, _ := json.Marshal(m)
	c.send <- msg
}
