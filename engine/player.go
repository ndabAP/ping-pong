package engine

type Player struct {
	Height, Width float64
}

func NewPlayer(w, h float64) Player {
	return Player{Width: w, Height: h}
}
