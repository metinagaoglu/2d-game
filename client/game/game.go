package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/metinagaoglu/2d-game/assets"
	"image/color"
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

	score int
	hot   float32
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

	g.hot -= 0.2
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
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
				fmt.Println("COLLISION")
				fmt.Println("i:", i, "j:", j)
				fmt.Println("BAMMMM")
			}
		}
	}

	for _, meteor := range g.meteors {
		if meteor.Collider().Intersects(g.player.Collider()) {
			fmt.Println("PLAYER HIT")
			g.Reset()
			break
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

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)
	text.Draw(screen, fmt.Sprintf("\n%.2f", g.hot), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) AddBullet(b *Bullet) {
	if g.hot >= 100 {
		return
	}
	g.bullets = append(g.bullets, b)
	g.hot += 5
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.score = 0
	// Send HTTP Request to save the score
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
