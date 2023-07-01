package engine

type Game struct {
	Width, Height float64
	P1, P2        Player
	Ball          Ball
}

func NewGame(
	gameWidth,
	gameHeight float64,
	p1,
	p2 Player,
	ball Ball,
) (g Game) {
	g.Width = gameWidth
	g.Height = gameHeight
	g.P1 = p1
	g.P2 = p2
	g.Ball = ball
	return
}
