package dice

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type (
	Dice       string
	ParsedDice struct {
		NumDice  int
		DiceSize int
	}
)

func (dice *Dice) Parse() ParsedDice {
	arr := strings.Split(string(*dice), "d")
	numDice, _ := strconv.Atoi(arr[0])
	diceSize, _ := strconv.Atoi(arr[1])

	return ParsedDice{numDice, diceSize}
}

func (dice *Dice) Roll() int {
	parsed := dice.Parse()
	min := parsed.NumDice
	max := (parsed.NumDice * parsed.DiceSize) + 1
	return dice.Random(min, max)
}

func (dice *Dice) RollWithModifier(mod string) int {
	modifier, _ := strconv.Atoi(mod)
	return dice.Roll() + modifier
}

func (dice *Dice) Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func SeedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}
