package slides

import (
	"machine"
	"time"

	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type Scroll struct {
	Device *pybadge.Device
	Line1  string
	Line2  string
	Line3  string
}

func (s *Scroll) Run() error {
	s.Device.Display.FillScreen(white)

	// calculate the width of the text, so we could center them later
	w32top, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, s.Line1)
	w32middle, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, s.Line2)
	w32bottom, _ := tinyfont.LineWidth(&freesans.Bold12pt7b, s.Line3)
	tinyfont.WriteLine(
		s.Device.Display,
		&freesans.Bold12pt7b,
		(width-int16(w32top))/2,
		34,
		s.Line1,
		getRainbowRGB(200),
	)
	tinyfont.WriteLine(
		s.Device.Display,
		&freesans.Bold12pt7b,
		(width-int16(w32middle))/2,
		60,
		s.Line2,
		getRainbowRGB(80),
	)
	tinyfont.WriteLine(
		s.Device.Display,
		&freesans.Bold12pt7b,
		(width-int16(w32bottom))/2,
		100,
		s.Line3,
		getRainbowRGB(120),
	)

	s.Device.Display.SetScrollArea(0, 0)
	for k := 0; k < 4; k++ {
		for i := int16(159); i >= 0; i-- {
			pressed, _ := s.Device.Buttons.Read8Input()
			if pressed&machine.BUTTON_SELECT_MASK > 0 {
				return nil
			}
			s.Device.Display.SetScroll(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
	s.Device.Display.SetScroll(0)
	s.Device.Display.StopScroll()
	return nil
}
