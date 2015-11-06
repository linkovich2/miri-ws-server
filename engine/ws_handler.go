package engine

type (
	connectionHandler interface {
		connect(c *connection) bool
		disconnect(c *connection) bool
		handle(m *message) bool
	}
)

type defaultConnectionHandler struct{}

func (d *defaultConnectionHandler) connect(c *connection) bool {
	return false
}

func (d *defaultConnectionHandler) disconnect(c *connection) bool {
	return false
}

func (d *defaultConnectionHandler) handle(m *message) bool {
	return false
}
