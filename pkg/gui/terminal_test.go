// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"fmt"
	"led-matrix/pkg/config"
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
	terminal     terminalGui
	mockTerminal *mockTerminalOutputter
}

func (t *terminalSuite) SetupTest() {
	t.terminal = terminalGui{}
	t.mockTerminal = new(mockTerminalOutputter)
}

func TestTerminalSuite(t *testing.T) {
	suite.Run(t, new(terminalSuite))
}

func (t *terminalSuite) Test_DisplayMatrix_Pass() {
	err := t.terminal.DisplayMatrix(letterA, time.Millisecond)
	t.Nil(err)
}
func (t *terminalSuite) Test_displayTerminalMatrixImpl_Pass() {
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "\n").Return(nil)
	err := displayTerminalMatrixImpl(t.mockTerminal, letterA, time.Millisecond)
	t.Nil(err)
}

func (t *terminalSuite) Test_displayTerminalMatrixImpl_NewLine_Error() {
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "\n").Return(fmt.Errorf("new line error"))
	err := displayTerminalMatrixImpl(t.mockTerminal, letterA, time.Millisecond)
	t.EqualError(err, "new line error")
}

func (t *terminalSuite) Test_displayTerminalMatrixImpl_lightLED_Error() {
	t.mockTerminal.On("Printf", " ").Return(fmt.Errorf("print err"))
	err := displayTerminalMatrixImpl(t.mockTerminal, letterA, time.Millisecond)
	t.EqualError(err, "print err")

}

func (t *terminalSuite) Test_NewTerminalGui() {
	cfg := config.PinConfig{
		RowPins: make([]int, expectedRow),
		ColPins: make([]int, expectedCol),
	}
	newT := NewTerminalGui(cfg)
	t.Equal(&terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
	}, newT)
}

func (t *terminalSuite) Test_Close() {
	term := terminalGui{}
	t.Nil(term.Close())
}
