package filters

import (
	"bytes"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var elvishWords = map[string]string{
	"the":           "i",
	"so":            "el",
	"i":             "amn",
	"a":             "en",
	"up":            "erra",
	"you":           "per",
	"will":          "ven",
	"want":          "va",
	"me":            "sur",
	"are":           "er",
	"is":            "en",
	"in":            "pa",
	"now":           "in",
	"elven":         "Elfien",
	"come":          "para",
	"comes":         "para'n",
	"elfen":         "Elve",
	"elves":         "Elve",
	"elf":           "Elfien",
	"elvish":        "Elve",
	"elvenkind":     "Elfien",
	"heaven":        "Yggdrasil",
	"with":          "vil",
	"and":           "olsa",
	"enlightenment": "Amnfaasur",
	"human":         "Hume",
	"humans":        "Hume",
	"dwarven":       "Dar'e",
	"dwarf":         "Dar'e",
	"dwarvish":      "Dar'e",
	"can":           "ko",
	"!":             "!",
	",":             ",",
	".":             ".",
	"?":             "?",
	";":             ";",
	"'":             "'",
	":":             ":",
	"-":             "-",
}

var elvishLetters = map[string]string{
	"a": "e",
	"A": "E",
	"b": "el",
	"B": "El",
	"c": "xo",
	"C": "Xo'",
	"d": "fa",
	"D": "Fa",
	"e": "o",
	"E": "Ol",
	"f": "f",
	"F": "F",
	"g": "amn",
	"G": "Amn",
	"h": "w",
	"H": "We",
	"i": "y",
	"I": "Ym",
	"j": "ist",
	"J": "Ist",
	"k": "sa",
	"K": "Sa'",
	"l": "l",
	"L": "L",
	"m": "li",
	"M": "Li",
	"n": "t",
	"N": "Nih",
	"o": "a",
	"O": "A",
	"p": "ph",
	"P": "H",
	"q": "ath",
	"Q": "At",
	"r": "ye",
	"R": "Ygg",
	"s": "da'",
	"S": "Dra",
	"t": "er",
	"T": "Er'",
	"u": "u",
	"U": "U",
	"v": "ve",
	"V": "Ve",
	"w": "se",
	"W": "Se",
	"x": "pu",
	"X": "Pu",
	"y": "a",
	"Y": "A",
	"z": "te",
	"Z": "Te",
}

var exoticElvishLetters = map[string]string{
	"ch": "p'a",
	"Ch": "P'a",
	"co": "po",
	"Co": "L'o",
	"sh": "il",
	"Sh": "Il",
	"qu": "ge",
	"Qu": "Ge",
	"ph": "ef",
	"Ph": "Ef",
	"gh": "tir",
	"Gh": "Tir",
	"sp": "t",
	"Sp": "T",
	"st": "yl",
	"St": "Yl",
	"th": "e",
	"Th": "Eh",
	"pe": "et",
	"Pe": "Et",
	"wh": "sh",
	"Wh": "Sh",
}

func TranslateToElvish(s string, understanding int) string {
	rand.Seed(time.Now().Unix())
	// @todo pre-filter out emotes like lol, lmfao, lmao, rofl, etc.
	logger.Write.Info("%v%% understanding", understanding)

	if s == strings.ToUpper(s) {
		s = strings.ToLower(s) // force lowercase for caps lock assholes @todo move this somewhere more centralized
	}

	if understanding >= 100 {
		return "[in Elvish] " + s
	}

	s = strings.Replace(s, "!", " !", -1) // @todo move this to a convenience method
	s = strings.Replace(s, "?", " ?", -1)
	s = strings.Replace(s, ";", " ;", -1)
	s = strings.Replace(s, ",", "  ,", -1)
	s = strings.Replace(s, ".", " .", -1)
	s = strings.Replace(s, "-", " -", -1)
	s = strings.Replace(s, ":", " :", -1)

	words := strings.Fields(s)
	logger.Write.Info("%v", words)
	res := ""

	for index, word := range words {
		var buffer bytes.Buffer
		var postfix string
		var newWord []string

		for i := 0; i < len(word); i++ {
			newWord = append(newWord, "")
		}

		if understanding > 35 {
			if rand.Intn(understanding-0)+0 > 30 {
				res = res + "[" + word + "] "
				continue
			}
		}

		if elvishWord, exists := elvishWords[strings.ToLower(word)]; exists {
			if (understanding > 10 && rand.Intn(understanding-0)+0 >= 10) || understanding > 50 {
				res = res + "[" + word + "] "
				continue
			} else {
				if index == 0 {
					buffer.WriteString(strings.ToUpper(elvishWord[:1]) + elvishWord[1:])
				} else {
					buffer.WriteString(elvishWord)
				}
			}
		} else {
			if string(word[0]) == strings.ToUpper(string(word[0])) && index != 0 && words[index - 1] != "." {
				buffer.WriteString(word)
			} else {
				var keys []string
				for k := range elvishLetters {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				if len(word) > 10 {
					word = word[:len(word)-6] + word[len(word)-1:]
				} else if len(word) > 6 {
					word = word[:len(word)-4] + word[len(word)-1:]
				} else if len(word) > 3 {
					word = word[:len(word)-2] + word[len(word)-1:]
				}

				if word[len(word)-2:] == "es" {
					word = word[:len(word) - 2]
					postfix = "ner"
				} else if word[len(word)-1:] == "s" {
					word = word[:len(word) - 1]
					postfix = "nu"
				}

				for letter, e := range exoticElvishLetters {
					if strings.Contains(word, letter) {
						i := strings.Index(word, letter)
						newWord[i] = e
						word = strings.Replace(word, letter, "$$", -1)
					}
				}

				for n := len(keys) - 1; n >= 0; n-- {
					letter := keys[n]
					if strings.Contains(word, letter) {
						newWord[strings.Index(word, letter)] = elvishLetters[letter]
					}
				}
			}
		}

		for _, n := range newWord {
			buffer.WriteString(n)
		}

		res = res + buffer.String() + postfix + " "
	}

	res = strings.Replace(res, "] [", " ", -1)
	res = strings.Replace(res, " !", "!", -1) // @todo move this to a convenience method
	res = strings.Replace(res, " ?", "?", -1)
	res = strings.Replace(res, " ;", ";", -1)
	res = strings.Replace(res, " ,", ",", -1)
	res = strings.Replace(res, " .", ".", -1)
	res = strings.Replace(res, " -", "-", -1)
	res = strings.Replace(res, " :", ":", -1)
	res = strings.Replace(res, "$", "", -1)
	return strings.TrimSpace(res)
}
