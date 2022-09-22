package main

import (
	"context"
	"log"
	"os"
)

type Game struct {
	width, height float64

	p1, p2 player

	ball ball

	engine *canvasEngine
}

var gameLogger = log.New(os.Stdout, "[GAME] ", 0)

func NewGame(gWidth, gHeight float64, p1Width, p1Height, p2Width, p2height float64, bWidth, bHeight float64) (g Game) {
	g.width = gWidth
	g.height = gHeight

	g.p1 = newPlayer1(p1Width, p1Height)
	g.p2 = newPlayer2(p2Width, p2height)

	g.ball = newBall(bWidth, bHeight)

	g.engine = newCanvasEngine(g)

	return
}

func (g Game) Start(ctx context.Context, frames chan []byte) {
	gameLogger.Println("start ...")

	g.engine.bootstrap().writeFrames(ctx, frames)
}
