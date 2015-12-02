package core

import (
	// "encoding/json"
	// "github.com/jonathonharrell/miri-ws-server/app/util"
	// "strconv"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Character struct {
	ID               bson.ObjectId       `bson:"_id,omitempty" json:"-"`
	Race             string              `json:"race"`
	Gender           string              `json:"gender"`
	AestheticTraits  map[string][]string `json:"aesthetic_traits"`
	FunctionalTraits map[string][]string `json:"functional_traits"`
	Background       string              `json:"background"`
	Name             string              `json:"name"`
	UserID           string              `json:"-" bson:"user_id"`
	Created          time.Time           `json:"created"`
}

// func (f *CharacterForm) validateFunctionalTraits(c *Character) (valid bool, errors []string) {
// 	validTraits := f.getAvailableFunctionalTraits()
// 	points := 0
// 	list := []string{}
//
// 	for _, ftc := range validTraits {
// 		for _, ft := range ftc.Traits {
// 			list = append(list, ft.ID)
//
// 			if in, _ := util.InArray(ft.ID, c.FunctionalTraits); !in {
// 				if ft.Required { // it's not in but it's required
// 					return false, []string{"'" + ft.Name + "' is a non-optional trait."}
// 				}
// 			} else {
// 				// it is in, validate anything else against it
// 				val, _ := strconv.Atoi(ft.Points)
// 				points = points + val
//
// 				// @todo need to validate for atleastonerequired, and any other constraints
// 			}
// 		}
// 	}
//
// 	if points < 0 {
// 		return false, []string{"You must have at least 0 points to continue."}
// 	}
//
// 	for _, trait := range c.FunctionalTraits {
// 		if in, _ := util.InArray(trait, list); !in {
// 			return false, []string{"Illegal trait included."}
// 		}
// 	}
//
// 	return true, []string{}
// }
//
// func (f *CharacterForm) getAvailableBackgrounds() []BackgroundShort {
// 	res := []BackgroundShort{}
//
// 	for _, bg := range backgrounds { // Category Level
// 		if bg.AllowAll {
// 			res = append(res, bg.Shorten())
// 			continue
// 		}
//
// 		// check against prerequisites to see if allowed @todo
// 	}
//
// 	return res
// }
