package engine

type Collision int

const (
	CollNone Collision = 0

	CollBottomLeft Collision = iota
	CollTopLeft
	CollBottomRight
	CollTopRight

	CollP1Bottom
	CollP1Top
	CollP2Bottom
	CollP2Top
	CollP1
	CollP2

	CollBottom
	CollTop

	CollLeft
	CollRight
)
