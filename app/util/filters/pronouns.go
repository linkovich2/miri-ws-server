package filters

import (
	"strings"
)

func GenderPronouns(str, possessive, nonpossessive string) string {
	if len(str) <= 0 {
		return str
	}

	str = strings.Replace(str, "[PossessivePronoun]", possessive, -1)
	str = strings.Replace(str, "[Pronoun]", nonpossessive, -1)
	str = strings.ToUpper(str[:1]) + str[1:]

	return str
}
