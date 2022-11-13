// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package gui

import (
	"fmt"
	"led-matrix/pkg/config"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LEDSuite struct {
	suite.Suite

	cfg      config.PinConfig
	mockRPIO *mockRpioProcessor
	mockPin  *mockPinProcessor
}

func (l *LEDSuite) SetupTest() {
	l.cfg = config.PinConfig{
		ColPins: []int{1, 2, 3},
		RowPins: []int{1, 2, 3},
	}
	l.mockRPIO = new(mockRpioProcessor)
	l.mockPin = new(mockPinProcessor)
}

func Test_LED_Suite(t *testing.T) {
	suite.Run(t, new(LEDSuite))
}

func (l *LEDSuite) Test_NewledGUI_Pass() {
	l.mockPin.On("Output")
	l.mockPin.On("Low")

	l.mockRPIO.On("Open").Return(nil)
	l.mockRPIO.On("Pin", 1).Return(l.mockPin)
	l.mockRPIO.On("Pin", 2).Return(l.mockPin)
	l.mockRPIO.On("Pin", 3).Return(l.mockPin)

	expectedOutput := &guiLED{
		colCount: 3,
		rowCount: 3,
		colPins:  []int{1, 2, 3},
		rowPins:  []int{1, 2, 3},
	}

	gui, err := NewledGUI(l.cfg, l.mockRPIO)
	l.Nil(err)
	l.Equal(expectedOutput, gui)
}

func (l *LEDSuite) Test_NewledGUI_Open_Error() {
	l.mockRPIO.On("Open").Return(fmt.Errorf("open error"))

	gui, err := NewledGUI(l.cfg, l.mockRPIO)
	l.EqualError(err, "open error")
	l.Equal(nil, gui)
}
