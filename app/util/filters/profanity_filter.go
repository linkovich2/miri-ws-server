package filters

import (
	"math/rand"
	"strings"
	"time"
)

var profanityMap = map[string][]string{
	"fuck": []string{"frick", "feck"},
}

func ReplaceProfanity(s string) string {
	rand.Seed(int64(time.Now().Nanosecond()))

	for word, replacements := range profanityMap {
		s = strings.Replace(s, word, replacements[rand.Intn(len(replacements))], -1)
	}

	return s
}
