package main

type player struct {
	height, width float64

	number int
}

func newPlayer1(w, h float64) player {
	return player{width: w, height: h, number: 1}
}

func newPlayer2(w, h float64) player {
	return player{width: w, height: h, number: 2}
}
