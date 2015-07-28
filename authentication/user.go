package authentication

import (
	"crypto/bcrypt"
)

type User struct {
	Email          string
	HashedPassword string
	ID             int

	// @todo: Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// ForgotPasswordToken
	// ForgotPasswordSentAt
	// CreatedAt
}

func CreateUser(email, password string) (User, error) {
  return &User{email, HashPassword(password)}
}

func ValidatePassword(pw string) error {

}

func ValidateEmail(email string) error {

}

func ValidateForgotPasswordToken(token string) {

}

func HashPassword(pw string) string {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func Match(pw string, hash string) bool {
	err := bcrypt.CompareHashAndPassword(hash, pw)
	if err != nil {
		return true
	}

	return false
}
