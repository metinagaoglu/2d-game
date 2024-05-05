package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/metinagaoglu/2d-game/assets"
	"log"
	"time"
	"math"
)

const (
	shootCooldown     = time.Millisecond * 100
	rotationPerSecond = math.Pi

	bulletSpawnOffset = 50.0
)

type Player struct {
	game *Game
	Position Vector
	sprite   *ebiten.Image
	rotation float64
	shootCooldown *Timer
}

func NewPlayer(game *Game) *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	log.Println("Bounds:", bounds)
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := Vector{
		X: ScreenWidth/2 - halfW,
		Y: ScreenHeight/2 - halfH,
	}

	return &Player{
		game: game,
		Position: pos,
		sprite:   PlayerSprite,
		rotation: 0,
		shootCooldown: NewTimer(shootCooldown),
	}
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.rotation += speed
	}

	// Ful screen
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		panic("Game Over") // TODO: fix this
	}


	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	p.Position.Y -= speed
	// }

	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	p.Position.Y += speed
	// }

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()

		//TODO: challenge
		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2


		spawnPos := Vector{
			p.Position.X + halfW + math.Sin(p.rotation)*bulletSpawnOffset,
			p.Position.Y + halfH + math.Cos(p.rotation)*-bulletSpawnOffset,
		}
		bullet := NewBullet(spawnPos, p.rotation)
		p.game.AddBullet(bullet)

		log.Println("X:", p.Position.X, "Y:", p.Position.Y)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.Position.X, p.Position.Y)

	screen.DrawImage(p.sprite, op)
}
