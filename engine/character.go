package engine

import (
	"encoding/json"
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
		Description string
	}
)

func (c *Character) Delete() {
}

func (c *Character) CharacterList() []Character {
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

func (h *HandlerInterface) CommandAuthenticated_CHARSEL(u *User, msg *json.RawMessage) {
	// Check the db for saved characters first!
	logger.Info("CommandAuthenticated_CHARSEL is called")
	// Once we're done change state to InGame? Or return to character select?
}
