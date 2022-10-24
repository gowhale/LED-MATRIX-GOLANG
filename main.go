// Package main is the entry point for our binary
package main

import (
	"elf-bar-awareness/pkg/gui"
	mx "elf-bar-awareness/pkg/matrix"
	"log"
	"time"
)

const (
	refreshTime = time.Second
	offsetLimit = 1000
)

func main() {
	screen := gui.NewTerminalGui()
	defer func() {
		err := screen.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	word := "lets get ready to rumble"
	matrix, err := mx.ConcatanateLetters(word)
	if err != nil {
		log.Panicln(err)
	}
	for offset := mx.OffsetStart; offset < offsetLimit; offset++ {
		trimmedMatrix, err := mx.TrimMatrix(matrix, gui.Rows, gui.Columns, offset)
		if err != nil {
			log.Fatalln(err)
		}
		if err := screen.AllVapesOff(); err != nil {
			log.Fatalln(err)
		}
		if err := screen.DisplayMatrix(trimmedMatrix, refreshTime); err != nil {
			log.Panicln(err)
		}
	}
	log.Println("fin.")
}
