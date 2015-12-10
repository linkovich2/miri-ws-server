package core

import "testing"

var (
	testGender     = Gender{DisallowedRaces: []string{"TEST"}}
	testGenderOnly = Gender{Only: "TEST"}
)

func TestGenderRaceAllowed(t *testing.T) {
	if testGender.RaceAllowed("TEST") {
		t.Error("Race [TEST] should NOT be allowed by gender")
	}

	if !testGender.RaceAllowed("WHATEVER") {
		t.Error("Race [WHATEVER] should be allowed by gender")
	}
}

func TestGenderOnly(t *testing.T) {
	if testGenderOnly.RaceAllowed("FAIL") {
		t.Error("Race [FAIL] should NOT be allowed by gender [only allows TEST]")
	}

	if !testGenderOnly.RaceAllowed("TEST") {
		t.Error("Race [TEST] should be allowed by gender [only allows TEST]")
	}
}
