package websocket

import "fmt"


// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	Connections map[*Connection]bool

	// Inbound messages from the connections.
	Inbound chan *Message

	// Register requests from the connections.
	Register chan *Connection

	// Unregister requests from connections.
	Unregister chan *Connection

	OnConnect func(c *Connection)
	OnMessage func(msg *Message)
	OnDisconnect func(c *Connection)
}

var h = Hub{
	Inbound:   make(chan *Message),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	Connections: make(map[*Connection]bool),
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c] = true
			h.OnConnect(c)
		case c := <-h.Unregister:
			if _, ok := h.Connections[c]; ok {
				h.OnDisconnect(c)
				delete(h.Connections, c)
				close(c.Send)
			}
		case m := <-h.Inbound:
			// handle message
			fmt.Println(string(m.Payload))
			h.OnMessage(m)
		}
	}
}

// Get's called whenever there is a new connection
func (h *Hub) SetOnConnectCallback(callback func(c *Connection)) {
	h.OnConnect = callback
}

// Get's called whenever a message is received from a connection
func (h *Hub) SetOnMessageCallback(callback func(msg *Message)) {
	h.OnMessage = callback
}

// Get's called when there is a disconnection
func (h *Hub) SetOnDisconnectCallback(callback func(c *Connection)) {
	h.OnDisconnect = callback
}

// Send a message to a lot of connections
func (h *Hub) Broadcast(msg []byte, targets []*Connection) {
	for c := range targets {
		c.Send <- msg
	}
}

// Send a message to one connection
func (h *Hub) Send(msg []byte, c *Connection) {
  c.Send <- msg
}
