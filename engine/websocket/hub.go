package websocket

import "fmt"

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	connections  map[*Connection]bool // Registered connections.
	inbound      chan *Message        // Inbound messages from the connections.
	register     chan *Connection     // Register requests from the connections.
	unregister   chan *Connection     // Unregister requests from connections.
	onConnect    func(c *Connection)
	onMessage    func(msg *Message)
	onDisconnect func(c *Connection)
}

var h = Hub{
	connections:  make(map[*Connection]bool),
	inbound:      make(chan *Message),
	register:     make(chan *Connection),
	unregister:   make(chan *Connection),
	onConnect:    func(c *Connection) {},
	onMessage:    func(m *Message) {},
	onDisconnect: func(c *Connection) {},
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			h.onConnect(c)
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				h.onDisconnect(c)
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			// handle message
			fmt.Println(string(m.Payload))
			h.onMessage(m)
		}
	}
}

// Get's called whenever there is a new connection
func (h *Hub) SetOnConnectCallback(callback func(c *Connection)) {
	h.onConnect = callback
}

// Get's called whenever a message is received from a connection
func (h *Hub) SetOnMessageCallback(callback func(msg *Message)) {
	h.onMessage = callback
}

// Get's called when there is a disconnection
func (h *Hub) SetOnDisconnectCallback(callback func(c *Connection)) {
	h.onDisconnect = callback
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
