package main

import (
	"context"
	"log"
	"os"
)

type Game struct {
	debug bool

	width, height float64

	p1, p2 player

	ball ball

	engine *canvasEngine
}

var gameLogger = log.New(os.Stdout, "[GAME] ", 0)

func NewGame(gWidth, gHeight int64, p1Width, p1Height, p2Width, p2height int64, bWidth, bHeight int64) (g Game) {
	if *debug {
		g.debug = *debug
		gameLogger.Println("debug mode")
	}

	// Validate
	if gWidth%2 != 0 ||
		gHeight%2 != 0 ||
		p1Width%2 != 0 ||
		p1Height%2 != 0 ||
		p2Width%2 != 0 ||
		p2height%2 != 0 ||
		bWidth%2 != 0 ||
		bHeight%2 != 0 {
		panic("values must be dividable by two")
	}

	g.width = float64(gWidth)
	g.height = float64(gHeight)

	g.p1 = newPlayer1(float64(p1Width), float64(p1Height))
	g.p2 = newPlayer2(float64(p2Width), float64(p2height))

	g.ball = newBall(float64(bWidth), float64(bHeight))

	g.engine = newCanvasEngine(g)

	return
}

func (g Game) Start(ctx context.Context, frames chan []byte) {
	gameLogger.Println("start ...")

	g.engine.bootstrap().writeFrames(ctx, frames)
}
