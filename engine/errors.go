package engine

import (
	"errors"
)

var (
	ErrP1Win = errors.New("p1 win")
	ErrP2Win = errors.New("p2 win")
)
