package main

import (
	_ "embed"
	"fmt"
	"machine"
	"math/rand"
	"runtime"

	"tinygo.org/x/tinyfont/freesans"

	"github.com/conejoninja/badger2040/bt"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
)

const (
	spriteWidth  = 128
	spriteHeight = 120
)

//go:embed assets/charizard.bin
var spriteCharizard []uint8

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
	{name: "Char Lizard", sprite: spriteCharizard, strength: "fire", weakness: "water"},
	{name: "Water Turtle", sprite: spriteBulbasaur, strength: "water", weakness: "electric"},
	{name: "Lightning Rat", sprite: spritePikachu, strength: "electric", weakness: "fire"},
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
	display.ClearBuffer()

	line := "Choose Your"
	w32, _ := tinyfont.LineWidth(&freesans.Bold24pt7b, line)
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 176-int16(w32)/2, 60, line, black)

	line = "Pokemon!"
	w32, _ = tinyfont.LineWidth(&freesans.Bold24pt7b, line)
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 176-int16(w32)/2, 88, line, black)

	display.Display()
	display.WaitUntilIdle()

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
			bt.Advertise()
			// return // Exit to menu
		default:
			runtime.Gosched()
		}
	}

confirm:
	fmt.Println("confirmed")
	player := Pokedex[m.selected]
	opponent := Pokedex[rand.Intn(len(Pokedex))]

	display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, player.name, black)
	display.Display()
	display.WaitUntilIdle()

	display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, "VS", black)
	display.Display()
	display.WaitUntilIdle()

	display.ClearBuffer()
	tinyfont.WriteLine(&display, &freesans.Bold18pt7b, 5, 74, opponent.name, black)
	display.Display()
	display.WaitUntilIdle()

	display.ClearBuffer()
	player.Draw(display, 120, 0)
	opponent.Draw(display, 0, 0)
	display.Display()
	display.WaitUntilIdle()

	winner := opponent
	loser := player
	blankingOffset := int16(120)
	if player.WinsAgainst(&opponent) {
		winner = player
		loser = opponent
		blankingOffset = 0
	}

	flashLoser := func() {
		fmt.Println("loser flashing")
		display.ClearBuffer()
		display.DisplayRect(blankingOffset, 0, spriteWidth, spriteHeight)
		display.WaitUntilIdle()
		loser.Draw(display, blankingOffset, 0)
		display.DisplayRect(blankingOffset, 0, spriteWidth, spriteHeight)
		display.WaitUntilIdle()
	}

	flashLoser()
	flashLoser()
	flashLoser()

	display.ClearBuffer()
	winner.Draw(display, 60, 0)
	display.WaitUntilIdle()
}

func ShowSplash(display uc8151.Device) {
	display.ClearBuffer()
	display.DrawBuffer(0, 25, 128, 246, []uint8(splash))
	display.Display()
	display.WaitUntilIdle()
}
