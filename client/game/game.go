package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	player      *Player
	attackTimer *Timer
}

func (g *Game) Update() error {
	// Update the attack timer
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	g := &Game{
		player:      NewPlayer(),
		attackTimer: NewTimer(1000),
	}
	return g
}
