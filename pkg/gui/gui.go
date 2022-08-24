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
	//Rows is the amount of rows in the display
	Rows = 5
	//Columns is the amount of columns in the display
	Columns = 30
	//VapeOn is the value for when a vape light is on
	VapeOn = 1
	//VapeOff is the value for when a vape light is off
	VapeOff = 0

	//magnification is how magnified the circles should be
	magnification    = 30
	rowColStartIndex = 0
)

var (
	vapeOnColour  = color.RGBA{0, 0, 255, 100}
	vapeOffColour = color.RGBA{0, 0, 0, 100}
)

//Screen represents a screen or hardware visual output
//go: generate go run github.com/vektra/mockery/cmd/mockery -name Screen -inpkg --filename screen_mock.go
type Screen interface {
	VapeLightOn(int) error
	VapeLightOff(int) error
	DisplayLetter([][]int) error
	AllVapesOff() error
}

//Canvas represents a canvas in the fyne.io GUI
//go: generate go run github.com/vektra/mockery/cmd/mockery -name Screen -inpkg --filename screen_mock.go
type Canvas interface {
	NewCircle(color.Color)
	VapeLightOff(int)
	BrailleToScreen(string)
}

type screen struct {
	vapes []*canvas.Circle
}

//NewGui returns a window and screen
func NewGui() (fyne.Window, Screen) {
	vapeAwareNess := app.New()
	vapeDisplay := vapeAwareNess.NewWindow("Vape Display")

	circlesCanvas, cirlesCircles := createCanvasAndDots()

	grid := container.NewAdaptiveGrid(Columns, circlesCanvas...)
	vapeDisplay.SetContent(grid)
	vapeDisplay.Resize(fyne.NewSize(float32(Columns)*magnification, float32(Rows)*magnification))
	vapeDisplay.SetFixedSize(true)

	return vapeDisplay, &screen{
		vapes: cirlesCircles,
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

func (s *screen) VapeLightOn(i int) error {
	if i < len(s.vapes) {
		s.vapes[i].FillColor = vapeOnColour
		canvas.Refresh(s.vapes[i])
		return nil
	}
	return fmt.Errorf("vapeID=%d is greater than len(s.vapes)=%d", i, len(s.vapes))
}

func (s *screen) VapeLightOff(i int) error {
	if i < len(s.vapes) {
		s.vapes[i].FillColor = vapeOffColour
		canvas.Refresh(s.vapes[i])
		return nil
	}
	return fmt.Errorf("vapeID=%d is greater than len(s.vapes)=%d", i, len(s.vapes))
}

func (s *screen) AllVapesOff() error {
	for i := range s.vapes {
		if err := s.VapeLightOff(i); err != nil {
			return err
		}
	}
	return nil
}

func (s *screen) DisplayLetter(matrix [][]int) error {
	count := rowColStartIndex
	for _, row := range matrix {
		for _, col := range row {
			if col == VapeOn {
				if err := s.VapeLightOn(count); err != nil {
					return err
				}
			}
			if col == VapeOff {
				if err := s.VapeLightOff(count); err != nil {
					return err
				}
			}
			count++
		}
	}
	return nil
}
