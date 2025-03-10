module github.com/metinagaoglu/2d-game

go 1.22.1

replace github.com/metinagaoglu/2d-game/assets => ./assets

replace github.com/metinagaoglu/2d-game/game => ./game

require (
	github.com/hajimehoshi/ebiten/v2 v2.7.2
	github.com/metinagaoglu/2d-game/assets v0.0.0-00010101000000-000000000000
)

require (
	github.com/ebitengine/gomobile v0.0.0-20240329170434-1771503ff0a8 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/purego v0.7.0 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	golang.org/x/image v0.15.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
