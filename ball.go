package main

import (
	. "github.com/gen2brain/raylib-go/raylib"
	"time"
)

type ball struct {
	game *game
	center Vector2
	radius float32
	color Color
	direction string
	trajectory string
	speed float32

}

func newBall(g *game)*ball{
	b := &ball{
		g,
		NewVector2(float32(g.options.screenWidth/2), float32(g.options.screenHeight/2)),
		10.0,
		White,
		"right",
		"straight",
		10.0,
	}


	return b
}

func (b *ball) reset(){
	dir := b.direction
	b.direction = ""
	b.center = NewVector2(float32(b.game.options.screenWidth/2), float32(b.game.options.screenHeight/2))
	time.Sleep(2 * time.Second)
	b.direction = dir
	b.trajectory = "straight"
}

func(b *ball) goRight(){
	switch b.trajectory {
	case "up":
		b.center.X += b.speed
		b.center.Y -= b.speed
		break

	case "straight":
		b.center.X += b.speed
		break

	case "down":
		b.center.X += b.speed
		b.center.Y += b.speed
		break
	}
}

func (b *ball) move(){
	if b.direction == "left" {
		b.goLeft()
	}else if b.direction == "right" {
		b.goRight()
	}
}

func(b *ball) goLeft(){
	switch b.trajectory {
	case "up":
		b.center.X -= b.speed
		b.center.Y -= b.speed/2
		break

	case "straight":
		b.center.X -= b.speed
		break

	case "down":
		b.center.X -= b.speed
		b.center.Y += b.speed/2
		break
	}
}

func(b *ball) changeDirection(){
	if b.direction == "left" {
		b.direction = "right"
	}else if b.direction == "right" {
		b.direction = "left"
	}
}

func (b *ball) checkCollision(){
	if b.center.Y <= 0 {
		b.trajectory = "down"
	}else if b.center.Y >= float32(b.game.options.screenHeight) {
		b.trajectory = "up"
	}

	if b.center.X <= 0 {
		b.game.paddle1.addPoint()
		go b.reset()
	}else if b.center.X >= float32(b.game.options.screenWidth) {
		b.game.paddle2.addPoint()
		go b.reset()
	}

	hit, traj := b.game.paddle1.isHit(b.center)
	if hit{
		b.trajectory = traj
		b.changeDirection()
	}
	hit, traj = b.game.paddle2.isHit(b.center)
	if hit {
		b.trajectory = traj
		b.changeDirection()
	}

}

func (b *ball) update(){
	b.checkCollision()
	b.move()
}

func(b *ball)draw(){
	DrawCircleV(b.center, b.radius, b.color)
}

