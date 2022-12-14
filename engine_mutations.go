package main

import (
	"math"
	"math/rand"
	"time"
)

func (e *canvasEngine) centerBall() *canvasEngine {
	engineLogger.Println("reset ball ...")

	// Center
	e.ballX = e.game.width / 2.0
	e.ballY = e.game.height / 2.0

	time.Sleep(250 * time.Millisecond)

	return e
}

func (e *canvasEngine) randomBallDir() *canvasEngine {
	engineLogger.Println("random ball ...")

	// Left or right
	if rand.Intn(10) < 5 {
		e.ballXVelocity = -default_ball_x_vel_ratio * e.game.width
		y := min_ball_y_vel_ratio*e.game.height + rand.Float64()*((max_y_vel_ratio*e.game.height)-(min_ball_y_vel_ratio*e.game.height))
		e.ballYVelocity = -y
	} else {
		e.ballXVelocity = default_ball_x_vel_ratio * e.game.width
		y := min_ball_y_vel_ratio*e.game.height + rand.Float64()*((max_y_vel_ratio*e.game.height)-(min_ball_y_vel_ratio*e.game.height))
		e.ballYVelocity = y
	}

	time.Sleep(250 * time.Millisecond)

	return e
}

func (e *canvasEngine) centerPlayers() *canvasEngine {
	engineLogger.Println("reset players ...")

	// P1
	e.p1X = 0 + default_padding
	e.p1Y = e.game.height/2 - e.game.p1.height/2
	// P2
	e.p2X = e.game.width - +e.game.p1.width - default_padding
	e.p2Y = e.game.height/2 - e.game.p2.height/2

	time.Sleep(250 * time.Millisecond)

	return e
}

func (e *canvasEngine) deglitchPlayers() *canvasEngine {
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

func (e *canvasEngine) deglitchBall() *canvasEngine {
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

func (e *canvasEngine) inverseBallXYVel() *canvasEngine {
	return e.inverseBallXVel().inverseBallYVel()
}

func (e *canvasEngine) inverseBallXVel() *canvasEngine {
	engineLogger.Println("inverse ball velocity x")
	if e.ballXVelocity > 0 {
		e.ballXVelocity = e.ballXVelocity * -1
	} else {
		e.ballXVelocity = math.Abs(e.ballXVelocity)
	}
	return e
}

func (e *canvasEngine) inverseBallYVel() *canvasEngine {
	engineLogger.Println("inverse ball velocity y")
	if e.ballYVelocity > 0 {
		e.ballYVelocity = e.ballYVelocity * -1
	} else {
		e.ballYVelocity = math.Abs(e.ballYVelocity)
	}
	return e
}
