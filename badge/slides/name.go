package slides

import (
	"image/color"
	"machine"

	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
)

type Name struct {
	Device *pybadge.Device
	Name   string
}

func (s Name) Run() error {
	s.myNameIs(s.Name)

	rainbow := make([]color.RGBA, 256)
	for i := 0; i < 256; i++ {
		rainbow[i] = getRainbowRGB(uint8(i))
	}

	w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, s.Name)
	for i := 0; i < 230; i++ {
		tinyfont.WriteLineColors(
			s.Device.Display,
			&freesans.Bold9pt7b,
			(width-int16(w32))/2,
			72,
			s.Name,
			rainbow[i:],
		)
		pressed, _ := s.Device.Buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			return nil
		}
	}
	return nil
}

func (s Name) myNameIs(name string) {
	s.Device.Display.FillScreen(white)

	var r int16 = 8

	// black corners detail
	s.Device.Display.FillRectangle(0, 0, r, r, black)
	s.Device.Display.FillRectangle(0, height-r, r, r, black)
	s.Device.Display.FillRectangle(width-r, 0, r, r, black)
	s.Device.Display.FillRectangle(width-r, height-r, r, r, black)

	// round corners
	tinydraw.FilledCircle(s.Device.Display, r, r, r, red)
	tinydraw.FilledCircle(s.Device.Display, width-r-1, r, r, red)
	tinydraw.FilledCircle(s.Device.Display, r, height-r-1, r, red)
	tinydraw.FilledCircle(s.Device.Display, width-r-1, height-r-1, r, red)

	// top band
	s.Device.Display.FillRectangle(r, 0, width-2*r-1, r, red)
	s.Device.Display.FillRectangle(0, r, width, 26, red)

	// bottom band
	s.Device.Display.FillRectangle(r, height-r-1, width-2*r-1, r+1, red)
	s.Device.Display.FillRectangle(0, height-2*r-1, width, r, red)

	// top text : my NAME is
	w32, _ := tinyfont.LineWidth(&freesans.Regular12pt7b, "my NAME is")
	tinyfont.WriteLine(s.Device.Display, &freesans.Regular12pt7b, (width-int16(w32))/2, 24, "my NAME is", white)

	// middle text
	w32, _ = tinyfont.LineWidth(&freesans.Bold9pt7b, name)
	tinyfont.WriteLine(s.Device.Display, &freesans.Bold9pt7b, (width-int16(w32))/2, 72, name, black)

	// gophers
	tinyfont.WriteLineColors(s.Device.Display, &gophers.Regular32pt, width-48, 110, "BE", []color.RGBA{getRainbowRGB(100), getRainbowRGB(200)})
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
