package core

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"text/template"
)

type Position struct {
	X, Y, Z int
}

type PositionRange struct {
	Max struct {
		X, Y int
	}

	Min struct {
		X, Y int
	}
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
