// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"elf-bar-awareness/pkg/config"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	expectedRow = 8
	expectedCol = 8
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

func (t *terminalSuite) Test_DisplayMatrix_Pass() {
	err := t.terminal.DisplayMatrix(letterA, time.Millisecond)
	t.Nil(err)
}

func (t *terminalSuite) Test_NewTerminalGui() {
	cfg := config.PinConfig{
		RowPins: []int{rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int()},
		ColPins: []int{rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int(), rand.Int()},
	}
	newT := NewTerminalGui(cfg)
	t.Equal(&terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
	}, newT)
}
