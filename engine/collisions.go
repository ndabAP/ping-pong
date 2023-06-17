package engine

type collision int

const (
	collNone collision = 0

	collBottomLeft collision = iota
	collTopLeft
	collBottomRight
	collTopRight

	collP1Bottom
	collP1Top
	collP2Bottom
	collP2Top
	collP1
	collP2

	collBottom
	collTop

	collLeft
	collRight
)
