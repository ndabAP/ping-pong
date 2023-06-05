package engine

// collision is an event type from the ball's point of view
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
	collP1 // TODO: Can be a the same time as above?
	collP2

	collBottom
	collTop

	collLeft
	collRight
)
