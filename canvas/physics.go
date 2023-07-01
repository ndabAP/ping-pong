package canvas

import (
	"math"
	"math/rand"

	"github.com/ndabAP/ping-pong/engine"
)

const (
	DEFAULT_FPS    = 60
	INPUT_BUF_SIZE = 2 << 8
)

// tick calculates the next frame
func (e *CanvasEngine) tick() {
	switch e.detectColl() {

	case
		engine.CollP1Top,
		engine.CollP1Bottom,
		engine.CollP2Top,
		engine.CollP2Bottom:
		e.inverseBallXYVelocity().deOutOfBoundsBall()

	case
		engine.CollP1,
		engine.CollP2:
		e.inverseBallXVelocity().deOutOfBoundsBall()

	case
		engine.CollBottomLeft,
		engine.CollTopLeft:
		e.err = engine.ErrP2Win
		return

	case
		engine.CollBottomRight,
		engine.CollTopRight:
		e.err = engine.ErrP2Win
		return

	case
		engine.CollTop,
		engine.CollBottom:
		e.inverseBallYVelocity().deOutOfBoundsBall()

	case engine.CollLeft:
		e.err = engine.ErrP2Win
		return

	case engine.CollRight:
		e.err = engine.ErrP1Win
		return

	case engine.CollNone:
	}

	e.advance().deOutOfBoundsPlayers()
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

	player_input_dist = 4
)

// State

func (e *CanvasEngine) ballDirP1() bool {
	return e.ballX <= e.game.Width/2
}

func (e *CanvasEngine) ballDirP2() bool {
	return e.ballX >= e.game.Width/2
}

// Collisions

// detectColl detects and returns a possible collision
func (e *CanvasEngine) detectColl() engine.Collision {
	switch {
	case e.isCollBottomLeft():
		return engine.CollBottomLeft

	case e.isCollTopLeft():
		return engine.CollTopLeft

	case e.isCollBottomRight():
		return engine.CollBottomRight

	case e.isCollTopRight():
		return engine.CollTopRight

	case e.isCollP1Bottom():
		return engine.CollP1Bottom

	case e.isCollP1Top():
		return engine.CollP1Top

	case e.isCollP2Bottom():
		return engine.CollP2Bottom

	case e.isCollP2Top():
		return engine.CollP2Top

	case e.isCollP1():
		return engine.CollP1

	case e.isCollP2():
		return engine.CollP2

	case e.isCollBottom():
		return engine.CollBottom

	case e.isCollTop():
		return engine.CollTop

	case e.isCollLeft():
		return engine.CollLeft

	case e.isCollRight():
		return engine.CollRight

	default:
		return engine.CollNone
	}
}

func (e *CanvasEngine) isCollP1() bool {
	x := e.ballX <= (e.p1X + e.game.P1.Width + magic_p)
	y1 := e.p1Y <= e.ballY
	y2 := (e.p1Y + e.game.P1.Height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *CanvasEngine) isCollP2() bool {
	x := (e.ballX + e.game.Ball.Height) >= e.p2X
	y1 := e.p2Y <= e.ballY
	y2 := (e.p2Y + e.game.P2.Height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *CanvasEngine) isCollTop() bool {
	return e.ballY <= baseline+e.game.Ball.Height+canvas_border_correction
}

func (e *CanvasEngine) isCollBottom() bool {
	return e.ballY+e.game.Ball.Height >= e.game.Height-canvas_border_correction
}

func (e *CanvasEngine) isCollLeft() bool {
	return e.ballX-e.game.Ball.Height-canvas_border_correction <= 0
}

func (e *CanvasEngine) isCollRight() bool {
	return e.ballX+e.game.Ball.Height+canvas_border_correction >= e.game.Width
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

func (e *CanvasEngine) reset() *CanvasEngine {
	e.err = nil
	return e.resetBall().resetPlayers()
}

func (e *CanvasEngine) resetBall() *CanvasEngine {
	// Center ball
	e.ballX = e.game.Width / 2.0
	e.ballY = e.game.Height / 2.0
	// Random direction
	if rand.Intn(10) < 5 {
		e.ballXVelocity = -default_ball_x_vel_ratio * e.game.Width
		y := min_ball_y_vel_ratio*e.game.Height + rand.Float64()*((max_y_vel_ratio*e.game.Height)-(min_ball_y_vel_ratio*e.game.Height))
		e.ballYVelocity = -y
	} else {
		e.ballXVelocity = default_ball_x_vel_ratio * e.game.Width
		y := min_ball_y_vel_ratio*e.game.Height + rand.Float64()*((max_y_vel_ratio*e.game.Height)-(min_ball_y_vel_ratio*e.game.Height))
		e.ballYVelocity = y
	}
	return e
}

func (e *CanvasEngine) resetPlayers() *CanvasEngine {
	// P1
	e.p1X = 0 + default_padding
	e.p1Y = e.game.Height/2 - e.game.P1.Height/2
	// P2
	e.p2X = e.game.Width - +e.game.P1.Width - default_padding
	e.p2Y = e.game.Height/2 - e.game.P2.Height/2
	return e
}

func (e *CanvasEngine) advance() *CanvasEngine {
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
		switch y := (e.p2Y + (e.game.P2.Height / 2)) - e.ballY; {
		case y > 0:
			e.p2YVelocity = max_y_vel_ratio * e.game.Height
			e.p2Y -= e.p2YVelocity / e.fps
		case y < 0:
			e.p2YVelocity = max_y_vel_ratio * e.game.Height
			e.p2Y += e.p2YVelocity / e.fps
		case y > -0.9 && y < 0.9:
			e.p2YVelocity = 0
		}
	}

	return e
}

func (e *CanvasEngine) p1Up() *CanvasEngine {
	e.p1YVelocity = player_input_dist
	e.p1Y += player_input_dist
	e.p1YVelocity = 0
	return e
}

func (e *CanvasEngine) p1Down() *CanvasEngine {
	e.p1YVelocity = player_input_dist
	e.p1Y -= player_input_dist
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
	if e.p1Y+e.game.P1.Height >= e.game.Height-default_padding {
		e.p1Y = e.game.Height - e.game.P1.Height - default_padding
		e.p1YVelocity = 0
	}
	// P2, top
	if e.p2Y-default_padding <= baseline {
		e.p2Y = baseline + default_padding
		e.p2YVelocity = 0
	}
	// P2, bottom
	if e.p2Y+e.game.P2.Height >= e.game.Height-default_padding {
		e.p2Y = e.game.Height - e.game.P2.Height - default_padding
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
	if e.ballY+e.game.Ball.Height >= e.game.Height {
		e.ballY = e.game.Height - e.game.Ball.Height - 1
	}
	// P1
	if e.ballX-e.game.Ball.Width <= e.p1X+e.game.P1.Width {
		e.ballX = e.p1X + e.game.P1.Width + magic_p
	}
	// P2
	if e.ballX+e.game.Ball.Width >= e.p2X {
		e.ballX = e.p2X - magic_p
	}
	return e
}
