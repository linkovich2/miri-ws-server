package auth

import (
  "golang.org/x/crypto/bcrypt"
)

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

func CreateUser(email, password string) *UserModel {
  hashed, _ := hashPassword(password)
  return &UserModel{Email: email, HashedPassword: string(hashed)}
}

// func ValidatePassword(pw string) error {
//
// }
//
// func ValidateEmail(email string) error {
//
// }
//
// func ValidateForgotPasswordToken(token string) {
//
// }

func hashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func Match(pw string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err != nil {
		return true
	}

	return false
}
