package main

import . "github.com/gen2brain/raylib-go/raylib"

type paddle struct {
	Rectangle
	game *game
	name string
	score int32
	color Color
	speed float32
}

func(p *paddle) physicsBody() []Vector2 {
	pb := make([]Vector2, 0)

	for y := p.Y; y <= p.Y+p.Height; y++ {
		for x := p.X; x <= p.X+p.Width; x++ {
			pb = append(pb, NewVector2(x, y))
		}
	}

	return pb
}



func (p *paddle) isHit(pos Vector2) (bool, string) {

	for _, vec := range p.physicsBody() {
		if vec.X == pos.X && vec.Y == pos.Y {
			middle := p.Y+p.Height/2
			middleTop := middle-p.Height/7
			middleBottom := middle+p.Height/7
			if vec.Y < middleBottom && vec.Y > middleTop {
				return true, "straight"
			}else if vec.Y < middleTop {
				return true, "up"
			}else if vec.Y > middleBottom {
				return true, "down"
			}

		}
	}
	return false, ""
}

func(p *paddle) update(){
	switch(p.name){
	case "player1":
		if IsKeyDown(KeyUp) {
			if p.Y > 0 {
				p.Y -= p.speed
			}
		}
		if IsKeyDown(KeyDown) {
			if p.Y+p.Height < float32(p.game.options.screenHeight) {
				p.Y += p.speed
			}
		}
		break
	case "computer":
		if p.game.ball.center.Y > p.Y+p.Height/2 {
			p.Y += p.speed
		}else if p.game.ball.center.Y < p.Y+p.Height/2 {
			p.Y -= p.speed
		}
		break
	}
}

func (p *paddle) addPoint(){
	p.score++
}

func (p *paddle) draw(){
	DrawRectangle(int32(p.X), int32(p.Y), int32(p.Width), int32(p.Height), p.color)
}

func newPaddle(n string, g *game)*paddle{

	var p *paddle

	if n == "player1" {
		p = &paddle{
			Rectangle{
				float32(g.options.screenWidth-45),
				float32(g.options.screenHeight/2 - 60),
				35,
				200,
			},
			g,
			n,
			0,
			White,
			5.0,
		}
	}else if n == "computer" {
		p = &paddle{
			Rectangle{
				10,
				float32(g.options.screenHeight/2 - 60),
				35,
				200,
			},
			g,
			n,
			0,
			White,
			3.0,
		}
	}


	return p
}