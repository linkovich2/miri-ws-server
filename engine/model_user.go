package engine

import "golang.org/x/crypto/bcrypt"

type modelUser struct {
	email          string
	hashedPassword string

	// @todo: Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// ForgotPasswordToken
	// ForgotPasswordSentAt
	// CreatedAt
}

type modelSession struct {
	sessionId string
	userId    string
}

func createUser(email, password string) error {
	hashed, _ := hashPassword(password)
	db.C("users").Insert(&modelUser{email: email, hashedPassword: string(hashed)})

	return nil
}

func hashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func matchPassword(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return true
	}

	return false
}
