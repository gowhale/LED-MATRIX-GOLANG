// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package gui

import (
	"fmt"
	"led-matrix/pkg/config"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LEDSuite struct {
	suite.Suite

	cfg           config.PinConfig
	guiStruct     guiLED
	exampleMatrix [][]int
	exampleCoords []coordinate

	mockRPIO   *mockRpioProcessor
	mockPin    *mockPinProcessor
	mockScreen *MockScreen
}

func (l *LEDSuite) SetupTest() {
	l.mockRPIO = new(mockRpioProcessor)
	l.mockPin = new(mockPinProcessor)
	l.mockScreen = new(MockScreen)

	l.cfg = config.PinConfig{
		ColPins: []int{1, 2, 3},
		RowPins: []int{1, 2, 3},
	}
	l.guiStruct = guiLED{
		colCount:       3,
		rowCount:       3,
		colPins:        []int{1, 2, 3},
		rowPins:        []int{1, 2, 3},
		rpioController: l.mockRPIO,
	}
	l.exampleMatrix = [][]int{
		{LEDOn, LEDOn, LEDOn},
		{LEDOn, LEDOn, LEDOn},
		{LEDOn, LEDOn, LEDOn},
	}
	l.exampleCoords = []coordinate{
		{0, 0},
		{1, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{0, 2},
		{1, 2},
		{2, 2},
	}
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

	gui, err := NewledGUI(l.cfg, l.mockRPIO)
	l.Nil(err)
	l.Equal(&l.guiStruct, gui)
}

func (l *LEDSuite) Test_NewledGUI_Open_Error() {
	l.mockRPIO.On("Open").Return(fmt.Errorf("open error"))

	gui, err := NewledGUI(l.cfg, l.mockRPIO)
	l.EqualError(err, "open error")
	l.Equal(nil, gui)
}

func (l *LEDSuite) Test_Close_Pass() {
	l.mockPin.On("Output")
	l.mockPin.On("Low")

	l.mockRPIO.On("Close").Return(nil)
	l.mockRPIO.On("Pin", 1).Return(l.mockPin)
	l.mockRPIO.On("Pin", 2).Return(l.mockPin)
	l.mockRPIO.On("Pin", 3).Return(l.mockPin)
	err := l.guiStruct.Close()
	l.Nil(err)
}

func (l *LEDSuite) Test_setRowPinLow() {
	l.mockPin.On("Low")
	l.mockPin.On("High")

	l.mockRPIO.On("Pin", 1).Return(l.mockPin)
	l.mockRPIO.On("Pin", 2).Return(l.mockPin)
	l.mockRPIO.On("Pin", 3).Return(l.mockPin)

	l.guiStruct.setRowPinLow(1)
}

func (l *LEDSuite) Test_setColPinHigh() {
	l.mockPin.On("Low")
	l.mockPin.On("High")

	l.mockRPIO.On("Pin", 1).Return(l.mockPin)
	l.mockRPIO.On("Pin", 2).Return(l.mockPin)
	l.mockRPIO.On("Pin", 3).Return(l.mockPin)

	l.guiStruct.CordinatesToLED(coordinate{1, 1})
}

func (l *LEDSuite) Test_CordinatesToLED() {
	l.mockPin.On("Low")
	l.mockPin.On("High")

	l.mockRPIO.On("Pin", 1).Return(l.mockPin)

	l.guiStruct.setColPinHigh(1)
}

func (l *LEDSuite) Test_letterToLED() {
	cords := letterToLED(l.exampleMatrix)
	l.Equal(l.exampleCoords, cords)
}

func (l *LEDSuite) Test_AllLEDSOff() {
	l.Nil(l.guiStruct.AllLEDSOff())
}

func (l *LEDSuite) Test_displayMatrixImp_Pass() {
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[0]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[1]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[2]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[3]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[4]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[5]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[6]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[7]).Return(nil)
	l.mockScreen.On("CordinatesToLED", l.exampleCoords[8]).Return(nil)
	err := displayMatrixImp(l.exampleMatrix, time.Millisecond, l.mockScreen)
	l.Nil(err)
}
