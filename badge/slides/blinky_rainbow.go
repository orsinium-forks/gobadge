package slides

import (
	"machine"

	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type BlinkyRainbow struct {
	Device *pybadge.Device
	Line1  string
	Line2  string
}

func (s *BlinkyRainbow) Run() error {
	s.Device.Display.FillScreen(white)

	// calculate the width of the text so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, s.Line1)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, s.Line2)
	for i := int16(0); i < 20; i++ {
		// show black text
		tinyfont.WriteLine(
			s.Device.Display, &freesans.Bold12pt7b,
			(width-int16(w32top))/2, 50, s.Line1, getRainbowRGB(uint8(i*12)),
		)
		tinyfont.WriteLine(
			s.Device.Display, &freesans.Bold9pt7b,
			(width-int16(w32bottom))/2, 100, s.Line2, getRainbowRGB(uint8(i*12)),
		)

		pressed, _ := s.Device.Buttons.Read8Input()
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			return nil
		}
	}
	return nil
}
