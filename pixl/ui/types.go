package ui

import (
	"fyne.io/fyne/v2"
	"github.com/davoodMossGreen/Pixl/apptype"
	"github.com/davoodMossGreen/Pixl/pxcanvas"
	"github.com/davoodMossGreen/Pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
