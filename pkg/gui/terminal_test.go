// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
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
	newT := NewTerminalGui(8, 8)
	t.Equal(&terminalGui{
		rows: 8,
		cols: 8,
	}, newT)
}
