// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"led-matrix/pkg/config"
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

type RPIOProc struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name rpioProcessor -inpkg --filename rpio_processor_mock.go
type rpioProcessor interface {
	Open() (err error)
	Close() (err error)
	Pin(p int) pinProcessor
}

func (*RPIOProc) Open() (err error) {
	return rpio.Open()
}

func (*RPIOProc) Close() (err error) {
	return rpio.Close()
}

func (*RPIOProc) Pin(p int) pinProcessor {
	return rpio.Pin(p)
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name pinProcessor -inpkg --filename pin_processor_mock.go
type pinProcessor interface {
	Output()
	Low()
	High()
}

type guiLED struct {
	rowCount, colCount int
	rowPins, colPins   []int
	rpioController     rpioProcessor
}

// NewledGUI returns ledGUI struct to display output on terminal
func NewledGUI(cfg config.PinConfig, rp rpioProcessor) (Screen, error) {
	log.Println("Creating LED GUI. Opening gpio")
	err := rp.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Setting all pins to low")
	for _, c := range cfg.ColPins {
		p := rp.Pin(c)
		p.Output()
		p.Low()
	}
	for _, c := range cfg.RowPins {
		p := rp.Pin(c)
		p.Output()
		p.Low()
	}

	return &guiLED{
		rowCount:       cfg.RowCount(),
		colCount:       cfg.ColCount(),
		rowPins:        cfg.RowPins,
		colPins:        cfg.ColPins,
		rpioController: rp,
	}, nil
}

func (l *guiLED) setRowPinLow(rowPin int) {
	for _, r := range l.rowPins {
		p := l.rpioController.Pin(r)
		p.High()
	}
	p := l.rpioController.Pin(rowPin)
	p.Low()
}

func (l *guiLED) setColPinHigh(col int) {
	p := l.rpioController.Pin(col)
	p.High()
	time.Sleep(time.Microsecond * sleep)
	p.Low()
}

// CordinatesToLED lights up a matrix's light at specified coordinate
// Only lights temporarily used for multiplexing the lights
func (l *guiLED) CordinatesToLED(cord coordinate) {
	l.setRowPinLow(l.rowPins[cord[cordYIndex]])
	l.setColPinHigh(l.colPins[cord[cordXIndex]])
}

func letterToLED(l [][]int) []coordinate {
	coordinates := []coordinate{}
	for i, row := range l {
		for j, col := range row {
			if col == LEDOn {
				coordinates = append(coordinates, coordinate{j, i})
			}
		}
	}
	return coordinates
}

// AllLEDSOff clears the termina
func (*guiLED) AllLEDSOff() error {
	return nil
}

func displayMatrixImp(matrix [][]int, t time.Duration, l Screen) error {
	startTime := time.Now()
	coordinates := letterToLED(matrix)
	for time.Since(startTime) < t {
		for _, c := range coordinates {
			l.CordinatesToLED(c)
		}
	}
	return nil
}

// DisplayMatrix displays the matrix provided
func (l *guiLED) DisplayMatrix(matrix [][]int, t time.Duration) error {
	return displayMatrixImp(matrix, t, l)
}

// Close turns all the gpio pins to low and closes connection
func (l *guiLED) Close() error {
	allPins := append(l.rowPins, l.colPins...)
	for _, p := range allPins {
		l.rpioController.Pin(p).Low()
	}
	return l.rpioController.Close()
}
