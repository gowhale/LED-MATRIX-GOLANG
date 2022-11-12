// Package gui is responsible for visual output
// File gui.go contains infomration used by all implementations of Screen interface
package gui

import "time"

const (
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
	Close() error                                                      // Defer to prevent errors next run
	DisplayMatrix(matrix [][]int, displayDuration time.Duration) error // Displays matrix for x amount of time
	Rows() int                                                         // Displays matrix for x amount of time
	Cols() int                                                         // Displays matrix for x amount of time
	CordinatesToLED([2]int)                                            // Displays matrix for x amount of time
}
