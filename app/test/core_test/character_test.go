package core_test

import (
	"fmt"
	. "github.com/jonathonharrell/miri-ws-server/app/core"
	"testing"
)

func TestCharacterStats(t *testing.T) {
	c := &Character{}
	stats := c.GetStats()

	if stats.Str.ToInt() != 10 {
		t.Error("Character should have 10 STR")
	}

	mod := stats.Str.GetModifier()
	if mod != "+0" {
		t.Error(fmt.Sprintf("Character's STR mod should be +0, not %s", mod))
	}

	stats.Str = 19
	mod = stats.Str.GetModifier()
	if mod != "+4" {
		t.Error(fmt.Sprintf("Character's STR mod should be +4, not %s", mod))
	}

	stats.Con = 4
	mod = stats.Con.GetModifier()
	if mod != "-3" {
		t.Error(fmt.Sprintf("Character's CON mod should be -3, not %s", mod))
	}
}
