package engine

const (
	notAuthenticated = iota
	authenticated
	inGame
)

type user struct {
	account    *modelUser
	connection *connection
	isAdmin    bool
	state      int
}
