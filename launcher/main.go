package launcher

import "github.com/tinygo-org/gobadge/pybadge"

type App struct {
	Name string
	Run  func(*pybadge.Device) error
}

type Menu interface {
	Run(*pybadge.Device, ...App) error
}
