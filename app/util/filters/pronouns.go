package filters

import (
	"strings"
)

func GenderPronouns(str, possessive, nonpossessive string, you bool) string {
	if len(str) <= 0 {
		return str
	}

	str = strings.Replace(str, "[PossessivePronoun]", possessive, -1)
	str = strings.Replace(str, "[Pronoun]", nonpossessive, -1)

	if you {
		str = strings.Replace(str, "[HaveHas]", "have", -1)
		str = strings.Replace(str, "[IsAre]", "are", -1)
	} else {
		str = strings.Replace(str, "[HaveHas]", "has", -1)
		str = strings.Replace(str, "[IsAre]", "is", -1)
	}

	str = strings.ToUpper(str[:1]) + str[1:]

	return str
}
