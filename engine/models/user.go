package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID                   bson.ObjectId `bson:"_id,omitempty"`
	Email                string
	HashedPassword       string
	IsAdmin              bool
	Characters           [3]Character
	ForgotPasswordToken  string
	ForgotPasswordSentAt time.Time

	// @todo Future stuff
	// LastLoginDate
	// LastLoginIP
	// CurrentLoginDate
	// CurrentLoginIP
	// CreatedAt
}
