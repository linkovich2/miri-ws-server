package engine

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
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
		Race          *Race
	}
	Race struct {
		Description string
	}
)

func (c *Character) Delete() {
}

func (c *Character) CreateCharacter() []Character {
	var ch []Character
	races, err := ioutil.ReadFile("races.json")
	if err != nil {
		log.Print("error: %v", err)
	}
	er := json.Unmarshal(races, &ch)
	if er != nil {
		log.Print("error: %v", er)
	}
	log.Print(ch[1].Race)
	return nil
}

func (h *HandlerInterface) CommandAuthenticated_CHARSEL(u *User) {
	// Check the db for saved characters first!
	var Finduser mgo.Collection
	q := Finduser.FindId(u.Account)
	log.Print(q)
	// Once we're done change state to InGame? Or return to character select?
}
