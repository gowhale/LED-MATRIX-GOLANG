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
	L3  = 20
	L6  = 16
	L7  = 19
	L8  = 13
	L9  = 6
	L10 = 5
	L11 = 12
	L14 = 21

	R7  = 23
	R10 = 18
	R13 = 4
	R14 = 27
	R15 = 22
	R16 = 17
	R18 = 26
	R21 = 25

	sleep          = 1 // amount of ms to keep single LED on whilst multiplexing
	runTimeSeconds = 100
)

var rows = []int{R7,
	L3,
	R10,
	L6,
	L11,
	R18,
	L14,
	R21}

var cols = []int{L10,
	L9,
	L8,
	L7,
	R13,
	R14,
	R15,
	R16}

func setRow(rowPin int) {
	for _, r := range rows {
		p := rpio.Pin(r)
		p.High()
	}
	p := rpio.Pin(rowPin)
	p.Low()
}

func flash(col int) {
	p := rpio.Pin(col)
	p.High()
	time.Sleep(time.Microsecond * sleep)
	p.Low()
}

func cordinatesToLED(cord []int) {
	setRow(rows[cord[1]])
	flash(cols[cord[0]])
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

type LEDGUI struct{}

// NewledGUI returns ledGUI struct to display output on terminal
func NewledGUI() (Screen, error) {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	log.Println("New LED GUI")
	for _, c := range cols {
		log.Println("pin =", c)
		p := rpio.Pin(c)
		p.Output()
	}
	for _, c := range rows {
		log.Println("pin =", c)
		p := rpio.Pin(c)
		p.Output()
	}
	return &LEDGUI{}, nil
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
			cordinatesToLED(c)
		}
	}
	return nil
}

func (*LEDGUI) Close() error {
	return rpio.Close()
}
