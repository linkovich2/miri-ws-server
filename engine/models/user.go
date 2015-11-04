package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	Email          string
	HashedPassword string
	IsAdmin        bool
	Characters     [3]Character

	// @todo Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// ForgotPasswordToken
	// ForgotPasswordSentAt
	// CreatedAt
}
