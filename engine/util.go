package engine

import (
  "time"

  "github.com/op/go-logging"
)

// set up logging
var log = logging.MustGetLogger("miri-logger")
var format = logging.MustStringFormatter(
  "%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func RunEvery(d time.Duration, f func()) {
  for _ = range time.Tick(d) {
    f()
  }
}
