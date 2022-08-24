// Package gui is responsible for visual output
// File fyneio.go implements fyne.io methods for gui interface
package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	//magnification is how magnified the circles should be
	magnification    = 30
	rowColStartIndex = 0
)

var (
	vapeOnColour  = color.RGBA{0, 0, 255, 100}
	vapeOffColour = color.RGBA{0, 0, 0, 100}
)

//Canvas represents a canvas in the fyne.io GUI
//go:generate go run github.com/vektra/mockery/cmd/mockery -name Canvas -inpkg --filename canvas_mock.go
type Canvas interface {
	NewCircle(color.Color)
	VapeLightOff(int)
	BrailleToScreen(string)
}

type fyneScreen struct {
	window fyne.Window
	vapes  []*canvas.Circle
}

//NewFyneGui returns a window and screen
func NewFyneGui() Screen {
	vapeAwareNess := app.New()
	vapeDisplay := vapeAwareNess.NewWindow("Vape Display")

	circlesCanvas, cirlesCircles := createCanvasAndDots()

	grid := container.NewAdaptiveGrid(Columns, circlesCanvas...)
	vapeDisplay.SetContent(grid)
	vapeDisplay.Resize(fyne.NewSize(float32(Columns)*magnification, float32(Rows)*magnification))
	vapeDisplay.SetFixedSize(true)

	return &fyneScreen{
		window: vapeDisplay,
		vapes:  cirlesCircles,
	}
}

func createCanvasAndDots() ([]fyne.CanvasObject, []*canvas.Circle) {
	circlesCanvas := []fyne.CanvasObject{}
	cirlesCircles := []*canvas.Circle{}
	for r := rowColStartIndex; r < Rows; r++ {
		for c := rowColStartIndex; c < Columns; c++ {
			newVape := canvas.NewCircle(color.Black)
			circlesCanvas = append(circlesCanvas, newVape)
			cirlesCircles = append(cirlesCircles, newVape)
		}
	}
	return circlesCanvas, cirlesCircles
}

func (s *fyneScreen) ShowAndRun() {
	s.window.ShowAndRun()
}

func (s *fyneScreen) VapeLightOn(i int) error {
	if i < len(s.vapes) {
		s.vapes[i].FillColor = vapeOnColour
		canvas.Refresh(s.vapes[i])
		return nil
	}
	return fmt.Errorf("vapeID=%d is greater than len(s.vapes)=%d", i, len(s.vapes))
}

func (s *fyneScreen) VapeLightOff(i int) error {
	if i < len(s.vapes) {
		s.vapes[i].FillColor = vapeOffColour
		canvas.Refresh(s.vapes[i])
		return nil
	}
	return fmt.Errorf("vapeID=%d is greater than len(s.vapes)=%d", i, len(s.vapes))
}

func (s *fyneScreen) AllVapesOff() error {
	return allVapesOff(s, s.vapes)
}

func allVapesOff(s Screen, vapes []*canvas.Circle) error {
	for i := range vapes {
		if err := s.VapeLightOff(i); err != nil {
			return err
		}
	}
	return nil
}

func (s *fyneScreen) DisplayMatrix(matrix [][]int) error {
	return DisplayMatrix(s, matrix)
}

func (*fyneScreen) NewRow() error {
	return nil
}
