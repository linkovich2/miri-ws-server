package engine

import "golang.org/x/crypto/bcrypt"

type UserModel struct {
	Email          string
	HashedPassword string

	// @todo: Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// ForgotPasswordToken
	// ForgotPasswordSentAt
	// CreatedAt
}

type SessionModel struct {
	SessionID string
	UserID    string
}

func CreateUser(email, password string) error {
	hashed, _ := hashPassword(password)
	db.C("users").Insert(&UserModel{Email: email, HashedPassword: string(hashed)})

	return nil
}

func hashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func Match(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return true
	}

	return false
}
