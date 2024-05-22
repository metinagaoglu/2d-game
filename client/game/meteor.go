package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	. "github.com/metinagaoglu/2d-game/assets"
	"math"
	"math/rand"
	"fmt"
)

const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
)

type Meteor struct {
	position Vector
	rotation float64
	sprite   *ebiten.Image
	movement Vector
	rotationSpeed float64
}

func NewMeteor(baseVelocity float64) *Meteor {
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

	// Randomized velocity
	velocity := baseVelocity + rand.Float64() * 1.5

	// Direction is the target minus the current position
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}
	// Normalize the direction
	normalizedDirection := direction.Normalize()
	fmt.Println(normalizedDirection)


	// Multiply the direction by velocity
	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Meteor{
		position:pos,
		sprite: sprite,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		movement: movement,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.movement.Y
	m.position.X += m.movement.X
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.sprite.Bounds()

	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}