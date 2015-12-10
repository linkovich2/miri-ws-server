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
	// pass := &Character{}
}
