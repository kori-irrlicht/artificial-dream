package core

import (
	loggers "gopkg.in/birkirb/loggers.v1"
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
)

var logger loggers.Contextual

func init() {
	logger = stdlib.NewDefaultLogger()
}

func SetLogger(l loggers.Contextual) {
	logger = l
}
