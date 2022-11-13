// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"led-matrix/pkg/config"
	"os"
	"os/exec"
	"time"
)

type terminalGui struct {
	colCount, rowCount int
	to                 terminalOutputter
}

// NewTerminalGui returns terminalGui struct to display output on terminal
func NewTerminalGui(cfg config.PinConfig) Screen {
	return &terminalGui{
		rowCount: len(cfg.RowPins),
		colCount: len(cfg.ColPins),
		to:       &terminalOutput{},
	}
}

// AllLEDSOff clears the termina
func (*terminalGui) AllLEDSOff() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DisplayMatrix displays the matrix provided
func (*terminalGui) DisplayMatrix(matrix [][]int, duration time.Duration) error {
	return displayTerminalMatrixImpl(&terminalOutput{}, matrix, duration)
}

func (*terminalGui) Close() error {
	return nil
}

// DisplayMatrix displays the matrix provided
func displayTerminalMatrixImpl(t terminalOutputter, matrix [][]int, duration time.Duration) error {
	count := rowColStartIndex
	for _, row := range matrix {
		for _, col := range row {
			if err := lightLED(t, col); err != nil {
				return err
			}
			count++
		}
		if err := t.Printf("\n"); err != nil {
			return err
		}
	}
	time.Sleep(duration)
	return nil
}

func lightLED(t terminalOutputter, col int) error {
	if col == LEDOn {
		return t.Printf("0")
	}
	return t.Printf(" ")
}

func (t *terminalGui) CordinatesToLED(cords coordinate) error {
	for y := 0; y < t.rowCount; y++ {
		for x := 0; x < t.colCount; x++ {
			if x == cords[cordXIndex] && y == cords[cordYIndex] {
				if err := t.to.Printf("0"); err != nil {
					return err
				}
			} else {
				if err := t.to.Printf(" "); err != nil {
					return err
				}
			}
		}
		if err := t.to.Printf("\n"); err != nil {
			return err
		}

	}
	return nil
}
