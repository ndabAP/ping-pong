package engine

import (
	"bytes"
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"time"
)

// CanvasEngine is a ping-pong engine for browsers with Canvas support
type CanvasEngine struct {
	// Static
	fps, tps float64

	game game

	// Dynamic
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

// NewCanvasEngine returns a new Canvas engine for browsers with Canvas support
func NewCanvasEngine(g game) *CanvasEngine {
	const def_fps = 50

	e := new(CanvasEngine)
	e.game = g
	e.fps = def_fps
	e.tps = 1000.0 / def_fps

	return e
}

// SetDebug sets the Canvas engines debug state
func (e *CanvasEngine) SetDebug(debug bool) *CanvasEngine {
	engineLogger.Printf("debug %t", debug)
	e.debug = debug
	return e
}

func (e *CanvasEngine) SetFPS(fps uint) *CanvasEngine {
	if int(fps)%2 != 0 {
		panic("fps must be dividable by two")
	}
	if fps == 0 {
		panic("fps must be greater zero")
	}
	engineLogger.Printf("fps %d", fps)
	e.fps = float64(fps)
	return e
}

// NewRound resets the ball, players and starts a new round. It accepts
// a frames channel to write into and input channel to read from
func (e *CanvasEngine) NewRound(ctx context.Context, frames chan<- []byte, input <-chan []byte) {
	engineLogger.Println("new round")

	time.Sleep(time.Millisecond * 1500) // 1.5 seconds

	e.resetEngine()

	// Calculates and writes frames
	go func() {
		clock := time.NewTicker(time.Duration(e.tps) * time.Millisecond)
		defer clock.Stop()

		for range clock.C {
			e.tick()

			switch e.err {
			case errP1Win:
				engineLogger.Println("p1 wins")
				e.p1Score += 1

				e.NewRound(ctx, frames, input)
				return

			case errP2Win:
				engineLogger.Println("p2 wins")
				e.p2Score += 1

				e.NewRound(ctx, frames, input)
				return
			}

			jsonTick, _ := json.Marshal(e)
			frames <- jsonTick

			engineLogger.Printf("tick: %s", string(jsonTick))
		}
	}()

	// Reads user input and moves player one according to it
	go func() {
		for key := range input {
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

	close(frames)
}

func (e *CanvasEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"fps": e.fps,
		"tps": e.tps,

		"gameWidth":  int(e.game.width),
		"gameHeight": int(e.game.height),
		"p1Width":    int(e.game.p1.width),
		"p1Height":   int(e.game.p1.height),
		"p2Width":    int(e.game.p2.width),
		"p2Height":   int(e.game.p2.height),
		"ballWidth":  int(e.game.ball.width),
		"ballHeight": int(e.game.ball.height),

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

// tick calculates the next frame
func (e *CanvasEngine) tick() {
	switch e.detectColl() {

	case
		collP1Top,
		collP1Bottom,
		collP2Top,
		collP2Bottom:
		e.inverseBallXYVelocity().deOutOfBoundsBall()

	case
		collP1,
		collP2:
		e.inverseBallXVelocity().deOutOfBoundsBall()

	case
		collBottomLeft,
		collTopLeft:
		e.err = errP2Win
		return

	case
		collBottomRight,
		collTopRight:
		e.err = errP1Win
		return

	case
		collTop,
		collBottom:
		e.inverseBallYVelocity().deOutOfBoundsBall()

	case collLeft:
		e.err = errP2Win
		return

	case collRight:
		e.err = errP1Win
		return

	case collNone:

	}

	e.advanceAll().deOutOfBoundsPlayers()
}

// Constants

const (
	baseline                 = 0
	default_padding          = 15
	canvas_border_correction = 1

	default_ball_x_vel_ratio = 0.25
	min_ball_y_vel_ratio     = 0.1
	max_y_vel_ratio          = 0.20

	magic_p = 3

	pinput_dist = 4
)

// State

func (e *CanvasEngine) ballDirP1() bool {
	return e.ballX <= e.game.width/2
}

func (e *CanvasEngine) ballDirP2() bool {
	return e.ballX >= e.game.width/2
}

// Collisions

// detectColl detects and returns a possible collision
func (e *CanvasEngine) detectColl() collision {
	switch {
	case e.isCollBottomLeft():
		return collBottomLeft

	case e.isCollTopLeft():
		return collTopLeft

	case e.isCollBottomRight():
		return collBottomLeft

	case e.isCollTopRight():
		return collTopRight

	case e.isCollP1Bottom():
		return collP1Bottom

	case e.isCollP1Top():
		return collP1Top

	case e.isCollP2Bottom():
		return collP2Bottom

	case e.isCollP2Top():
		return collP2Top

	case e.isCollP1():
		return collP1

	case e.isCollP2():
		return collP2

	case e.isCollBottom():
		return collBottom

	case e.isCollTop():
		return collTop

	case e.isCollLeft():
		return collLeft

	case e.isCollRight():
		return collRight

	default:
		return collNone
	}
}

func (e *CanvasEngine) isCollP1() bool {
	x := e.ballX <= (e.p1X + e.game.p1.width + magic_p)
	y1 := e.p1Y <= e.ballY
	y2 := (e.p1Y + e.game.p1.height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *CanvasEngine) isCollP2() bool {
	x := (e.ballX + e.game.ball.height) >= e.p2X
	y1 := e.p2Y <= e.ballY
	y2 := (e.p2Y + e.game.p2.height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *CanvasEngine) isCollTop() bool {
	return e.ballY <= baseline+e.game.ball.height+canvas_border_correction
}

func (e *CanvasEngine) isCollBottom() bool {
	return e.ballY+e.game.ball.height >= e.game.height-canvas_border_correction
}

func (e *CanvasEngine) isCollLeft() bool {
	return e.ballX-e.game.ball.height-canvas_border_correction <= 0
}

func (e *CanvasEngine) isCollRight() bool {
	return e.ballX+e.game.ball.height+canvas_border_correction >= e.game.width
}

func (e *CanvasEngine) isCollP1Top() bool {
	return e.isCollP1() && e.isCollTop()
}

func (e *CanvasEngine) isCollP2Top() bool {
	return e.isCollP2() && e.isCollTop()
}

func (e *CanvasEngine) isCollP1Bottom() bool {
	return e.isCollP1() && e.isCollBottom()
}

func (e *CanvasEngine) isCollP2Bottom() bool {
	return e.isCollP2() && e.isCollBottom()
}

func (e *CanvasEngine) isCollTopLeft() bool {
	return e.isCollTop() && e.isCollLeft()
}

func (e *CanvasEngine) isCollBottomLeft() bool {
	return e.isCollBottom() && e.isCollLeft()
}

func (e *CanvasEngine) isCollTopRight() bool {
	return e.isCollTop() && e.isCollRight()
}

func (e *CanvasEngine) isCollBottomRight() bool {
	return e.isCollBottom() && e.isCollRight()
}

// Mutations

func (e *CanvasEngine) resetEngine() *CanvasEngine {
	e.err = nil
	return e.resetBall().resetPlayers()
}

func (e *CanvasEngine) resetBall() *CanvasEngine {
	// Center ball
	e.ballX = e.game.width / 2.0
	e.ballY = e.game.height / 2.0
	// Random direction
	if rand.Intn(10) < 5 {
		e.ballXVelocity = -default_ball_x_vel_ratio * e.game.width
		y := min_ball_y_vel_ratio*e.game.height + rand.Float64()*((max_y_vel_ratio*e.game.height)-(min_ball_y_vel_ratio*e.game.height))
		e.ballYVelocity = -y
	} else {
		e.ballXVelocity = default_ball_x_vel_ratio * e.game.width
		y := min_ball_y_vel_ratio*e.game.height + rand.Float64()*((max_y_vel_ratio*e.game.height)-(min_ball_y_vel_ratio*e.game.height))
		e.ballYVelocity = y
	}
	return e
}

func (e *CanvasEngine) resetPlayers() *CanvasEngine {
	// P1
	e.p1X = 0 + default_padding
	e.p1Y = e.game.height/2 - e.game.p1.height/2
	// P2
	e.p2X = e.game.width - +e.game.p1.width - default_padding
	e.p2Y = e.game.height/2 - e.game.p2.height/2
	return e
}

func (e *CanvasEngine) advanceAll() *CanvasEngine {
	return e.advanceBall().advancePlayers()
}

// advanceBall advances the ball one tick or frame
func (e *CanvasEngine) advanceBall() *CanvasEngine {
	e.ballX += e.ballXVelocity / e.fps
	e.ballY += e.ballYVelocity / e.fps
	return e
}

// advancePlayers advances the players one tick or frame
func (e *CanvasEngine) advancePlayers() *CanvasEngine {
	switch {
	case e.ballDirP1():
		e.p2YVelocity = 0

	case e.ballDirP2():
		switch y := (e.p2Y + (e.game.p2.height / 2)) - e.ballY; {
		case y > 0:
			e.p2YVelocity = max_y_vel_ratio * e.game.height
			e.p2Y -= e.p2YVelocity / e.fps
		case y < 0:
			e.p2YVelocity = max_y_vel_ratio * e.game.height
			e.p2Y += e.p2YVelocity / e.fps
		case y > -0.9 && y < 0.9:
			e.p2YVelocity = 0
		}
	}

	return e
}

func (e *CanvasEngine) p1Up() *CanvasEngine {
	e.p1YVelocity = pinput_dist
	e.p1Y += pinput_dist
	e.p1YVelocity = 0
	return e
}

func (e *CanvasEngine) p1Down() *CanvasEngine {
	e.p1YVelocity = pinput_dist
	e.p1Y -= pinput_dist
	e.p1YVelocity = 0
	return e
}

func (e *CanvasEngine) inverseBallXYVelocity() *CanvasEngine {
	return e.inverseBallXVelocity().inverseBallYVelocity()
}

func (e *CanvasEngine) inverseBallXVelocity() *CanvasEngine {
	if e.ballXVelocity > 0 {
		e.ballXVelocity = e.ballXVelocity * -1
	} else {
		e.ballXVelocity = math.Abs(e.ballXVelocity)
	}
	return e
}

func (e *CanvasEngine) inverseBallYVelocity() *CanvasEngine {
	if e.ballYVelocity > 0 {
		e.ballYVelocity = e.ballYVelocity * -1
	} else {
		e.ballYVelocity = math.Abs(e.ballYVelocity)
	}
	return e
}

func (e *CanvasEngine) deOutOfBoundsPlayers() *CanvasEngine {
	// P1, top
	if e.p1Y-default_padding <= baseline {
		e.p1Y = baseline + default_padding
		e.p1YVelocity = 0
	}
	// P1, bottom
	if e.p1Y+e.game.p1.height >= e.game.height-default_padding {
		e.p1Y = e.game.height - e.game.p1.height - default_padding
		e.p1YVelocity = 0
	}
	// P2, top
	if e.p2Y-default_padding <= baseline {
		e.p2Y = baseline + default_padding
		e.p2YVelocity = 0
	}
	// P2, bottom
	if e.p2Y+e.game.p2.height >= e.game.height-default_padding {
		e.p2Y = e.game.height - e.game.p2.height - default_padding
		e.p2YVelocity = 0
	}
	return e
}

func (e *CanvasEngine) deOutOfBoundsBall() *CanvasEngine {
	// Top
	if e.ballY <= baseline {
		e.ballY = baseline - 1
	}
	// Bottom
	if e.ballY+e.game.ball.height >= e.game.height {
		e.ballY = e.game.height - e.game.ball.height - 1
	}
	// P1
	if e.ballX-e.game.ball.width <= e.p1X+e.game.p1.width {
		e.ballX = e.p1X + e.game.p1.width + magic_p
	}
	// P2
	if e.ballX+e.game.ball.width >= e.p2X {
		e.ballX = e.p2X - magic_p
	}
	return e
}
