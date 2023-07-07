package engine

type Game struct {
	Width, Height float64
	P1, P2        Player
	Ball          Ball
}

func NewGame(
	w,
	h float64,
	p1,
	p2 Player,
	ball Ball,
) (g Game) {
	if w <= 0 || h <= 0 {
		panic("width and height must be greater 0")
	}
	g.Width = w
	g.Height = h
	g.P1 = p1
	g.P2 = p2
	g.Ball = ball
	return
}
