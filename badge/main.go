package badge

import (
	"fmt"
	"image/color"

	"github.com/tinygo-org/gobadge/pybadge"
)

const (
	width  = 160
	height = 128
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
)

type Screen interface {
	Run() error
}

func Run(d *pybadge.Device) error {
	screens := []Screen{
		ScreenName{Device: d, Name: "gram"},
	}
	for _, s := range screens {
		err := s.Run()
		if err != nil {
			return fmt.Errorf("run screen: %v", err)
		}
	}
	return nil
}
