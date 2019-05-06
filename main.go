package main

import (
	"fmt"
	. "github.com/gen2brain/raylib-go/raylib"
)

func main () {

	game := newGame()

	for !WindowShouldClose() {
		game.play()
	}

	CloseWindow()
}

// GAME
type options struct {
	screenWidth int32
	screenHeight int32
	title string
	fps int32
	paused bool
}

type game struct {
	ball *ball
	paddle1 *paddle
	paddle2 *paddle
	*options

}

func newGame()*game{
	g := &game{
		options:&options{
			800,
			600,
			"go pong",
			60,
			false,
		},
	}

	g.ball = newBall(g)
	g.paddle1 = newPaddle("player1", g)
	g.paddle2 = newPaddle("computer", g)


	InitWindow(g.screenWidth, g.screenHeight, g.title)
	SetTargetFPS(g.fps)


	return g
}

func (g *game) update(){
	g.ball.update()
	g.paddle1.update()
	g.paddle2.update()
}

func (g *game) drawScore(){

	scoreString := fmt.Sprintf("%v / %v", g.paddle2.score, g.paddle1.score)
	DrawText(scoreString, g.options.screenWidth/2-10, 30, 30, White)
}

func (g *game) draw(){
	g.drawScore()
	g.ball.draw()
	g.paddle1.draw()
	g.paddle2.draw()
}

func(g *game) play(){
	BeginDrawing()

	ClearBackground(Black)

	g.update()
	g.draw()

	EndDrawing()
}