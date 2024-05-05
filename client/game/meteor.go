package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/metinagaoglu/2d-game/assets"
	"math"
	"math/rand"
)

type Meteor struct {
	position Vector
	sprite   *ebiten.Image
	movement Vector
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}
	// The distance from the center the meteor should spawn at — half the width
	r := ScreenWidth / 2.0

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	pos := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// TODO: move Direction Logic:

	// Randomized velocity
	velocity := 0.25 + rand.Float64()*1.5

	// Direction is the target minus the current position
	normalizedDirection := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	// Multiply the direction by velocity
	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Meteor{
		position:pos,
		sprite: sprite,
		movement: movement,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.movement.Y
	m.position.X += m.movement.X
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}
