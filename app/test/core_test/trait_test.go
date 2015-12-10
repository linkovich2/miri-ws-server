package core_test

import (
	. "github.com/jonathonharrell/miri-ws-server/app/core"
	"testing"
)

var (
	testTrait = AestheticTrait{
		DisallowedRaces:   []string{"TEST"},
		DisallowedGenders: []string{"T"},
	}
	testTraitOnly = AestheticTrait{Only: "TEST"}
)

func TestTraitValidateCharacterAgainst(t *testing.T) {
	onlyPass := &Character{Race: "TEST"}
	onlyFail := &Character{Race: "NOPE"}
	racePass := onlyFail
	raceFail := onlyPass
	genderPass := &Character{Race: "T", Gender: "F"}
	genderFail := &Character{Race: "T", Gender: "T"}

	if !testTraitOnly.IsAllowedForCharacter(onlyPass) {
		t.Error("Only should have passed")
	}
	if testTraitOnly.IsAllowedForCharacter(onlyFail) {
		t.Error("Only should have failed")
	}
	if !testTrait.IsAllowedForCharacter(genderPass) {
		t.Error("Gender should have passed")
	}
	if testTrait.IsAllowedForCharacter(genderFail) {
		t.Error("Gender should have failed")
	}
	if !testTrait.IsAllowedForCharacter(racePass) {
		t.Error("Race should have passed")
	}
	if testTrait.IsAllowedForCharacter(raceFail) {
		t.Error("Race should have failed")
	}
}
