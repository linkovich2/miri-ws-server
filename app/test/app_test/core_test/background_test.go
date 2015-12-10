package core_test

import (
	. "github.com/jonathonharrell/miri-ws-server/app/core"
	"testing"
)

var (
	testBackgroundAll  = Background{AllowAll: true}
	testBackgroundRace = Background{
		Prerequisites: BackgroundPrerequisites{
			Races: []string{"TEST"},
		},
	}
	testBackgroundGender = Background{
		Prerequisites: BackgroundPrerequisites{
			Genders: []string{"T"},
		},
	}
	testBackgroundAestheticTrait = Background{
		Prerequisites: BackgroundPrerequisites{
			AestheticTraits: []string{"A_T"},
		},
	}
	testBackgroundFunctionalTrait = Background{
		Prerequisites: BackgroundPrerequisites{
			FunctionalTraits: []string{"F_T"},
		},
	}
)

func TestAllowAll(t *testing.T) {
	c := &Character{}
	if !testBackgroundAll.IsAllowedForCharacter(c) {
		t.Error("IsAllowedForCharacter failed when it should have passed [allow all]")
	}
}

func TestRaceReq(t *testing.T) {
	pass := &Character{Race: "TEST", Gender: "T", AestheticTraits: make(map[string][]string), FunctionalTraits: make(map[string][]string)}
	fail := &Character{Race: "FAIL", Gender: "T", AestheticTraits: make(map[string][]string), FunctionalTraits: make(map[string][]string)}
	if testBackgroundRace.IsAllowedForCharacter(fail) {
		t.Error("IsAllowedForCharacter passed when it should have failed [race req]")
	}

	if !testBackgroundRace.IsAllowedForCharacter(pass) {
		t.Error("IsAllowedForCharacter failed when it should have passed [race req]")
	}
}

func TestGenderReq(t *testing.T) {
	pass := &Character{Race: "HUMAN", Gender: "T"}
	fail := &Character{Race: "HUMAN", Gender: "M"}
	if testBackgroundGender.IsAllowedForCharacter(fail) {
		t.Error("IsAllowedForCharacter passed when it should have failed [gender req]")
	}

	if !testBackgroundGender.IsAllowedForCharacter(pass) {
		t.Error("IsAllowedForCharacter passed when it should have failed [gender req]")
	}
}

func TestAestheticTraitReq(t *testing.T) {
	pass := &Character{Race: "HUMAN", Gender: "T", AestheticTraits: map[string][]string{"T": []string{"A_T"}}, FunctionalTraits: make(map[string][]string)}
	fail := &Character{Race: "HUMAN", Gender: "T", AestheticTraits: map[string][]string{"T": []string{"BLAH"}}, FunctionalTraits: make(map[string][]string)}
	if !testBackgroundAestheticTrait.IsAllowedForCharacter(pass) {
		t.Error("IsAllowedForCharacter failed when it should have passed [aesthetic trait req]")
	}

	if testBackgroundAestheticTrait.IsAllowedForCharacter(fail) {
		t.Error("IsAllowedForCharacter passed when it should have failed [aesthetic trait req]")
	}
}

func TestFunctionalTraitReq(t *testing.T) {
	pass := &Character{Race: "HUMAN", Gender: "T", FunctionalTraits: map[string][]string{"T": []string{"F_T"}}, AestheticTraits: make(map[string][]string)}
	fail := &Character{Race: "HUMAN", Gender: "T", FunctionalTraits: map[string][]string{"T": []string{"BLAH"}}, AestheticTraits: make(map[string][]string)}
	if !testBackgroundFunctionalTrait.IsAllowedForCharacter(pass) {
		t.Error("IsAllowedForCharacter failed when it should have passed [functional trait req]")
	}

	if testBackgroundFunctionalTrait.IsAllowedForCharacter(fail) {
		t.Error("IsAllowedForCharacter passed when it should have failed [functional trait req]")
	}
}
