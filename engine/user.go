package engine

const (
	notAuthenticated = iota
	authenticated
	inGame
)

func stateString(s int) string {
	switch s {
	case 0:
		return "NotAuthenticated"
	case 1:
		return "Authenticated"
	default:
		return ""
	}
}

type user struct {
	account    *modelUser
	connection *connection
	isAdmin    bool
	state      int
}
