package core

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

type (
	Position struct {
		X, Y, Z int
	}

	PositionRange struct {
		Max struct {
			X, Y int
		}

		Min struct {
			X, Y int
		}
	}
)

var oppositeDirections = map[string]string{
	"north":     "south",
	"south":     "north",
	"east":      "west",
	"west":      "east",
	"northeast": "southwest",
	"northwest": "southeast",
	"southeast": "northwest",
	"southwest": "northeast",
}

/**
 * Get position template string from position object: Position{14, 22, 2} => "14:22:2"
 */
func (p *Position) ToString() string {
	tmpl, _ := template.New("positionString").Parse("{{.X}}:{{.Y}}:{{.Z}}")
	var s bytes.Buffer
	tmpl.Execute(&s, p)

	return s.String()
}

func (p *Position) AdjacentPositions() map[string]string {
	north := Position{p.X, p.Y + 1, p.Z}
	south := Position{p.X, p.Y - 1, p.Z}
	northeast := Position{p.X + 1, p.Y + 1, p.Z}
	northwest := Position{p.X - 1, p.Y + 1, p.Z}
	southeast := Position{p.X + 1, p.Y - 1, p.Z}
	southwest := Position{p.X - 1, p.Y - 1, p.Z}
	east := Position{p.X + 1, p.Y, p.Z}
	west := Position{p.X - 1, p.Y, p.Z}

	return map[string]string{
		"north":     north.ToString(),
		"south":     south.ToString(),
		"northeast": northeast.ToString(),
		"northwest": northwest.ToString(),
		"southeast": southeast.ToString(),
		"southwest": southwest.ToString(),
		"east":      east.ToString(),
		"west":      west.ToString(),
	}
}

func (p *Position) Move(direction string) (Position, error) {
	direction = GetDirectionFromVariations(direction)
	switch direction {
	case "north":
		return Position{p.X, p.Y + 1, p.Z}, nil
	case "south":
		return Position{p.X, p.Y - 1, p.Z}, nil
	case "northeast":
		return Position{p.X + 1, p.Y + 1, p.Z}, nil
	case "northwest":
		return Position{p.X - 1, p.Y + 1, p.Z}, nil
	case "southeast":
		return Position{p.X + 1, p.Y - 1, p.Z}, nil
	case "southwest":
		return Position{p.X - 1, p.Y - 1, p.Z}, nil
	case "east":
		return Position{p.X + 1, p.Y, p.Z}, nil
	case "west":
		return Position{p.X - 1, p.Y, p.Z}, nil
	}

	return *p, errors.New("Direction provided was invalid")
}

/**
 * Get a Position struct from a template, IE: "15:22:2" => Position{15, 22, 2}
 */
func GetPosition(positionString string) (Position, error) {
	arr := strings.Split(positionString, ":")
	if len(arr) < 3 {
		return Position{}, errors.New("Position string must have exactly 3 keys, {X}:{Y}:{Z}")
	}

	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	z, _ := strconv.Atoi(arr[2])

	return Position{x, y, z}, nil
}

func GetOppositeDirection(direction string) (string, error) {
	if val, exists := oppositeDirections[direction]; exists {
		return val, nil
	} else {
		return "", errors.New(fmt.Sprintf("Invalid direction supplied: [%s]", direction))
	}
}

func GetDirectionFromVariations(d string) string {
	alts := map[string]string{
		"n":  "north",
		"s":  "south",
		"e":  "east",
		"w":  "west",
		"nw": "northwest",
		"ne": "northeast",
		"se": "southeast",
		"sw": "southwest",
	}

	if dir, exists := alts[d]; exists {
		return dir
	} else {
		return d
	}
}
