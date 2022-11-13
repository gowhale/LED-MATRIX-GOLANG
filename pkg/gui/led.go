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
	sleep          = 1 // amount of ms to keep single LED on whilst multiplexing
)

func (s *LEDGUI) setRowPinLow(rowPin int) {
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

func (s *LEDGUI) CordinatesToLED(cord coordinate) {
	s.setRowPinLow(s.rowPins[cord[1]])
	setColPinHigh(s.colPins[cord[0]])
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

type LEDGUI struct {
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

	return &LEDGUI{
		rowCount: cfg.RowCount(),
		colCount: cfg.ColCount(),
		rowPins:  cfg.RowPins,
		colPins:  cfg.ColPins,
	}, nil
}

// AllVapesOff clears the termina
func (*LEDGUI) AllVapesOff() error {
	return nil
}

// DisplayMatrix displays the matrix provided
func (s *LEDGUI) DisplayMatrix(matrix [][]int, t time.Duration) error {
	startTime := time.Now()
	coordinates := letterToLED(matrix)
	for time.Since(startTime) < t {
		for _, c := range coordinates {
			s.CordinatesToLED(c)
		}
	}
	return nil
}

func (l *LEDGUI) Close() error {
	allPins := append(l.rowPins, l.colPins...)
	for _, p := range allPins {
		rpio.Pin(p).Low()
	}
	return rpio.Close()
}

func (l *LEDGUI) Rows() int {
	return l.rowCount
}

func (l *LEDGUI) Cols() int {
	return l.colCount
}
