package core

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

const (
	BaseCharacterMoveSpeed = 4

	StateDefault = iota
	StateMoving
)

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
		State            int                 `json:"-"`
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
