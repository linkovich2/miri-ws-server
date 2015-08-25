package engine

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
)

const (
	M = iota
	F
)

/*character should be something every User has*/
type (
	Character struct {
		/*Types that should be contained inside a character*/
		CharacterName string
		Race          []Race
	}
	Race struct {
		RaceName    string
		Gender      string
		Description string
	}
)

func (c *Character) Delete() {
}

func (c *Character) CharacterList(u *User) []Character {
	var ch Character
	races, err := ioutil.ReadFile("races.json")
	if err != nil {
		log.Print("error: %v", err)
	}
	er := json.Unmarshal(races, &ch.Race)
	if er != nil {
		log.Print("error: %v", er)
	}
	log.Print(ch.Race)
	return nil
}

func (h *HandlerInterface) CommandAuthenticated_CHARLIST(u *User, msg *json.RawMessage) {
	// Check the db for saved characters first!
	log.Print("CommandAuthenticated_CHARLIST is called", db.C("users").Find(bson.M{"email": u.Account.Email}))
	// Once we're done change state to InGame? Or return to character select?
}
