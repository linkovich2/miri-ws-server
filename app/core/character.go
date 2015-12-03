package core

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Character struct {
	ID               bson.ObjectId       `bson:"_id,omitempty" json:"id"`
	Race             string              `json:"race"`
	Gender           string              `json:"gender"`
	AestheticTraits  map[string][]string `json:"aesthetic_traits"`
	FunctionalTraits map[string][]string `json:"functional_traits"`
	Background       string              `json:"background"`
	Name             string              `json:"name"`
	UserID           bson.ObjectId       `json:"-" bson:"user_id"`
	Created          time.Time           `json:"created"`
}
