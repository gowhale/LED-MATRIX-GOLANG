// Package gui is responsible for visual output
// File gui.go contains infomration used by all implementations of Screen interface
package gui

const (
	//Rows is the amount of rows in the display
	Rows = 8
	//Columns is the amount of columns in the display
	Columns = 8
	//VapeOn is the value for when a vape light is on
	VapeOn = 1
	//VapeOff is the value for when a vape light is off
	VapeOff = 0

	rowColStartIndex = 0
)

// Screen represents a screen or hardware visual output
//
//go:generate go run github.com/vektra/mockery/cmd/mockery -name Screen -inpkg --filename screen_mock.go
type Screen interface {
	VapeLightOn(int) error
	VapeLightOff(int) error
	DisplayMatrix([][]int) error
	AllVapesOff() error
	ShowAndRun()
	NewRow() error
}

// DisplayMatrix displays the matrix provided
func DisplayMatrix(s Screen, matrix [][]int) error {
	count := rowColStartIndex
	for _, row := range matrix {
		for _, col := range row {
			if err := lightVape(s, col, count); err != nil {
				return err
			}
			count++
		}
		if err := s.NewRow(); err != nil {
			return err
		}
	}
	return nil
}

func lightVape(s Screen, col, count int) error {
	if col == VapeOn {
		return s.VapeLightOn(count)
	}
	return s.VapeLightOff(count)
}
