package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"

	"bytes"
)

func DescribeCharacter(c *core.Character) string {
	response := bytes.NewBuffer([]byte{})
	race := content.Races[c.Race]
	gender := content.Genders[c.Gender]

	if race.Descriptor[:1] == "a" || race.Descriptor[:1] == "e" || race.Descriptor[:1] == "i" || race.Descriptor[:1] == "o" || race.Descriptor[:1] == "u" {
		response.Write([]byte("An "))
	} else {
		response.Write([]byte("A "))
	}

	response.Write([]byte(race.Descriptor))
	response.Write([]byte(" "))

	if race.GenderHuman {
		response.Write([]byte(gender.Human))
	} else {
		response.Write([]byte(gender.Scientific))
	}

	// @todo build an aesthetic traits based descriptions
	// @todo future: need at least one more description based on deeds, skills and functional traits

	return response.String()
}
