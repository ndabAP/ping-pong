package main

type player struct {
	height, width float64

	score uint
}

func newPlayer1(w, h float64) player {
	return player{width: w, height: h, score: 0}
}

func newPlayer2(w, h float64) player {
	return player{width: w, height: h, score: 0}
}
