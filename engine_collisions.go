package main

type collision int

const (
	noCollision collision = 0

	collisionGroundLeftWall collision = iota
	collisionCeilingLeftWall
	collisionGroundRightWall
	collisionCeilingRightWall

	collisionP1Ground
	collisionP1Ceiling
	collisionP2Ground
	collisionP2Ceiling
	collisionP1
	collisionP2

	collisionGround
	collisionCeiling

	collisionLeftWall
	collisionRightWall
)

func (e *canvasEngine) detectCollision() collision {
	switch {
	case e.isCollisionGroundLeftWall():
		engineLogger.Println("collision ball ground left wall")
		return collisionGroundLeftWall

	case e.isCollisionCeilingLeftWall():
		engineLogger.Println("collision ball ceiling left wall")
		return collisionCeilingLeftWall

	case e.isCollisionGroundRightWall():
		engineLogger.Println("collision ball ground right wall")
		return collisionGroundLeftWall

	case e.isCollisionCeilingRightWall():
		engineLogger.Println("collision ball ceiling right wall")
		return collisionCeilingRightWall

	case e.isCollisionP1Ground():
		engineLogger.Println("collision ball p1 ground")
		return collisionP1Ground

	case e.isCollisionP1Ceiling():
		engineLogger.Println("collision ball p1 ceiling")
		return collisionP1Ceiling

	case e.isCollisionP2Ground():
		engineLogger.Println("collision ball p2 ground")
		return collisionP2Ground

	case e.isCollisionP2Ceiling():
		engineLogger.Println("collision ball p2 ceiling")
		return collisionP2Ceiling

	case e.isCollisionP1():
		engineLogger.Println("collision ball p1")
		return collisionP1

	case e.isCollisionP2():
		engineLogger.Println("collision ball p2")
		return collisionP2

	case e.isCollisionGround():
		engineLogger.Println("collision ball ground")
		return collisionGround

	case e.isCollisionCeiling():
		engineLogger.Println("collision ball ceiling")
		return collisionCeiling

	case e.isCollisionLeftWall():
		engineLogger.Println("collision left wall")
		return collisionLeftWall

	case e.isCollisionRightWall():
		engineLogger.Println("collision right wall")
		return collisionRightWall

	default:
		engineLogger.Println("no collision")
		return noCollision
	}
}

func (e *canvasEngine) isCollisionP1() bool {
	x := e.ballX <= (e.p1X + e.game.p1.width)
	y1 := e.p1Y <= e.ballY
	y2 := (e.p1Y + e.game.p1.height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *canvasEngine) isCollisionP2() bool {
	x := (e.ballX + e.game.ball.height) >= e.p2X
	y1 := e.p2Y <= e.ballY
	y2 := (e.p2Y + e.game.p2.height) >= e.ballY
	y := y1 && y2
	return x && y
}

func (e *canvasEngine) isCollisionCeiling() bool {
	return e.ballY <= baseline+e.game.ball.height
}

func (e *canvasEngine) isCollisionGround() bool {
	return e.ballY+e.game.ball.height >= e.game.height
}

func (e *canvasEngine) isCollisionLeftWall() bool {
	return e.ballX-e.game.ball.height-canvas_border_correction <= 0
}

func (e *canvasEngine) isCollisionRightWall() bool {
	return e.ballX+e.game.ball.height+canvas_border_correction >= e.game.width
}

func (e *canvasEngine) isCollisionP1Ceiling() bool {
	return e.isCollisionP1() && e.isCollisionCeiling()
}

func (e *canvasEngine) isCollisionP2Ceiling() bool {
	return e.isCollisionP2() && e.isCollisionCeiling()
}

func (e *canvasEngine) isCollisionP1Ground() bool {
	return e.isCollisionP1() && e.isCollisionGround()
}

func (e *canvasEngine) isCollisionP2Ground() bool {
	return e.isCollisionP2() && e.isCollisionGround()
}

func (e *canvasEngine) isCollisionCeilingLeftWall() bool {
	return e.isCollisionCeiling() && e.isCollisionLeftWall()
}

func (e *canvasEngine) isCollisionGroundLeftWall() bool {
	return e.isCollisionGround() && e.isCollisionLeftWall()
}

func (e *canvasEngine) isCollisionCeilingRightWall() bool {
	return e.isCollisionCeiling() && e.isCollisionRightWall()
}

func (e *canvasEngine) isCollisionGroundRightWall() bool {
	return e.isCollisionGround() && e.isCollisionRightWall()
}
