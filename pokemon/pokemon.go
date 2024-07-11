package pokemon

import (
	_ "embed"
	"image/color"
	"machine"
	"math/rand"

	"tinygo.org/x/tinyfont/freesans"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
)

const (
	spriteWidth  = 128
	spriteHeight = 120
)

var black = color.RGBA{1, 1, 1, 255}

//go:embed assets/charizard.bin
var spriteCharlizard []uint8

//go:embed assets/pikachu.bin
var spritePikachu []uint8

//go:embed assets/bulbasaur.bin
var spriteBulbasaur []uint8

//go:embed assets/splash.bin
var splash []uint8

type Pokemon struct {
	name     string
	sprite   []uint8
	strength string
	weakness string
}

func (p *Pokemon) Draw(display uc8151.Device, x, y int16) {
	display.DrawBuffer(x, y, spriteWidth, spriteHeight, p.sprite)
}

func (p *Pokemon) WinsAgainst(p2 *Pokemon) bool {
	return p.strength == p2.weakness
}

var Pokedex = []Pokemon{
	{"Char Lizard", spriteCharlizard, "fire", "water"},
	{"Lightning Rat", spritePikachu, "electric", "fire"},
	{"Water Turtle", spriteBulbasaur, "water", "electric"},
}

type SelectionModel struct {
	display  uc8151.Device
	selected int
}

func (m *SelectionModel) Draw() {
	m.display.ClearBuffer()
	Pokedex[m.selected].Draw(m.display, 0, 0)
	m.display.Display()
	m.display.WaitUntilIdle()
}

func NewSelectionModel(display uc8151.Device) SelectionModel {
	return SelectionModel{display, -1}
}

func Go(display uc8151.Device) {
	m := NewSelectionModel(display)

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
			goto confirm // Confirm selection
		case machine.BUTTON_DOWN.Get():
			return // Exit to menu
		}
	}

confirm:
	player := Pokedex[m.selected]
	opponent := Pokedex[rand.Intn(len(Pokedex))]

	m.display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, player.name, black)
	m.display.Display()
	m.display.WaitUntilIdle()

	m.display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, "VS", black)
	m.display.Display()
	m.display.WaitUntilIdle()

	m.display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, opponent.name, black)
	m.display.Display()
	m.display.WaitUntilIdle()
}


func Logo(display uc8151.Device) {
	display.ClearBuffer()
	display.DrawBuffer(0, 25, 128, 246, []uint8(splash))
	display.Display()
	display.WaitUntilIdle()
}
