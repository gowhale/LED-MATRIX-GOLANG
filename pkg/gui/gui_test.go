// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package gui

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type guiSuite struct {
	suite.Suite
	mockScreen *MockScreen
}

func (t *guiSuite) SetupTest() {
	t.mockScreen = new(MockScreen)
}

func TestGuiSuite(t *testing.T) {
	suite.Run(t, new(guiSuite))
}

func (t *guiSuite) Test_displayLEDMatrix_Pass() {
	testMatrix := [][]int{{LEDOn}}

	t.mockScreen.On("CordinatesToLED", coordinate{0, 0}).Return(nil)

	err := displayLEDMatrix(testMatrix, time.Millisecond*500, t.mockScreen)
	t.Nil(err)
}

func (t *guiSuite) Test_displayLEDMatrix_Error() {
	testMatrix := [][]int{{LEDOn}}

	t.mockScreen.On("CordinatesToLED", coordinate{0, 0}).Return(fmt.Errorf("cord error"))

	err := displayLEDMatrix(testMatrix, time.Millisecond*500, t.mockScreen)
	t.EqualError(err, "cord error")
}
