//usr/bin/env go run "$0" "$@"; exit
package main

import (
	"context"

	"github.com/ndabAP/ping-pong/canvas"
	"github.com/ndabAP/ping-pong/engine"
)

func main() {
	game := engine.NewGame(
		800,
		600,
		engine.NewPlayer(10, 150),
		engine.NewPlayer(10, 150),
		engine.NewBall(5, 5),
	)
	engine := canvas.New(game)

	// Frames
	framesch := make(chan []byte)
	go func() {
		for range framesch {
			// ...
		}
	}()
	// User input
	inputch := make(chan []byte, 1)
	defer close(inputch)
	go func() {
		for {
			buf := make([]byte, canvas.INPUT_BUF_SIZE)
			// ...
			inputch <- buf
		}
	}()

	ctx := context.Background()
	engine.NewRound(ctx, framesch, inputch)
}
