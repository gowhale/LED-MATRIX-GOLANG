// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"elf-bar-awareness/pkg/config"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// consts starting with L represent left side pins
// consts starting with R represent right side pins
const (
	sleep = 1 // amount of ms to keep single LED on whilst multiplexing

	cordXIndex = 0
	cordYIndex = 0
)

type guiLED struct {
	rowCount, colCount int
	rowPins, colPins   []int
}

// NewledGUI returns ledGUI struct to display output on terminal
func NewledGUI(cfg config.PinConfig) (Screen, error) {
	log.Println("Creating LED GUI. Opening gpio")
	err := rpio.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Setting all pins to low")
	for _, c := range cfg.ColPins {
		p := rpio.Pin(c)
		p.Output()
		p.Low()
	}
	for _, c := range cfg.RowPins {
		p := rpio.Pin(c)
		p.Output()
		p.Low()
	}

	return &guiLED{
		rowCount: cfg.RowCount(),
		colCount: cfg.ColCount(),
		rowPins:  cfg.RowPins,
		colPins:  cfg.ColPins,
	}, nil
}

func (l *guiLED) setRowPinLow(rowPin int) {
	for _, r := range s.rowPins {
		p := rpio.Pin(r)
		p.High()
	}
	p := rpio.Pin(rowPin)
	p.Low()
}

func setColPinHigh(col int) {
	p := rpio.Pin(col)
	p.High()
	time.Sleep(time.Microsecond * sleep)
	p.Low()
}

// CordinatesToLED lights up a matrix's light at specified coordinate
// Only lights temporarily used for multiplexing the lights
func (l *guiLED) CordinatesToLED(cord coordinate) {
	l.setRowPinLow(l.rowPins[cord[cordYIndex]])
	setColPinHigh(l.colPins[cord[cordXIndex]])
}

func letterToLED(l [][]int) []coordinate {
	coordinates := []coordinate{}
	for i, row := range l {
		for j, col := range row {
			if col == VapeOn {
				coordinates = append(coordinates, coordinate{j, i})
			}
		}
	}
	return coordinates
}

// AllVapesOff clears the termina
func (*guiLED) AllVapesOff() error {
	return nil
}

// DisplayMatrix displays the matrix provided
func (l *guiLED) DisplayMatrix(matrix [][]int, t time.Duration) error {
	startTime := time.Now()
	coordinates := letterToLED(matrix)
	for time.Since(startTime) < t {
		for _, c := range coordinates {
			l.CordinatesToLED(c)
		}
	}
	return nil
}

// Close turns all the gpio pins to low and closes connection
func (l *guiLED) Close() error {
	allPins := append(l.rowPins, l.colPins...)
	for _, p := range allPins {
		rpio.Pin(p).Low()
	}
	return rpio.Close()
}
