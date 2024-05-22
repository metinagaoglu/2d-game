package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	meteorSpawnTime = 1 * time.Second

	baseMeteorVelocity  = 0.50
	meteorSpeedUpAmount = 0.05
	meteorSpeedUpTime   = 1 * time.Second
)

type Game struct {
	player      *Player
	attackTimer *Timer
	meteors     []*Meteor
	meteorTimer *Timer
	bullets     []*Bullet

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
	g.meteorTimer.Update()

	if g.meteorTimer.IsReady() {
		fmt.Println("METEOR SPAWN")
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
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
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
		attackTimer:   NewTimer(1000),
		meteors:       []*Meteor{},
		meteorTimer:   NewTimer(meteorSpawnTime),
		baseVelocity:  baseMeteorVelocity,
		velocityTimer: NewTimer(meteorSpeedUpTime),
	}

	g.player = NewPlayer(g)

	return g
}
