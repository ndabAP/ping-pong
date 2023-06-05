package engine

type game struct {
	width, height float64
	p1, p2        player
	ball          ball
}

func NewGame(
	gameWidth,
	gameHeight uint,
	p1,
	p2 player,
	ball ball,
) (g game) {
	g.width = float64(gameWidth)
	g.height = float64(gameHeight)
	g.p1 = p1
	g.p2 = p2
	g.ball = ball

	return
}
