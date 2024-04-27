package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/metinagaoglu/2d-game/assets"
	"log"
)

type Player struct {
	Position Vector
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
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
		Position: pos,
		sprite:   PlayerSprite,
	}
}

func (p *Player) Update() {
	speed := 5.0

	// Ful screen
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		panic("Game Over") // TODO: fix this
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Position.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Position.X += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Position.Y -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Position.Y += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		log.Println("X:", p.Position.X, "Y:", p.Position.Y)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.sprite, op)
}
