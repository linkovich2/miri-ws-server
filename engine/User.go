package engine

const (
	NotAuthenticated = iota
	Authenticated
	InGame
)

type (
	LoginData struct {
		Email    string
		Password string
	}

	User struct {
		Account    *UserModel
		Connection *Connection
		IsAdmin    bool
		State      int
	}
)
