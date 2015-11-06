package app

import "testing"

func TestToString(t *testing.T) {
	p := Position{1, 1, 1}
	s := p.ToString()

	if s != "1:1:1" {
		t.Error("Expected 1:1:1, got ", s)
	}
}

func TestGetPosition(t *testing.T) {
	pString := "248:122:3"
	s, _ := GetPosition(pString)

	if s.X != 248 {
		t.Error("Expected position object, got ", s.X)
	}

	if s.Y != 122 {
		t.Error("Expected position object, got ", s.Y)
	}

	if s.Z != 3 {
		t.Error("Expected position object, got ", s.Z)
	}
}

func TestGetPositionError(t *testing.T) {
	pString := "222"
	_, err := GetPosition(pString)

	if err == nil {
		t.Error("Expected error, got ", err)
	}
}
