// Package gui is responsible for visual output
// File gui.go contains infomration used by all implementations of Screen interface
package gui

import "time"

const (
	//LEDOn is the value for when a light is on
	LEDOn = 1
	//LEDOff is the value for when a light is off
	LEDOff = 0

	rowColStartIndex = 0
)

type coordinate [2]int // represents an x and y on matrix

// Screen represents a screen or hardware visual output
//
//go:generate go run github.com/vektra/mockery/cmd/mockery -name Screen -inpkg --filename screen_mock.go
type Screen interface {
	AllLEDSOff() error                                                 // Set all LEDS to off
	Close() error                                                      // Defer to prevent errors next run
	DisplayMatrix(matrix [][]int, displayDuration time.Duration) error // Displays matrix for x amount of time
	CordinatesToLED(coordinate)                                        // CordinatesToLED lights up a matrix's light at specified coordinate
}
