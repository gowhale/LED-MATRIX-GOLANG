// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package gui

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const (
	vOffFuncName   = "VapeLightOff"
	vOnFuncName    = "VapeLightOn"
	newRowFuncName = "NewRow"
)

type guiSuite struct {
	suite.Suite
	mockScreen *MockScreen
}

func (t *guiSuite) SetupTest() {
	t.mockScreen = new(MockScreen)
}

func TestQuizTestSuite(t *testing.T) {
	suite.Run(t, new(guiSuite))
}

func (t *guiSuite) Test_DisplayMatrix_Pass() {
	testMatrix := letterA

	t.mockScreen.On(vOffFuncName, mock.Anything).Return(nil)
	t.mockScreen.On(vOnFuncName, mock.Anything).Return(nil)
	t.mockScreen.On(newRowFuncName).Return(nil)

	err := DisplayMatrix(t.mockScreen, testMatrix)
	t.Nil(err)
}
