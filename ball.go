package main

type ball struct {
	width, height float64
}

func newBall(w, h float64) ball {
	return ball{width: w, height: h}
}
