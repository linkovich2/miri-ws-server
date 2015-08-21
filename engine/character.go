package engine

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

/*character should be something every User has*/
type (
	Character struct {
		/*Types that should be contained inside a character*/
		Race string
	}
)

func (c *Character) CommandAuthenticated_CHARSEL() {
	// Check the db for saved characters first!
	var ch []Character
	races, err := ioutil.ReadFile("races.json")
	if err != nil {
		log.Print("error: %v", err)
	}
	er := json.Unmarshal(races, ch)
	if er != nil {
		log.Print("error: %v", er)
	}
}
