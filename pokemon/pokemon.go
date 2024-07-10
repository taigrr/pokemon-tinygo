package pokemon

import (
	"machine"
	_ "embed"

	"tinygo.org/x/drivers/uc8151"
)

const spriteWidth = 120
const spriteHeight = 128

//go:embed assets/charizard.bin
var spriteCharlizard []uint8

//go:embed assets/pikachu.bin
var spritePikachu []uint8

//go:embed assets/bulbasaur.bin
var spriteBulbasaur []uint8

type Pokemon struct {
	name string
	sprite []uint8
	strength string
	weakness string
}

func (p *Pokemon) Draw(display uc8151.Device, x, y int16) {
	display.DrawBuffer(x, y, spriteWidth, spriteHeight, p.sprite)
}

func (p *Pokemon) WinsAgainst(p2 *Pokemon) bool {
	return p.strength == p2.weakness
}

var (
	Pokedex = []Pokemon{
		Pokemon{"Char Lizard", spriteCharlizard, "fire", "water"},
		Pokemon{"Lightning Rat", spritePikachu, "electric", "fire"},
		Pokemon{"Water Turtle", spriteBulbasaur, "water", "electric"},
	}
)

type Model struct {
	display uc8151.Device
	selected int
}

func (m *Model) Draw() {
	m.display.ClearBuffer()
	m.display.Display()
	Pokedex[m.selected].Draw(m.display, 0, 0)
	m.display.WaitUntilIdle()
}

func Go(display uc8151.Device) {
	m := Model{display, -1}

	for {
		switch {
		case machine.BUTTON_A.Get():
			m.selected = 0
			m.Draw()
		case machine.BUTTON_B.Get():
			m.selected = 1
			m.Draw()
		case machine.BUTTON_C.Get():
			m.selected = 2
			m.Draw()
		case machine.BUTTON_UP.Get():
			// TODO
			break
		case machine.BUTTON_DOWN.Get():
			break
		}
	}
}
