package engine

const (
	notAuthenticated = iota
	authenticated
	inGame
)

type (
	formLogin struct {
		email    string
		password string
	}

	user struct {
		account    *modelUser
		connection *connection
		isAdmin    bool
		state      int
	}
)
