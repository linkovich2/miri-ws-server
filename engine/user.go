package engine

const (
	notAuthenticated = iota
	authenticated
	inGame
)

func stateString(s int) string {
	switch s {
	case notAuthenticated:
		return "NotAuthenticated"
	case authenticated:
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
