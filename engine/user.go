package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/models"
)

const (
	NotAuthenticated = iota
	Authenticated
	InGame
)

func StateString(s int) string {
	switch s {
	case NotAuthenticated:
		return "NotAuthenticated"
	case Authenticated:
		return "Authenticated"
	default:
		return ""
	}
}

type User struct {
	Account    *models.User
	Connection *Connection
	State      int
	Character  *Character
}
