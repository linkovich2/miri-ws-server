package server

import (
	"log"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type ConnectionHub struct {
	connections       map[*Connection]bool // Registered connections.
	connectionHandler ConnectionHandler
	inbound           chan *InboundMessage // Inbound messages from the connections.
	register          chan *Connection // Register requests from the connections.
	unregister        chan *Connection // Unregister requests from connections.
}

var hub = ConnectionHub{
	connections:       make(map[*Connection]bool),
	connectionHandler: &defaultConnectionHandler{},
	inbound:           make(chan *InboundMessage),
	register:          make(chan *Connection),
	unregister:        make(chan *Connection),
}

func GetHub() ConnectionHub {
	return hub
}

func (h *ConnectionHub) Run() {
	if h.connectionHandler == nil {
		log.Fatal("Connection Handler must be set for websocket hub to function.")
		return
	}

	for {
		select {
		case c := <-h.register:
			h.connections[c] = true

			logger.Write.Info("New Connection [%s]", c.ID)
			h.connectionHandler.Connect(c)

		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				// run any other logic on disconnect we need here
				logger.Write.Notice("Connection [%s] disconnected", c.ID)
				h.connectionHandler.Disconnect(c)

				delete(h.connections, c)
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
func (h *ConnectionHub) Broadcast(m []byte, targets []*Connection) {
	for _, c := range targets {
		c.send <- m
	}
}

// Send a message to one connection
func (h *ConnectionHub) Send(m []byte, c *Connection) {
	c.send <- m
}
