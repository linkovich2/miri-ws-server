package auth

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/jonathonharrell/miri-ws-server/engine/database"
)

type UserModel struct {
	Email          []byte
	HashedPassword []byte

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
	SessionID []byte
	UserID    []byte
}

func CreateUser(email, password []byte) error {
	hashed, _ := hashPassword(password)
	database.DB.C("users").Insert(&UserModel{Email: email, HashedPassword: hashed})

	return nil
}

func hashPassword(pw []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pw, 10)
}

func Match(pw, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		return true
	}

	return false
}
