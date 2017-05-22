package core

import (
	loggers "gopkg.in/birkirb/loggers.v1"
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
)

var logger loggers.Contextual

func init() {
	logger = stdlib.NewDefaultLogger()
}

// SetLogger sets a new logger for the core lib
// The default is the "log" library of golang
func SetLogger(l loggers.Contextual) {
	logger = l
}
