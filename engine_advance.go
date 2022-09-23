package main

import "errors"

var (
	ErrP1Win = errors.New("p1 win")
	ErrP2Win = errors.New("p2 win")
)

func (e *canvasEngine) advanceGame() error {
	switch e.detectCollision() {

	case collisionP1Ceiling, collisionP1Ground, collisionP2Ceiling, collisionP2Ground:
		e.inverseBallVelXY().fixBall()

	case collisionP1:
		e.inverseBallVelX().fixBall()

	case collisionP2:
		e.inverseBallVelX().fixBall()

	case collisionGroundLeftWall, collisionCeilingLeftWall:
		return ErrP2Win

	case collisionGroundRightWall, collisionCeilingRightWall:
		return ErrP1Win

	case collisionCeiling:
		e.inverseBallVelY().fixBall()

	case collisionGround:
		e.inverseBallVelY().fixBall()

	case collisionLeftWall:
		return ErrP2Win

	case collisionRightWall:
		return ErrP1Win

	case noCollision:
		// Continue
	}

	e.advanceBall().advancePlayers().fixPlayers()

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
		e.p1Y += e.ballYVelocity / e.fps
		e.p1YVelocity = e.ballYVelocity

	case e.ballDirP2():
		e.p2Y += e.ballYVelocity / e.fps
		e.p2YVelocity = e.ballYVelocity

	}
	return e
}
