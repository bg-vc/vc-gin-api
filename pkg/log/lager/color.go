package lager

import (
	"fmt"
)

const (
	//Black is a constant of type int
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var (
	//InfoByte is a variable of type []byte
	DebugByte = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Blue, "DEBUG"))
	WarnByte  = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Yellow, "WARN"))
	ErrorByte = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Red, "ERROR"))
	FatalByte = []byte(fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", Magenta, "FATAL"))
)
