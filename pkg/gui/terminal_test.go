// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/gowhale/led-matrix-golang/pkg/config"

	"github.com/stretchr/testify/suite"
)

const (
	expectedRow = 8
	expectedCol = 8

	exitCodePass = 0
	exitCodeFail = 1
)

type terminalSuite struct {
	suite.Suite
	terminal     terminalGui
	mockTerminal *mockTerminalOutputter
}

func (t *terminalSuite) SetupTest() {
	t.terminal = terminalGui{}
	t.mockTerminal = new(mockTerminalOutputter)

	execCommand = fakeExecCommandPass
}

func (*terminalSuite) AfterTest() {
	execCommand = exec.Command
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
		to:       &terminalOutput{},
	}, newT)
}

func (t *terminalSuite) Test_Close() {
	term := terminalGui{}
	t.Nil(term.Close())
}

func TestHelperProcess(*testing.T) {
	helper := os.Getenv("GO_WANT_HELPER_PROCESS")
	//pass
	if helper == "1" {
		os.Exit(exitCodePass)
		return
	}
	//fail
	if helper == "2" {
		os.Exit(exitCodeFail)
		return
	}
}

func fakeExecCommandPass(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func (t *terminalSuite) Test_CordinatesToLED_Pass() {
	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       &terminalOutput{},
	}

	t.Nil(term.CordinatesToLED(coordinate{1, 1}))
}

func (t *terminalSuite) Test_CordinatesToLED_mockTerminal_Pass() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(nil)
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "#\n#").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.Nil(term.CordinatesToLED(coordinate{1, 1}))
}

func (t *terminalSuite) Test_CordinatesToLED_CordPrint_Error() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(fmt.Errorf("print err"))
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "#\n#").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.EqualError(term.CordinatesToLED(coordinate{1, 1}), "print err")
}

func (t *terminalSuite) Test_CordinatesToLED_Space_Error() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(nil)
	t.mockTerminal.On("Printf", " ").Return(fmt.Errorf("print err"))
	t.mockTerminal.On("Printf", "#\n#").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.EqualError(term.CordinatesToLED(coordinate{1, 1}), "print err")
}

func (t *terminalSuite) Test_CordinatesToLED_NewLine_Error() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(nil)
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "#\n#").Return(fmt.Errorf("print err"))
	t.mockTerminal.On("Printf", "0").Return(nil)
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.EqualError(term.CordinatesToLED(coordinate{1, 1}), "print err")
}

func (t *terminalSuite) Test_CordinatesToLED_Zero_Error() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(nil)
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "#\n#").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(fmt.Errorf("print err"))
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.EqualError(term.CordinatesToLED(coordinate{1, 1}), "print err")
}

func (t *terminalSuite) Test_CordinatesToLED_EndLine_Error() {
	t.mockTerminal.On("Printf", "cord: x=%d y=%d\n%s\n#", 1, 1, "##########").Return(nil)
	t.mockTerminal.On("Printf", " ").Return(nil)
	t.mockTerminal.On("Printf", "#\n#").Return(nil)
	t.mockTerminal.On("Printf", "0").Return(fmt.Errorf("print err"))
	t.mockTerminal.On("Printf", "#\n%s\n", "##########").Return(nil)

	term := &terminalGui{
		rowCount: expectedRow,
		colCount: expectedCol,
		to:       t.mockTerminal,
	}
	t.EqualError(term.CordinatesToLED(coordinate{1, 1}), "print err")
}
