package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util/filters"

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

	response.Write([]byte("; "))

	for i, cat := range c.AestheticTraits {
		for _, t := range cat {
			response.Write([]byte(filters.GenderPronouns(content.AestheticTraits[i].Traits[t].Description, gender.Possessive, gender.Pronoun, false)))
			response.Write([]byte(" "))
		}
	}

	return response.String()
}

// @todo future: need at least one more description based on deeds, skills and functional traits

// @todo this should log an error if it fails
func SaveCharacter(c *core.Character) {
	logger.Write.Info("Save Character [%s] called", c.Name)
}
