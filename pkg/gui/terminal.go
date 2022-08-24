// Package gui is responsible for visual output
// File terminal.go implements terminal methods for gui interface
package gui

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	amountOfHoursToWaitToEndDefer = 100
)

type terminalGui struct{}

//NewTerminalGui returns terminalGui struct to display output on terminal
func NewTerminalGui() Screen {
	return &terminalGui{}
}

// ShowAndRun waits so the update of letters can occure
func (*terminalGui) ShowAndRun() {
	time.Sleep(amountOfHoursToWaitToEndDefer * time.Hour)
}

// VapeLightOn prints out "0"
func (*terminalGui) VapeLightOn(_ int) error {
	_, err := fmt.Printf("0")
	return err
}

// VapeLightOn prints out " "
func (*terminalGui) VapeLightOff(_ int) error {
	_, err := fmt.Printf(" ")
	return err
}

// AllVapesOff clears the termina
func (*terminalGui) AllVapesOff() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// DisplayMatrix displays the matrix provided
func (s *terminalGui) DisplayMatrix(matrix [][]int) error {
	return DisplayMatrix(s, matrix)
}

func (*terminalGui) NewRow() error {
	_, err := fmt.Printf("\n")
	return err
}
