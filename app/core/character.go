package core

import (
	"bytes"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"

	"github.com/jonathonharrell/miri-ws-server/app/util/filters"
)

const (
	BaseCharacterMoveSpeed = 4

	StateMoving = iota
	StateLoggingOut
)

var stateStrings = map[int]string{
	StateLoggingOut: "loggingOut",
	StateMoving:     "moving",
}

type (
	Character struct {
		ID               bson.ObjectId       `bson:"_id,omitempty" json:"id"`
		Race             string              `json:"race"`
		Gender           string              `json:"gender"`
		AestheticTraits  map[string][]string `json:"aesthetic_traits"`
		FunctionalTraits map[string][]string `json:"functional_traits"`
		Background       string              `json:"background"`
		Name             string              `json:"name"`
		UserID           bson.ObjectId       `json:"-" bson:"user_id"`
		Created          time.Time           `json:"created"`
		Position         string              `json:"-"`
		Realm            string              `json:"-"`
		stats            StatBlock           `json:"-"`
		statsCached      bool                `json:"-"`
		State            []int               `json:"-"`
		Targets          []string            `json:"-"`
	}

	Stat      int
	StatBlock struct {
		Str Stat
		Dex Stat
		Con Stat
		Wis Stat
		Int Stat
		Cha Stat
	}
)

func (c *Character) SetDefaultStats() {
	var s Stat = 10
	c.stats = StatBlock{s, s, s, s, s, s}
}

func (c *Character) GetStats() StatBlock {
	if c.statsCached {
		return c.stats
	}

	// create and cache a stat block until something about the character has changed
	c.SetDefaultStats()

	// @todo calc stats from any modifying properties in core, such as traits, items, etc.
	// @note stats cannot be lower then 1 or higher then 20
	c.statsCached = true
	return c.stats
}

func (s *Stat) GetModifier() string {
	i := int(*s)
	mod := (i-i%2)/2 - 5
	if mod >= 0 {
		return strings.Join([]string{"+", strconv.Itoa(mod)}, "")
	} else {
		return strconv.Itoa(mod)
	}
}

func (s *Stat) ToInt() int {
	return int(*s)
}

func (c *Character) GetSpeed() int {
	stats := c.GetStats()
	mod, _ := strconv.Atoi(stats.Dex.GetModifier())
	return BaseCharacterMoveSpeed + mod
}

// returns a movement style for movement messages, for now walk is the only string
// but in the future we may have stuff like mount, boat, or other keywords for
// use with movement message creation
func (c *Character) GetMovementStyle() string {
	// @todo stub
	return "walk"
}

func (c *Character) GetStateString() []string {
	res := []string{}
	for _, val := range c.State {
		res = append(res, stateStrings[val])
	}

	return res
}

func (c *Character) AddState(s int) error {
	if c.HasState(s) {
		return errors.New("Character has that state already!")
	}

	c.State = append(c.State, s)
	return nil
}

func (c *Character) RemoveState(s int) {
	// @todo this one might not need error reporting, but we should think about it
	for i, val := range c.State {
		if val == s {
			c.State = append(c.State[:i], c.State[i+1:]...)
			break
		}
	}
}

func (c *Character) HasState(s int) bool {
	for _, val := range c.State {
		if val == s {
			return true
		}
	}

	return false
}

func (c *Character) FirstName() string {
	splitName := strings.Split(c.Name, " ")
	return splitName[0]
}

func (c *Character) LastName() string {
	splitName := strings.Split(c.Name, " ")
	return splitName[1]
}

func (c *Character) Describe() string {
	response := bytes.NewBuffer([]byte(c.ShortDescription()))
	response.Write([]byte("; "))

	gender := GetGender(c.Gender)
	traits := GetAestheticTraits()

	for i, cat := range c.AestheticTraits {
		for _, t := range cat {
			response.Write([]byte(filters.GenderPronouns(traits[i].Traits[t].Description, gender.Possessive, gender.Pronoun, false)))
			response.Write([]byte(" "))
		}
	}

	return response.String()
}

func (c *Character) ShortDescription() string {
	response := bytes.NewBuffer([]byte{})
	race := GetRace(c.Race)
	gender := GetGender(c.Gender)

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
