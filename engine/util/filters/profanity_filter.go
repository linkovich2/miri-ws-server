package filters

import (
	"math/rand"
	"strings"
	"time"
)

var c map[string][]string

func Init() {
	c = make(map[string][]string)
	c["fuck"] = []string{"frick", "feck"} // @todo we add our filter options here
	// @todo would be better probably to load this in from JSON so we can edit them from a tool more easily

	// For things like "lol", we need to filter based on that and replace it if it's part of the string
	//   or forward to an action if it's the whole string
	//   call this the "RP Filter"
}

func ProfanityFilter(s string) string {
	rand.Seed(int64(time.Now().Nanosecond()))

	for word, replacements := range c {
		s = strings.Replace(s, word, replacements[rand.Intn(len(replacements))], -1)
	}

	return s
}
