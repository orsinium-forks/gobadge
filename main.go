package main

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/gobadge/badge"
	"github.com/tinygo-org/gobadge/launcher"
	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/drivers/lis3dh"

	"tinygo.org/x/drivers/ws2812"

	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/drivers/st7735"
)

var display st7735.Device
var buttons shifter.Device
var leds ws2812.Device
var bzrPin machine.Pin
var accel lis3dh.Device
var snakeGame = Game{
	colors: []color.RGBA{
		color.RGBA{0, 0, 0, 255},
		color.RGBA{0, 200, 0, 255},
		color.RGBA{250, 0, 0, 255},
		color.RGBA{160, 160, 160, 255},
	},
	snake: Snake{
		body: [208][2]int16{
			{0, 3},
			{0, 2},
			{0, 1},
		},
		length:    3,
		direction: 3,
	},
	appleX: -1,
	appleY: -1,
	status: START,
}

func main() {
	device := pybadge.NewDevice()
	display = *device.Display
	buttons = *device.Buttons
	leds = *device.LEDs
	bzrPin = *device.Buzzer
	accel = *device.Accel

	launcher.Tiny{}.Run(
		&device,
		launcher.App{
			Name: "Badge",
			Run:  badge.Run,
		},
		launcher.App{
			Name: "Snake",
			Run:  snakeGame.Start,
		},
		launcher.App{
			Name: "Rainbow LEDs",
			Run:  Leds,
		},
		launcher.App{
			Name: "Accelerometer",
			Run:  Accel3D,
		},
		launcher.App{
			Name: "Music",
			Run:  Music,
		},
	)

}
