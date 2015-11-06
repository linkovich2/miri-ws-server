package server

type (
	ConnectionHandler interface {
		Connect(c *Connection) bool
		Disconnect(c *Connection) bool
		Handle(m *InboundMessage) bool
	}
)

type defaultConnectionHandler struct{}

func (d *defaultConnectionHandler) Connect(c *Connection) bool {
	return false
}

func (d *defaultConnectionHandler) Disconnect(c *Connection) bool {
	return false
}

func (d *defaultConnectionHandler) Handle(m *InboundMessage) bool {
	return false
}
