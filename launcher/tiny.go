package launcher

import (
	"image/color"
	"time"

	"github.com/tinygo-org/gobadge/pybadge"
	"tinygo.org/x/drivers/shifter"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

type Tiny struct {
}

var _ Menu = Tiny{}

func (m Tiny) Run(d *pybadge.Device, apps ...App) error {
	for {
		index := m.selectApp(d, apps)
		err := apps[index].Run(d)
		if err != nil {
			return err
		}
	}
}

func (m Tiny) selectApp(d *pybadge.Device, apps []App) int16 {
	d.Display.FillScreen(color.RGBA{0, 0, 0, 255})

	bgColor := color.RGBA{0, 40, 70, 255}
	d.Display.FillScreen(bgColor)
	tinydraw.FilledTriangle(d.Display, 0, 128, 0, 45, 45, 0, color.RGBA{255, 255, 255, 255})
	tinydraw.FilledTriangle(d.Display, 45, 0, 0, 128, 145, 0, color.RGBA{255, 255, 255, 255})
	tinydraw.FilledTriangle(d.Display, 0, 128, 15, 128, 145, 0, color.RGBA{255, 255, 255, 255})
	for i := int16(0); i < 8; i++ {
		tinydraw.Line(d.Display, 0, 110+i, 110+i, 0, bgColor)
	}

	selected := int16(0)
	numOpts := int16(len(apps))
	for i, app := range apps {
		i := int16(i)
		tinydraw.Circle(d.Display, 32, 37+10*i, 4, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 39, 39+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 39, 40+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 39, 41+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 40, 41+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 41, 41+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 41, 40+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 41, 39+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 40, 39+10*i, app.Name, color.RGBA{0, 0, 0, 255})
		tinyfont.WriteLine(d.Display, &proggy.TinySZ8pt7b, 40, 40+10*i, app.Name, color.RGBA{250, 250, 0, 255})
	}

	tinydraw.FilledCircle(d.Display, 32, 37, 2, color.RGBA{200, 200, 0, 255})

	released := true
	for {
		pressed, _ := d.Buttons.ReadInput()

		if released && d.Buttons.Pins[shifter.BUTTON_UP].Get() && selected > 0 {
			selected--
			tinydraw.FilledCircle(d.Display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(d.Display, 32, 37+10*(selected+1), 2, color.RGBA{255, 255, 255, 255})
		}
		if released && d.Buttons.Pins[shifter.BUTTON_DOWN].Get() && selected < (numOpts-1) {
			selected++
			tinydraw.FilledCircle(d.Display, 32, 37+10*selected, 2, color.RGBA{200, 200, 0, 255})
			tinydraw.FilledCircle(d.Display, 32, 37+10*(selected-1), 2, color.RGBA{255, 255, 255, 255})
		}
		if released && d.Buttons.Pins[shifter.BUTTON_START].Get() {
			break
		}
		if pressed == 0 {
			released = true
		} else {
			released = false
		}
		time.Sleep(200 * time.Millisecond)
	}
	return selected
}
