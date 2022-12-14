package matrix

import (
	"fmt"

	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

const (
	// OffsetStart is the start val for offsetting
	OffsetStart = 0
)

// TrimMatrix will take in a matrix then trim it according to the amount of rows, cols and offset
func TrimMatrix(matrix [][]int, cfg config.PinConfig, offset int) ([][]int, error) {
	newMatrix := make([][]int, cfg.RowCount())
	for i := range newMatrix {
		newMatrix[i] = make([]int, cfg.ColCount())
	}
	for i := range matrix {
		count := OffsetStart
		currentOffset := offset + count

		for (count < len(matrix[i])) && ((offset - currentOffset) < cfg.ColCount()) && currentOffset < len(matrix[i]) && count < cfg.ColCount() {
			newMatrix[i][count] = matrix[i][currentOffset]
			currentOffset++
			count++
		}
	}
	for i := range newMatrix {
		for len(newMatrix[i]) < cfg.ColCount() {
			newMatrix[i] = append(newMatrix[i], gui.LEDOff)
		}
	}
	return newMatrix, nil
}

// ConcatanateLetters will create a matrix
func ConcatanateLetters(word string) ([][]int, error) {
	sentenceAsLEDs := [][]int{}
	for index, l := range word {
		newLetter, ok := gui.LetterMap[string(l)]
		if !ok {
			return nil, fmt.Errorf("l=%s not in LetterMap", string(l))
		}
		if index == OffsetStart {
			sentenceAsLEDs = append([][]int{}, newLetter...)
		}
		if index != OffsetStart {
			for j, r := range newLetter {
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], gui.LEDOff)
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], r...)
			}
		}
	}
	return sentenceAsLEDs, nil
}
