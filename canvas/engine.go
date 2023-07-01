package canvas

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ndabAP/ping-pong/engine"
)

// CanvasEngine is a ping-pong engine for browsers with Canvas support
type CanvasEngine struct {
	// Static
	fps, tps float64

	game engine.Game

	// State
	p1Score, p2Score int

	ballX, ballY       float64
	p1X, p1Y, p2X, p2Y float64

	p1YVelocity, p2YVelocity     float64
	ballXVelocity, ballYVelocity float64

	// Error of the current tick
	err error

	// Engine debug state
	debug bool
}

var engineLogger = log.New(os.Stdout, "[ENGINE] ", 0)

// New returns a new Canvas engine for browsers with Canvas support
func New(g engine.Game) *CanvasEngine {
	e := new(CanvasEngine)
	e.game = g
	e.fps = DEFAULT_FPS
	e.tps = 1000.0 / e.fps

	return e
}

// SetDebug sets the Canvas engines debug state
func (e *CanvasEngine) SetDebug(debug bool) *CanvasEngine {
	engineLogger.Printf("debug %t", debug)
	e.debug = debug
	return e
}

func (e *CanvasEngine) SetFPS(fps uint) *CanvasEngine {
	if fps <= 0 {
		panic("fps must be greater zero")
	}
	engineLogger.Printf("fps %d", fps)
	e.fps = float64(fps)
	e.tps = 1000.0 / e.fps
	return e
}

// NewRound resets the ball, players and starts a new round. It accepts
// a frames channel to write into and input channel to read from
func (e *CanvasEngine) NewRound(ctx context.Context, framesch chan<- []byte, inputch <-chan []byte) {
	engineLogger.Println("new round")

	time.Sleep(time.Millisecond * 1500) // 1.5 seconds

	e.reset()

	// Calculates and writes frames
	go func() {
		clock := time.NewTicker(time.Duration(e.tps) * time.Millisecond)
		defer clock.Stop()

		for range clock.C {
			e.tick()

			switch e.err {
			case engine.ErrP1Win:
				engineLogger.Println("p1 wins")
				e.p1Score += 1

				e.NewRound(ctx, framesch, inputch)
				return

			case engine.ErrP2Win:
				engineLogger.Println("p2 wins")
				e.p2Score += 1

				e.NewRound(ctx, framesch, inputch)
				return
			}

			jsonTick, _ := json.Marshal(e)
			framesch <- jsonTick

			engineLogger.Printf("tick: %s", string(jsonTick))
		}
	}()

	// Reads user input and moves player one according to it
	go func() {
		for key := range inputch {
			key = bytes.Trim(key, "\x00")

			switch k := string(key); k {
			case "ArrowUp":
				engineLogger.Printf("key %s", k)
				e.p1Down() // The Canvas origin is top left

			case "ArrowDown":
				engineLogger.Printf("key %s", k)
				e.p1Up()
			}
		}
	}()

	<-ctx.Done()
	engineLogger.Println("exiting")

	close(framesch)
}

func (e *CanvasEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"fps": e.fps,
		"tps": e.tps,

		"gameWidth":  int(e.game.Width),
		"gameHeight": int(e.game.Height),
		"p1Width":    int(e.game.P1.Width),
		"p1Height":   int(e.game.P1.Height),
		"p2Width":    int(e.game.P2.Width),
		"p2Height":   int(e.game.P2.Height),
		"ballWidth":  int(e.game.Ball.Width),
		"ballHeight": int(e.game.Ball.Height),

		"p1Score": e.p1Score,
		"p2Score": e.p2Score,

		"ballX": int(e.ballX),
		"ballY": int(e.ballY),
		"p1X":   int(e.p1X),
		"p1Y":   int(e.p1Y),
		"p2X":   int(e.p2X),
		"p2Y":   int(e.p2Y),

		"p1YVelocity":   int(e.p1YVelocity),
		"p2YVelocity":   int(e.p2YVelocity),
		"ballXVelocity": int(e.ballXVelocity),
		"ballYVelocity": int(e.ballYVelocity),

		"error": e.err,

		"debug": e.debug,
	})
}
