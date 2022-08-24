// Package gui is responsible for visual output
// File terminal_test.go tests the terminal.go file
package gui

import (
	"fmt"
	color "image/color"
	"testing"

	"fyne.io/fyne/v2/canvas"
	"github.com/stretchr/testify/suite"
)

type fyneioSuite struct {
	suite.Suite
	fyneGui    fyneScreen
	mockScreen *MockScreen
}

func (t *fyneioSuite) SetupTest() {
	t.fyneGui = fyneScreen{}
	t.mockScreen = new(MockScreen)
}

func TestFyneioSuite(t *testing.T) {
	suite.Run(t, new(fyneioSuite))
}

func (*fyneioSuite) Test_NewFyneGui() {
	// How to actually test this?
	_ = NewFyneGui()
}

func (t *fyneioSuite) Test_VapeLightOn_Pass() {
	t.fyneGui.vapes = []*canvas.Circle{canvas.NewCircle(color.Black)}
	err := t.fyneGui.VapeLightOn(VapeOff)
	t.Nil(err)
}

func (t *fyneioSuite) Test_VapeLightOn_Error() {
	err := t.fyneGui.VapeLightOn(VapeOff)
	t.EqualError(err, "vapeID=0 is greater than len(s.vapes)=0")
}

func (t *fyneioSuite) Test_VapeLightOff_Pass() {
	t.fyneGui.vapes = []*canvas.Circle{canvas.NewCircle(color.Black)}
	err := t.fyneGui.VapeLightOff(VapeOff)
	t.Nil(err)
}

func (t *fyneioSuite) Test_VapeLightOff_Error() {
	err := t.fyneGui.VapeLightOff(VapeOff)
	t.EqualError(err, "vapeID=0 is greater than len(s.vapes)=0")
}

func (t *fyneioSuite) Test_NewRow_Pass() {
	err := t.fyneGui.NewRow()
	t.Nil(err)
}

// func (t *fyneioSuite) Test_DisplayMatrix_Pass() {
// 	err := t.terminal.DisplayMatrix(nil)
// 	t.Nil(err)
// }

func (t *fyneioSuite) Test_NewTerminalGui() {
	newT := NewTerminalGui()
	t.Equal(&terminalGui{}, newT)
}

func (t *fyneioSuite) Test_createCanvasAndDots() {
	canvasObjs, circleObjs := createCanvasAndDots()
	t.Equal(Rows*Columns, len(canvasObjs))
	t.Equal(Rows*Columns, len(circleObjs))
}

func (t *fyneioSuite) Test_AllVapesOff_Pass() {
	testVape := []*canvas.Circle{canvas.NewCircle(color.Black)}

	t.mockScreen.On("VapeLightOff", VapeOff).Return(nil)

	err := allVapesOff(t.mockScreen, testVape)
	t.Nil(err)
}

func (t *fyneioSuite) Test_AllVapesOff_Error() {
	testVape := []*canvas.Circle{canvas.NewCircle(color.Black)}

	t.mockScreen.On("VapeLightOff", VapeOff).Return(fmt.Errorf("off error"))

	err := allVapesOff(t.mockScreen, testVape)
	t.EqualError(err, "off error")
}
