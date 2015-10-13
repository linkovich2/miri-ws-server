package logger

import (
	"fmt"
	"github.com/op/go-logging"
)

// set up logging
var Write = logging.MustGetLogger("miri-logger")
var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}",
)

func NewLine() {
	fmt.Printf("\n")
}
