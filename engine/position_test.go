package engine

import "testing"

func testToString(t *testing.T) {
	p := position{1, 1, 1}
	s := p.toString()

	if s != "1:1:1" {
		t.Error("Expected 1:1:1, got ", s)
	}
}

func testGetPosition(t *testing.T) {
	pString := "248:122:3"
	s, _ := getPosition(pString)

	if s.x != 248 {
		t.Error("Expected position object, got ", s.x)
	}

	if s.y != 122 {
		t.Error("Expected position object, got ", s.y)
	}

	if s.z != 3 {
		t.Error("Expected position object, got ", s.z)
	}
}

func testGetPositionError(t *testing.T) {
	pString := "222"
	_, err := getPosition(pString)

	if err == nil {
		t.Error("Expected error, got ", err)
	}
}
