package app

import (
	"encoding/json"
	"log"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type connectionHub struct {
	connections       map[*connection]bool // Registered connections.
	connectionHandler connectionHandler
	inbound           chan *message    // Inbound messages from the connections.
	register          chan *connection // Register requests from the connections.
	unregister        chan *connection // Unregister requests from connections.
}

var hub = connectionHub{
	connections:       make(map[*connection]bool),
	connectionHandler: &defaultConnectionHandler{},
	inbound:           make(chan *message),
	register:          make(chan *connection),
	unregister:        make(chan *connection),
}

func (h *connectionHub) Run() {
	if h.connectionHandler == nil {
		log.Fatal("Connection Handler must be set for websocket hub to function.")
		return
	}

	for {
		select {
		case c := <-h.register:
			h.connections[c] = true

			logger.Info("New Connection [%s]", c.id)
			h.connectionHandler.connect(c)

		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				// run any other logic on disconnect we need here
				logger.Notice("Connection [%s] disconnected", c.id)
				h.connectionHandler.disconnect(c)

				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			h.connectionHandler.handle(m)
		}
	}
}

func (h *connectionHub) setHandler(c connectionHandler) {
	h.connectionHandler = c
}

// Send a message to a lot of connections
func (h *connectionHub) broadcast(m *message, targets []*connection) {
	msg, _ := json.Marshal(m.payload)
	for _, c := range targets {
		c.send <- msg
	}
}

// Send a message to one connection
func (h *connectionHub) send(m *message) {
	msg, _ := json.Marshal(m.payload)
	m.connection.send <- msg
}
