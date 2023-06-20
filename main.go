package main

import (
	"github.com/tinygo-org/gobadge/badge"
	"github.com/tinygo-org/gobadge/launcher"
	"github.com/tinygo-org/gobadge/pybadge"
)

func main() {
	device := pybadge.NewDevice()

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
