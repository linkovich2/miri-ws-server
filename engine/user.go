package engine

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
	Account    *ModelUser
	Connection *Connection
	IsAdmin    bool
	State      int
	Characters []Character
}
