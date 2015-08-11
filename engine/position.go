package engine

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"text/template"
)

type position struct {
	x, y, z int
}

type positionRange struct {
	max struct {
		X, Y int
	}

	min struct {
		X, Y int
	}
}

/**
 * Get position template string from position object: Position{14, 22, 2} => "14:22:2"
 */
func (p *position) toString() string {
	tmpl, _ := template.New("positionString").Parse("{{.X}}:{{.Y}}:{{.Z}}")
	var s bytes.Buffer
	tmpl.Execute(&s, p)

	return s.String()
}

/**
 * Get a Position struct from a template, IE: "15:22:2" => Position{15, 22, 2}
 */
func getPosition(positionString string) (position, error) {
	arr := strings.Split(positionString, ":")
	if len(arr) < 3 {
		return position{}, errors.New("Position string must have exactly 3 keys, {X}:{Y}:{Z}")
	}

	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	z, _ := strconv.Atoi(arr[2])

	return position{x, y, z}, nil
}
