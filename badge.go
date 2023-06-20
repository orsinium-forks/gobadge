package main

import (
	"image/color"
	"machine"
	"time"

	qrcode "github.com/skip2/go-qrcode"
	"github.com/tinygo-org/gobadge/pybadge"
)

const (
	WIDTH  = 160
	HEIGHT = 128
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
)

const (
	logoDisplayTime = 10 * time.Second
)

func QR(d *pybadge.Device, msg string) {
	qr, err := qrcode.New(msg, qrcode.Medium)
	if err != nil {
		println(err, 123)
	}

	qrbytes := qr.Bitmap()
	size := int16(len(qrbytes))

	factor := int16(HEIGHT / len(qrbytes))

	bx := (WIDTH - size*factor) / 2
	by := (HEIGHT - size*factor) / 2
	d.Display.FillScreen(color.RGBA{109, 0, 140, 255})
	for y := int16(0); y < size; y++ {
		for x := int16(0); x < size; x++ {
			if qrbytes[y][x] {
				d.Display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, black)
			} else {
				d.Display.FillRectangle(bx+x*factor, by+y*factor, factor, factor, white)
			}
		}
	}

	time.Sleep(logoDisplayTime)
	pressed, _ := d.Buttons.Read8Input()
	if pressed&machine.BUTTON_SELECT_MASK > 0 {
		return
	}
}
