//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"context"

	"github.com/ndabAP/ping-pong/canvas"
	"github.com/ndabAP/ping-pong/engine"
)

func main() {
	framesch := make(chan []byte)
	inputch := make(chan []byte, 1)
	g := newEngine(framesch, inputch)

	ctx := context.Background()

	// Frames
	go func() {
		for range framesch {
			// ...
		}
	}()
	// AI input
	defer close(inputch)
	go func() {
		for {
			buf := make([]byte, canvas.INPUT_BUF_SIZE)
			// ...
			inputch <- buf
		}
	}()
	g.NewRound(ctx, framesch, inputch)
}

func newEngine(framesch, inputch chan []byte) *canvas.CanvasEngine {
	game := engine.NewGame(
		800,
		600,
		engine.NewPlayer(10, 150),
		engine.NewPlayer(10, 150),
		engine.NewBall(5, 5),
	)
	engine := canvas.New(game)
	return engine
}
