package engine

type player struct {
	height, width float64
}

func NewPlayer(w, h uint) player {
	return player{width: float64(w), height: float64(h)}
}
