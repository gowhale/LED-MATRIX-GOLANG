// Package gui is responsible for visual output
// File gui.go contains infomration used by all implementations of Screen interface
package gui

import "time"

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
	AllVapesOff() error                                                // Set all vapes to off
	Close() error                                                      // Closes the connections
	DisplayMatrix(matrix [][]int, displayDuration time.Duration) error // Displays matrix for x amount of time
}
