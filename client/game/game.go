package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
	"fmt"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	baseMeteorVelocity = 0.25
	meteorSpeedUpAmount = 0.1
	meteorSpeedUpTime = 5 * time.Second
)

type Game struct {
	player      *Player
	attackTimer *Timer
	meteors    []*Meteor
	meteorTimer *Timer
	bullets []*Bullet

	baseVelocity  float64
	velocityTimer *Timer
}

func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}

	// Update the attack timer
	g.player.Update()
	g.attackTimer.Update()
	g.meteorTimer.Update()

	if g.meteorTimer.IsReady() {
		g.meteorTimer.Reset()
	
		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, meteor := range g.meteors {
		meteor.Update()
	}

	for _, bullet := range g.bullets {
		bullet.Update()
	}

	// Collision detection
	for i, meteor := range g.meteors {
		for j, bullet := range g.bullets {
			//fmt.Println(meteor.Collider())
			//fmt.Println(bullet.Collider())
			if meteor.Collider().Intersects(bullet.Collider()) {
				fmt.Println("COLLISION")
				fmt.Println("i:", i, "j:", j)
				fmt.Println("BAMMMM")
			}
		}
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
		baseVelocity: baseMeteorVelocity,
		velocityTimer: NewTimer(meteorSpeedUpTime),
	}

	g.player = NewPlayer(g)

	return g
}
