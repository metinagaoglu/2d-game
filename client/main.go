package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/metinagaoglu/2d-game/assets"
	. "github.com/metinagaoglu/2d-game/game"
	"time"
)

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	playerPosition Vector
	attackTimer    *Timer
}

func (g *Game) Update() error {
	// Update the attack timer
	g.attackTimer.Update()

	if g.attackTimer.IsReady() {
		fmt.Println("Attack ready")
		g.attackTimer.Reset()
	}

	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerPosition.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerPosition.X += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerPosition.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerPosition.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return fmt.Errorf("Game closed")
	}

	// Handle sapce for debug
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("Space pressed")
		fmt.Println("X: ", g.playerPosition.X, "Y: ", g.playerPosition.Y)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	// White
	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerPosition: Vector{X: 100, Y: 300},
		attackTimer:    NewTimer(5 * time.Second),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}

}
