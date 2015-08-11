package util

import (
	"time"
)

func RunEvery(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}
