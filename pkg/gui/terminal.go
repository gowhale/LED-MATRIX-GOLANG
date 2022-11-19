// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"os"
	"os/exec"
	"time"

	"github.com/gowhale/led-matrix/pkg/config"
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

var execCommand = exec.Command

// AllLEDSOff clears the terminal
func (*terminalGui) AllLEDSOff() error {
	cmd := execCommand("clear")
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

func printRow(t *terminalGui, x, y, count int, cords coordinate) error {
	if count%t.colCount == 0 && count > 0 {
		if err := t.to.Printf("#\n#"); err != nil {
			return err
		}
	}
	if x == cords[cordXIndex] && y == cords[cordYIndex] {
		if err := t.to.Printf("0"); err != nil {
			return err
		}
	}
	return t.to.Printf(" ")
}

func (t *terminalGui) CordinatesToLED(cords coordinate) error {
	if err := t.AllLEDSOff(); err != nil {
		return err
	}
	if err := t.to.Printf("cord: x=%d y=%d\n##########\n#", cords[cordXIndex], cords[cordYIndex]); err != nil {
		return err
	}

	count := 0
	for x := 0; x < t.colCount; x++ {
		for y := 0; y < t.rowCount; y++ {
			if err := printRow(t, x, y, count, cords); err != nil {
				return err
			}
			count++
		}
	}

	return t.to.Printf("#\n##########")
}
