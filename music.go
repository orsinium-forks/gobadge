package main

import (
	"image/color"
	"machine"
	"time"

	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func Music(d *pybadge.Device) error {
	white := color.RGBA{255, 255, 255, 255}
	d.Display.FillScreen(white)

	tinyfont.WriteLine(d.Display, &freesans.Bold24pt7b, 0, 50, "MUSIC", color.RGBA{0, 100, 250, 255})
	tinyfont.WriteLine(d.Display, &freesans.Bold9pt7b, 20, 100, "Press any key", color.RGBA{200, 0, 0, 255})

	for {
		pressed, _ := d.Buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			break
		}

		if pressed&machine.BUTTON_START_MASK > 0 {
			tone(d, 5274)
		}
		if pressed&machine.BUTTON_A_MASK > 0 {
			tone(d, 1046)
		}
		if pressed&machine.BUTTON_B_MASK > 0 {
			tone(d, 1975)
		}

		if pressed&machine.BUTTON_LEFT_MASK > 0 {
			tone(d, 329)
		}
		if pressed&machine.BUTTON_RIGHT_MASK > 0 {
			tone(d, 739)
		}
		if pressed&machine.BUTTON_UP_MASK > 0 {
			tone(d, 369)
		}
		if pressed&machine.BUTTON_DOWN_MASK > 0 {
			tone(d, 523)
		}
	}
	return nil
}

func tone(d *pybadge.Device, tone int) {
	for i := 0; i < 10; i++ {
		d.Buzzer.High()
		time.Sleep(time.Duration(tone) * time.Microsecond)

		d.Buzzer.Low()
		time.Sleep(time.Duration(tone) * time.Microsecond)
	}
}
