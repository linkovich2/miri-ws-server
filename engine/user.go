package engine

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type user struct {
	id                  bson.ObjectId `bson:"_id,omitempty"`
	email               string
	password            string
	role                string
	salt                string
	provider            string
	resetPasswordToken  string
	resetPasswordSent   time.Time
	lastLogin           time.Time
	created             time.Time
	updated             time.Time
	failedLoginAttempts int
}

func (u *user) isAdmin() bool {
	return bool(u.role == "admin")
}
