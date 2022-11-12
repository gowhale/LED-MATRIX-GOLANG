// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"elf-bar-awareness/pkg/config"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type terminalSuite struct {
	suite.Suite
	terminal terminalGui
}

func (t *terminalSuite) SetupTest() {
	t.terminal = terminalGui{}
}

func TestTerminalSuite(t *testing.T) {
	suite.Run(t, new(terminalSuite))
}

func (t *terminalSuite) Test_NewRow_Pass() {
}

func (t *terminalSuite) Test_DisplayMatrix_Pass() {
	err := t.terminal.DisplayMatrix(letterA, time.Millisecond)
	t.Nil(err)
}

func (t *terminalSuite) Test_NewTerminalGui() {
	cfg := config.PinConfig{RowPins: []int{1, 3, 4, 5, 6, 7, 8, 9}, ColPins: []int{10, 11, 12, 13, 14, 15, 16, 17}}
	newT := NewTerminalGui(cfg)
	t.Equal(&terminalGui{
		rowCount: 8,
		colCount: 8,
	}, newT)
}
