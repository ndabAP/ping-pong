package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	// Real random
	rand.Seed(time.Now().UnixNano())
}

var engineLogger = log.New(os.Stdout, "[ENGINE] ", 0)

type canvasEngine struct {
	game Game

	ballX float64
	ballY float64
	p1X   float64
	p1Y   float64
	p2X   float64
	p2Y   float64

	p1YVelocity   float64
	p2YVelocity   float64
	ballXVelocity float64
	ballYVelocity float64

	fps float64
	tps float64
}

const (
	// Frames per second
	_fps = 50.0
	// Ticks per second
	_tps = 1000.0 * (1.0 / _fps)
)

func newCanvasEngine(g Game) *canvasEngine {
	e := &canvasEngine{}

	e.game = g

	e.fps = _fps
	e.tps = _tps

	return e
}

func (e *canvasEngine) bootstrap() *canvasEngine {
	engineLogger.Println("bootstrap ...")

	time.Sleep(750 * time.Millisecond)

	e.resetBall().resetPlayers().randomBall().log()

	return e
}

func (e *canvasEngine) writeFrames(gameCtx context.Context, frames chan []byte) {
	go func() {
		// tps ticks or millseconds for 1 frame, since: _tps * _fps = y
		ticker := time.NewTicker(_tps * time.Millisecond)
		ticks := 0

		for {
			select {
			case <-ticker.C:
				engineLogger.Println("next tick ...")

				if err := e.advanceGame(); err != nil {
					// Restart
					e.bootstrap().writeFrames(gameCtx, frames)
					return
				}

				e.log()

				jsonFrame, err := e.jsonFrame()
				if err != nil {
					panic(err)
				}
				frames <- jsonFrame

				ticks++

				engineLogger.Printf("ticks %d", ticks)

			case <-gameCtx.Done():
				ticker.Stop()

				return
			}
		}
	}()
}

func (e *canvasEngine) jsonFrame() ([]byte, error) {
	return json.Marshal(e.mapFrame())
}

const (
	baseline           = 0
	default_padding    = 15
	default_ball_x_vel = 0.25
	min_y_vel          = 0.1
	max_y_vel          = 0.15
)

func (e *canvasEngine) ballDirP1() bool {
	return e.ballX <= e.game.width/2
}

func (e *canvasEngine) ballDirP2() bool {
	return e.ballX >= e.game.width/2
}

func (e *canvasEngine) log() *canvasEngine {
	jsonBytes, err := json.MarshalIndent(e.mapFrame(), "", "	")
	if err != nil {
		panic(err)
	}
	engineLogger.Printf("%s", jsonBytes)
	return e
}

func (e *canvasEngine) mapFrame() map[string]interface{} {
	return map[string]interface{}{
		"gameWidth":     e.game.width,
		"gameHeight":    e.game.height,
		"p1Width":       e.game.p1.width,
		"p1Height":      e.game.p1.height,
		"p2Width":       e.game.p2.width,
		"p2Height":      e.game.p2.height,
		"ballWidth":     e.game.ball.width,
		"ballHeight":    e.game.ball.height,
		"ballX":         e.ballX,
		"ballY":         e.ballY,
		"p1X":           e.p1X,
		"p1Y":           e.p1Y,
		"p2X":           e.p2X,
		"p2Y":           e.p2Y,
		"p1YVelocity":   e.p1YVelocity,
		"p2YVelocity":   e.p2YVelocity,
		"ballXVelocity": e.ballXVelocity,
		"ballYVelocity": e.ballYVelocity,
	}
}
