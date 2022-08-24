// Package main is the entry point for our binary
package main

import (
	"elf-bar-awareness/pkg/gui"
	"fmt"
	"log"
	"time"
)

const (
	refreshTime = time.Millisecond * 100
	offsetStart = 0
	offsetLimit = 1000
)

func main() {
	screen := gui.NewTerminalGui()
	defer func() {
		screen.ShowAndRun()
	}()
	word := "    lets get ready to rumble"
	matrix, err := concatanateLetters(word)
	if err != nil {
		log.Panicln(err)
	}
	go func() {
		offset := offsetStart
		for offset < offsetLimit {
			trimmedMatrix, err := trimMatrix(matrix, offset)
			if err != nil {
				log.Panicln(err)
			}
			if err := screen.AllVapesOff(); err != nil {
				log.Fatalln(err)
			}
			if err := screen.DisplayMatrix(trimmedMatrix); err != nil {
				log.Panicln(err)
			}
			time.Sleep(refreshTime)
			offset++
		}
	}()
}

func trimMatrix(matrix [][]int, offset int) ([][]int, error) {
	newMatrix := make([][]int, gui.Rows)
	for i := range newMatrix {
		newMatrix[i] = make([]int, gui.Columns)
	}
	for i := range matrix {
		count := offsetStart
		currentOffset := offset + count

		for (count < len(matrix[i])) && ((offset - currentOffset) < gui.Columns) && currentOffset < len(matrix[i]) && count < gui.Columns {
			newMatrix[i][count] = matrix[i][currentOffset]
			currentOffset++
			count++
		}
	}
	for i := range newMatrix {
		for len(newMatrix[i]) < gui.Columns {
			newMatrix[i] = append(newMatrix[i], gui.VapeOff)
		}
	}
	return newMatrix, nil
}

func concatanateLetters(word string) ([][]int, error) {
	sentenceAsLEDs := [][]int{}
	for index, l := range word {
		newLetter, ok := gui.LetterMap[string(l)]
		if !ok {
			return nil, fmt.Errorf("l=%s not in LetterMap", string(l))
		}
		if index == offsetStart {
			sentenceAsLEDs = append([][]int{}, newLetter...)
		}
		if index != offsetStart {
			for j, r := range newLetter {
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], gui.VapeOff)
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], r...)
			}
		}
	}
	return sentenceAsLEDs, nil
}
