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
	windows, screen := gui.NewGui()
	word := "    lets get ready to rumble"
	matrix, err := concatanateLetters(word)
	if err != nil {
		log.Panicln(err)
	}
	go func() {
		offset := offsetStart
		for offset < offsetLimit {
			if err := screen.AllVapesOff(); err != nil {
				log.Fatalln(err)
			}
			trimmedMatrix, err := trimMatrix(matrix, offset)
			if err != nil {
				log.Panicln(err)
			}
			if err := screen.DisplayLetter(trimmedMatrix); err != nil {
				log.Panicln(err)
			}
			time.Sleep(refreshTime)
			offset++
			if _, err := fmt.Printf("offset=%d", offset); err != nil {
				log.Panicln(err)
			}
		}
	}()
	if _, err := fmt.Println("showing"); err != nil {
		log.Panicln(err)
	}
	windows.ShowAndRun()
}

func trimMatrix(matrix [][]int, offset int) ([][]int, error) {
	if _, err := fmt.Println("Creating new matrix"); err != nil {
		return nil, err
	}
	newMatrix := make([][]int, gui.Rows)
	for i := range newMatrix {
		newMatrix[i] = make([]int, gui.Columns)
	}

	if _, err := fmt.Println("Copying old matrix with offset"); err != nil {
		return nil, err
	}
	for i := range matrix {
		count := offsetStart
		currentOffset := offset + count

		for (count < len(matrix[i])) && ((offset - currentOffset) < gui.Columns) && currentOffset < len(matrix[i]) && count < gui.Columns {
			// fmt.Println((count < len(matrix[i])), len(matrix[i]), currentOffset-offset, count)
			newMatrix[i][count] = matrix[i][currentOffset]
			currentOffset++
			count++
		}
	}
	if _, err := fmt.Println("filling new matrix"); err != nil {
		return nil, err
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
				log.Println(len(r))
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], gui.VapeOff)
				sentenceAsLEDs[j] = append(sentenceAsLEDs[j], r...)
			}
		}
	}

	return sentenceAsLEDs, nil
}
