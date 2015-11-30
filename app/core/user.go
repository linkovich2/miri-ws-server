package core

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	Email               string
	Password            string
	Role                string
	Salt                string
	Provider            string
	ResetPasswordToken  string
	ResetPasswordSent   time.Time
	LastLogin           time.Time
	Created             time.Time
	Updated             time.Time
	FailedLoginAttempts int

	Characters []Character
}

func (u *User) IsAdmin() bool {
	return bool(u.Role == "admin")
}
