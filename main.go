package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/tinygo-org/gobadge/badge"
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

	for {
		switch menu() {
		case 0:
			badge.Run(&device)
			break
		case 1:
			snakeGame.Start()
			break
		case 2:
			Leds()
			break
		case 3:
			Accel3D()
			break
		case 4:
			Music()
			break
		default:
			break
		}
		println("LOOP")
		time.Sleep(1 * time.Second)
	}
}

func getRainbowRGB(i uint8) color.RGBA {
	if i < 85 {
		return color.RGBA{i * 3, 255 - i*3, 0, 255}
	} else if i < 170 {
		i -= 85
		return color.RGBA{255 - i*3, 0, i * 3, 255}
	}
	i -= 170
	return color.RGBA{0, i * 3, 255 - i*3, 255}
}
