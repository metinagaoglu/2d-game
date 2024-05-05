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
	meteors    []*Meteor
	meteorTimer *Timer
	bullets []*Bullet
}

func (g *Game) Update() error {
	// Update the attack timer
	g.player.Update()
	g.attackTimer.Update()
	g.meteorTimer.Update()

	if g.meteorTimer.IsReady() {
		g.meteorTimer.Reset()
	
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, meteor := range g.meteors {
		meteor.Update()
	}

	for _, bullet := range g.bullets {
		bullet.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, meteor := range g.meteors {
		meteor.Draw(screen)
	}

	for _, bullet := range g.bullets {
		bullet.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}

func NewGame() *Game {
	g := &Game{
		attackTimer: NewTimer(1000),
		meteors:    []*Meteor{},
		meteorTimer: NewTimer(3000),
	}

	g.player = NewPlayer(g)

	return g
}
