package main

import (
	"math"
	"math/rand"
)

func (e *canvasEngine) resetBall() *canvasEngine {
	// Center
	e.ballX = e.game.width / 2.0
	e.ballY = e.game.height / 2.0
	return e
}

func (e *canvasEngine) randomBall() *canvasEngine {
	// Left or right
	if rand.Intn(10) < 5 {
		e.ballXVelocity = -default_ball_x_vel * e.game.width
		v := min_y_vel*e.game.height + rand.Float64()*((max_y_vel*e.game.height)-(min_y_vel*e.game.height))
		e.ballYVelocity = -v
	} else {
		e.ballXVelocity = default_ball_x_vel * e.game.width
		v := min_y_vel*e.game.height + rand.Float64()*((max_y_vel*e.game.height)-(min_y_vel*e.game.height))
		e.ballYVelocity = v
	}
	return e
}

func (e *canvasEngine) resetPlayers() *canvasEngine {
	// P1
	e.p1X = 0 + default_padding
	e.p1Y = e.game.height/2 - e.game.p1.height/2
	// P2
	e.p2X = e.game.width - +e.game.p1.width - default_padding
	e.p2Y = e.game.height/2 - e.game.p2.height/2
	return e
}

func (e *canvasEngine) fixPlayers() *canvasEngine {
	// Top
	if e.p1Y-default_padding <= baseline {
		e.p1Y = baseline + default_padding
		e.p1YVelocity = 0
	}
	// Bottom
	if e.p1Y+e.game.p1.height >= e.game.height-default_padding {
		e.p1Y = e.game.height - e.game.p1.height - default_padding
		e.p1YVelocity = 0
	}
	// Top
	if e.p2Y-default_padding <= baseline {
		e.p2Y = baseline + default_padding
		e.p2YVelocity = 0
	}
	// Bottom
	if e.p2Y+e.game.p2.height >= e.game.height-default_padding {
		e.p2Y = e.game.height - e.game.p2.height - default_padding
		e.p2YVelocity = 0
	}
	return e
}

func (e *canvasEngine) fixBall() *canvasEngine {
	// Top
	if e.ballY <= baseline {
		e.ballY = baseline - math.SmallestNonzeroFloat64
	}
	// Bottom
	if e.ballY+e.game.ball.height >= e.game.height {
		e.ballY = e.game.height - e.game.ball.height - math.SmallestNonzeroFloat64
	}
	return e
}

func (e *canvasEngine) inverseBallVelXY() *canvasEngine {
	return e.inverseBallVelX().inverseBallVelY()
}

func (e *canvasEngine) inverseBallVelX() *canvasEngine {
	if e.ballXVelocity > 0 {
		engineLogger.Println("inverse ball velocity x")
		e.ballXVelocity = e.ballXVelocity * -1
	} else {
		engineLogger.Println("inverse ball velocity x")
		e.ballXVelocity = math.Abs(e.ballXVelocity)
	}
	return e
}

func (e *canvasEngine) inverseBallVelY() *canvasEngine {
	if e.ballYVelocity > 0 {
		engineLogger.Println("inverse ball velocity y")
		e.ballYVelocity = e.ballYVelocity * -1
	} else {
		engineLogger.Println("inverse ball velocity y")
		e.ballYVelocity = math.Abs(e.ballYVelocity)
	}
	return e
}
