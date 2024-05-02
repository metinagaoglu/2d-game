package game 

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	. "github.com/metinagaoglu/2d-game/assets"
)


type Meteor struct {
	position Vector
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	return &Meteor{
		position: Vector{
			X: rand.Float64() * ScreenWidth,
			Y: -50,
		},
		sprite:   sprite,
	}
}

func (m *Meteor) Update() {
	m.position.Y += 0.5
	m.position.X += 0.1
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}