package engine

import "golang.org/x/crypto/bcrypt"

type ModelUser struct {
	Email          string
	HashedPassword string
	IsAdmin        bool

	// @todo Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// ForgotPasswordToken
	// ForgotPasswordSentAt
	// CreatedAt
}

// @todo FUTURE FEATURE need a session model for DB storage

func CreateUser(email, password string) error {
	// @todo make sure email is unique, make sure email is valid, make sure pass is at least 6 characters VALIDATION

	hashed, _ := HashPassword(password)
	db.C("users").Insert(&ModelUser{Email: email, HashedPassword: string(hashed), IsAdmin: false})

	return nil
}

func HashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func MatchPassword(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return true
	}

	return false
}
