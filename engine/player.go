package engine

type Player struct {
	Height, Width float64
}

func NewPlayer(w, h float64) Player {
	if w <= 0 || h <= 0 {
		panic("width and height must be greater 0")
	}
	return Player{Width: w, Height: h}
}
