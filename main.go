package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
)

var display uc8151.Device

var (
	black = color.RGBA{1, 1, 1, 255}
	white = color.RGBA{0, 0, 0, 255}
)

const (
	WIDTH  = 296
	HEIGHT = 128
)

func main() {
	fmt.Println("booting up")
	led3v3 := machine.ENABLE_3V3
	led3v3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3v3.High()

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(
		machine.SPI0,
		machine.EPD_CS_PIN,
		machine.EPD_DC_PIN,
		machine.EPD_RESET_PIN,
		machine.EPD_BUSY_PIN,
	)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	machine.BUTTON_A.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.BUTTON_B.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.BUTTON_C.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.BUTTON_UP.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	machine.BUTTON_DOWN.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	ShowSplash(display)
	time.Sleep(3 * time.Second)

	Go(display)
}
