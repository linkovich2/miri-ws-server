package dice

import "testing"

func TestParseDice(t *testing.T) {
	var d Dice = "3d6"
	p := d.Parse()

	if p.NumDice != 3 {
		t.Error("Expected 3 dice, got ", p.NumDice)
	}

	if p.DiceSize != 6 {
		t.Error("Expected 6-sided die, got ", p.DiceSize)
	}

	var anotherDie Dice = "2d10"
	p = anotherDie.Parse()

	if p.NumDice != 2 {
		t.Error("Expected 2 dice, got ", p.NumDice)
	}

	if p.DiceSize != 10 {
		t.Error("Expected 10-sided die, got ", p.DiceSize)
	}
}

func TestRoll(t *testing.T) {
	SeedRandom()

	var d Dice = "3d6"
	var res int

	for i := 0; i < 50; i++ {
		res = d.Roll()
		t.Logf("%d\n", res)

		if res > 18 || res < 3 {
			t.Error("Rolled outside range of 3-18: ", res)
		}
	}
}

func TestAnotherRoll(t *testing.T) {
	SeedRandom()

	var d Dice = "2d4"
	var res int

	for i := 0; i < 50; i++ {
		res = d.Roll()
		t.Logf("Rolled 2d4: %d\n", res)

		if res > 8 || res < 2 {
			t.Error("Rolled outside of range 2-8: ", res)
		}
	}
}

func TestRollWithModifier(t *testing.T) {
	SeedRandom()

	var d Dice = "1d20"
	var res int
	for i := 0; i < 50; i++ {
		res = d.RollWithModifier("+2")
		t.Logf("%d\n", res)

		if res > 22 || res < 3 {
			t.Error("Rolled outside range of 3-22: ", res)
		}
	}
}
