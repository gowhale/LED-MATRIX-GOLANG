// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

// consts starting with L represent left side pins
// consts starting with R represent right side pins
const (
	L10 = 5
	L9  = 6
	L11 = 12
	L8  = 13
	L6  = 16
	L7  = 19
	L3  = 20
	L14 = 21

	R13 = 4
	R16 = 17
	R10 = 18
	R15 = 22
	R7  = 23
	R8  = 24
	R21 = 25
	R18 = 26
	R14 = 27

	sleep          = 1 // amount of ms to keep single LED on whilst multiplexing
	runTimeSeconds = 100
)

var RowsPins = []int{R7,
	L3,
	R10,
	L6,
	L11,
	R18,
	L14,
	R21}

var ColPins = []int{L10,
	L9,
	L8,
	L7,
	R13,
	R14,
	R15,
	R16}

func setRowPinLow(rowPin int) {
	for _, r := range RowsPins {
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

func CordinatesToLED(cord []int) {
	setRowPinLow(RowsPins[cord[1]])
	setColPinHigh(ColPins[cord[0]])
}

func letterToLED(l [][]int) [][]int {
	coordinates := [][]int{}
	for i, row := range l {
		for j, col := range row {
			if col == VapeOn {
				coordinates = append(coordinates, []int{j, i})
			}
		}
	}
	return coordinates
}

type LEDGUI struct {
	rows, cols int
}

// type matrixMapping struct{}

// func sizeValidation(rows, cols int) matrixMapping

// NewledGUI returns ledGUI struct to display output on terminal
func NewledGUI(rows, cols int) (Screen, error) {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	log.Println("New LED GUI")
	for _, c := range ColPins {
		log.Println("pin =", c)
		p := rpio.Pin(c)
		p.Output()
		p.Low()
	}
	for _, c := range RowsPins {
		log.Println("pin =", c)
		p := rpio.Pin(c)
		p.Output()
		p.Low()
	}
	return &LEDGUI{
		rows: rows,
		cols: cols,
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
			CordinatesToLED(c)
		}
	}
	return nil
}

func (*LEDGUI) Close() error {
	allPins := append(RowsPins, ColPins...)
	for _, p := range allPins {
		rpio.Pin(p).Low()
	}
	return rpio.Close()
}

func (l *LEDGUI) Rows() int {
	return l.rows
}

func (l *LEDGUI) Cols() int {
	return l.cols
}
