package engine

import ()

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	connections map[*Connection]bool // Registered connections.
	inbound     chan *Message        // Inbound messages from the connections.
	register    chan *Connection     // Register requests from the connections.
	unregister  chan *Connection     // Unregister requests from connections.
}

var hub = Hub{
	connections: make(map[*Connection]bool),
	inbound:     make(chan *Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			users[c.ID] = &User{Connection: c, State: NotAuthenticated}
			// maybe we should also try to authenticate, if we want to use cookies or whatever
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				// run any other logic on disconnect we need here
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			Interpret(m, users[m.Connection.ID])
		}
	}
}

// Send a message to a lot of connections
func (h *Hub) Broadcast(msg []byte, targets []*Connection) {
	for _, c := range targets {
		c.send <- msg
	}
}

// Send a message to one connection
func (h *Hub) Send(msg []byte, c *Connection) {
	c.send <- msg
}
