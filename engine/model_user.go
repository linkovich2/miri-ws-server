package engine

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type ModelUser struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
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

func CreateUser(email, password string, u *User) (errors []string) {
	existing := ModelUser{}
	err := db.C("users").Find(bson.M{"email": email}).One(&existing)

	if err == nil { // checking for existing user
		errors = append(errors, "A user already exists with that email.")
	}

	if !govalidator.IsEmail(email) {
		errors = append(errors, "Not a valid email")
	}

	if len(password) < 6 {
		errors = append(errors, "Password must be at least 6 characters long.")
	}

	if len(errors) <= 0 {
		hashed, _ := HashPassword(password)
		user := &ModelUser{Email: email, HashedPassword: string(hashed), IsAdmin: false}
		db.C("users").Insert(user)

		u.State = Authenticated
		u.Account = user
		logger.Info("New User Created: %s", email)
	}

	return errors
}

func Authenticate(email, password string, u *User) (success bool, errors []string) {
	existing := ModelUser{}
	err := db.C("users").Find(bson.M{"email": email}).One(&existing)

	if err != nil { // no existing user
		errors = append(errors, "Invalid email or password.")
		return false, errors
	}

	success = MatchPassword(password, existing.HashedPassword)

	if err == nil && !success { // if user found but not matching password
		errors = append(errors, "Invalid email or password.")
	} else {
		u.State = Authenticated
		u.Account = &existing
		logger.Info("User logged in: %s", email)
	}

	return success, errors
}

func HashPassword(pw string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 10)
}

func MatchPassword(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	if err == nil {
		return true
	}

	return false
}
