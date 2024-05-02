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
	//	fmt.Println("Meteor: ", meteor)
		meteor.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, meteor := range g.meteors {
		meteor.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	g := &Game{
		player:      NewPlayer(),
		attackTimer: NewTimer(1000),
		meteors:    []*Meteor{},
		meteorTimer: NewTimer(3000),
	}

	return g
}
