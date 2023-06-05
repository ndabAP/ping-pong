package engine

import (
	"errors"
	"log"
	"os"
)

var engineLogger = log.New(os.Stdout, "[ENGINE] ", 0)

var (
	errP1Win = errors.New("p1 win")
	errP2Win = errors.New("p2 win")
)
