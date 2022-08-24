// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type terminalSuite struct {
	suite.Suite
	terminal   terminalGui
}

func (t *terminalSuite) SetupTest() {
	t.terminal = terminalGui{}
}

func TestTerminalSuite(t *testing.T) {
	suite.Run(t, new(terminalSuite))
}

func (t *terminalSuite) Test_VapeLightOn_Pass() {
	err := t.terminal.VapeLightOn(VapeOff)
	t.Nil(err)
}

func (t *terminalSuite) Test_VapeLightOff_Pass() {
	err := t.terminal.VapeLightOff(VapeOff)
	t.Nil(err)
}

func (t *terminalSuite) Test_NewRow_Pass() {
	err := t.terminal.NewRow()
	t.Nil(err)
}

func (t *terminalSuite) Test_DisplayMatrix_Pass() {
	err := t.terminal.DisplayMatrix(letterA)
	t.Nil(err)
}

func (t *terminalSuite) Test_NewTerminalGui() {
	newT := NewTerminalGui()
	t.Equal(&terminalGui{}, newT)
}
