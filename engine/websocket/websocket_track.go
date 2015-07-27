package websocket

import (
	"fmt"
	"io"
)

// Tracker is an object that is capable of keeping tabs on events
type Tracker interface {
	Track(...interface{})
}

type tracker struct {
	out io.Writer
}

func (t *tracker) Track(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// this will turn your shit on
func Run(w io.Writer) Tracker {
	return &tracker{out: w}
}

type nilTracker struct{}

func (t *nilTracker) Track(a ...interface{}) {}

// this will turn this shit off
func Off() Tracker {
	return &nilTracker{}
}
