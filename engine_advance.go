package main

import (
	"errors"
)

var (
	ErrP1Win = errors.New("p1 win")
	ErrP2Win = errors.New("p2 win")
)

func (e *canvasEngine) advanceGame() error {
	switch e.detectCollision() {

	case collisionP1Ceiling, collisionP1Ground, collisionP2Ceiling, collisionP2Ground:
		e.inverseBallXYVel().deglitchBall()

	case collisionP1:
		e.inverseBallXVel().deglitchBall()

	case collisionP2:
		e.inverseBallXVel().deglitchBall()

	case collisionGroundLeftWall, collisionCeilingLeftWall:
		return ErrP2Win

	case collisionGroundRightWall, collisionCeilingRightWall:
		return ErrP1Win

	case collisionCeiling:
		e.inverseBallYVel().deglitchBall()

	case collisionGround:
		e.inverseBallYVel().deglitchBall()

	case collisionLeftWall:
		return ErrP2Win

	case collisionRightWall:
		return ErrP1Win

	case noCollision:
		// Continue
	}

	e.advanceBall().advancePlayers().deglitchPlayers()

	return nil
}

func (e *canvasEngine) advanceBall() *canvasEngine {
	e.ballX += e.ballXVelocity / e.fps
	e.ballY += e.ballYVelocity / e.fps
	return e
}

func (e *canvasEngine) advancePlayers() *canvasEngine {
	switch {
	case e.ballDirP1():
		e.p2YVelocity = 0

		switch y := (e.p1Y + (e.game.p1.height / 2)) - e.ballY; {
		case y > 0:
			// Go up
			e.p1YVelocity = max_y_vel_ratio * e.game.height
			e.p1Y -= e.p1YVelocity / e.fps
		case y < 0:
			// Go down
			e.p1YVelocity = max_y_vel_ratio * e.game.height
			e.p1Y += e.p1YVelocity / e.fps
		case y > -0.9 && y < 0.9:
			// Perfect
			e.p1YVelocity = 0
		}

	case e.ballDirP2():
		e.p1YVelocity = 0

		switch y := (e.p2Y + (e.game.p2.height / 2)) - e.ballY; {
		case y > 0:
			// Go up
			e.p2YVelocity = max_y_vel_ratio * e.game.height
			e.p2Y -= e.p2YVelocity / e.fps
		case y < 0:
			// Go down
			e.p2YVelocity = max_y_vel_ratio * e.game.height
			e.p2Y += e.p2YVelocity / e.fps
		case y > -0.9 && y < 0.9:
			// Perfect
			e.p2YVelocity = 0
		}
	}
	return e
}
