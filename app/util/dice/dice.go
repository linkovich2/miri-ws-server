package dice

import (
	"math"
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

// roll the dice!
func (dice *Dice) Roll() int {
	parsed := dice.Parse()
	min := parsed.NumDice
	max := (parsed.NumDice * parsed.DiceSize) + 1
	return dice.Random(min, max)
}

// Roll with a modifier, ex. d.RollWithModifier("+2")
func (dice *Dice) RollWithModifier(mod string) int {
	modifier, _ := strconv.Atoi(mod)
	return dice.Roll() + modifier
}

// Rolls two dice and takes the greater of the two
func (dice *Dice) RollWithAdvantage() int {
	return int(math.Max(float64(dice.Roll()), float64(dice.Roll())))
}

// Helper method for rolling with advantage (greater of two rolls), with modifier
func (dice *Dice) RollWithAdvantageAndModifier(mod string) int {
	return int(math.Max(float64(dice.RollWithModifier(mod)), float64(dice.RollWithModifier(mod))))
}

// Helper random method
func (dice *Dice) Random(min, max int) int {
	return rand.Intn(max-min) + min
}

// Helper seed method
func SeedRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}
