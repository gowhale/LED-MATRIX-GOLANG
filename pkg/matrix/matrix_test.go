// Package gui is responsible for visual output
// File gui_test.go tests the gui.go file
package matrix

import (
	"elf-bar-awareness/pkg/gui"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testRows   = 1
	testCols   = 5
	testOffSet = 5
)

type matrixSuite struct {
	suite.Suite
}

func (*matrixSuite) SetupTest() {}

func TestQuizTestSuite(t *testing.T) {
	suite.Run(t, new(matrixSuite))
}

func (t *matrixSuite) Test_DisplayMatrix_Pass() {
	testMatrix := [][]int{{gui.VapeOff, gui.VapeOff, gui.VapeOff, gui.VapeOff, gui.VapeOff}}

	m, err := TrimMatrix(testMatrix, testRows, testCols, testOffSet)
	t.Equal(testMatrix, m)
	t.Nil(err)
}

// func (t *matrixSuite) Test_DisplayMatrix_Append_Pass() {
// 	testMatrix := [][]int{}
// 	expectedMatrix := [][]int{{gui.VapeOff, gui.VapeOff}}

// 	m, err := TrimMatrix(testMatrix, 1, 10, 0)
// 	t.Equal(expectedMatrix, m)
// 	t.Nil(err)
// }

func (t *matrixSuite) Test_ConcatanateLetters_SingleLetter_Pass() {
	m, err := ConcatanateLetters("a")
	t.Equal(gui.LetterMap["a"], m)
	t.Nil(err)
}

func (t *matrixSuite) Test_ConcatanateLetters_TwoLetter_Pass() {
	m, err := ConcatanateLetters("  ")
	letterOne := gui.LetterMap[" "]
	letterTwo := gui.LetterMap[" "]
	expectedMatrix := letterOne
	for i := range letterTwo {
		expectedMatrix[i] = append(expectedMatrix[i], letterTwo[i]...)
		expectedMatrix[i] = append(expectedMatrix[i], gui.VapeOff)
	}
	t.Equal(expectedMatrix, m)
	t.Nil(err)
}

func (t *matrixSuite) Test_ConcatanateLetters_SingleLetter_Error() {
	m, err := ConcatanateLetters("ç")
	t.Nil(m)
	t.EqualError(err, "l=ç not in LetterMap")
}
