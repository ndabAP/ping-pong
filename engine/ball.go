package engine

type ball struct {
	width, height float64
}

func NewBall(w, h uint) ball {
	return ball{width: float64(w), height: float64(h)}
}
