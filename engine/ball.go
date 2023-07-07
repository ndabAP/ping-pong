package engine

type Ball struct {
	Width, Height float64
}

func NewBall(w, h float64) Ball {
	if w <= 0 || h <= 0 {
		panic("width and height must be greater 0")
	}
	return Ball{Width: w, Height: h}
}
