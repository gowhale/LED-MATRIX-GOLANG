// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"elf-bar-awareness/pkg/config"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type terminalGui struct {
	colCount, rowCount int
}

// NewTerminalGui returns terminalGui struct to display output on terminal
func NewTerminalGui(cfg config.PinConfig) Screen {
	return &terminalGui{
		rowCount: len(cfg.RowPins),
		colCount: len(cfg.ColPins),
	}
}

// AllVapesOff clears the termina
func (*terminalGui) AllVapesOff() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DisplayMatrix displays the matrix provided
func (s *terminalGui) DisplayMatrix(matrix [][]int, t time.Duration) error {
	err := DisplayMatrix(s, matrix)
	if err != nil {
		return err
	}
	time.Sleep(t)
	return nil
}

func (*terminalGui) Close() error {
	return nil
}

// DisplayMatrix displays the matrix provided
func DisplayMatrix(s Screen, matrix [][]int) error {
	count := rowColStartIndex
	for _, row := range matrix {
		for _, col := range row {
			if err := lightVape(s, col, count); err != nil {
				return err
			}
			count++
		}
		_, err := fmt.Printf("\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func lightVape(s Screen, col, count int) error {
	if col == VapeOn {
		_, err := fmt.Printf("0")
		return err
	}
	_, err := fmt.Printf(" ")
	return err
}

func (t *terminalGui) Rows() int {
	return t.rowCount
}

func (t *terminalGui) Cols() int {
	return t.colCount
}

func (t *terminalGui) CordinatesToLED(cord coordinate) {

}
