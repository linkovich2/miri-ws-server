package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util/filters"

	"bytes"
)

func DescribeCharacter(c *core.Character) string {
	response := bytes.NewBuffer([]byte(ShortDescriptionForCharacter(c)))
	response.Write([]byte("; "))

	gender := content.Gender(c.Gender)
	traits := content.AestheticTraits()

	for i, cat := range c.AestheticTraits {
		for _, t := range cat {
			response.Write([]byte(filters.GenderPronouns(traits[i].Traits[t].Description, gender.Possessive, gender.Pronoun, false)))
			response.Write([]byte(" "))
		}
	}

	return response.String()
}

func ShortDescriptionForCharacter(c *core.Character) string {
	response := bytes.NewBuffer([]byte{})
	race := content.Race(c.Race)
	gender := content.Gender(c.Gender)

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

	return response.String()
}

// @todo future: need at least one more description based on deeds, skills and functional traits

// @todo this should log an error if it fails
func SaveCharacter(c *core.Character) {
	logger.Write.Info("Save Character [%s] called", c.Name)
}
