package engine

type Ball struct {
	Width, Height float64
}

func NewBall(w, h float64) Ball {
	return Ball{Width: w, Height: h}
}
