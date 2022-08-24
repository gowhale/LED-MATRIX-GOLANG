// Package main is the entry point for our binary
package main

import (
	"elf-bar-awareness/pkg/gui"
	mx "elf-bar-awareness/pkg/matrix"
	"log"
	"time"
)

const (
	refreshTime = time.Millisecond * 100
	offsetLimit = 1000
)

func main() {
	screen := gui.NewTerminalGui()
	defer func() {
		screen.ShowAndRun()
	}()
	word := "    lets get ready to rumble"
	matrix, err := mx.ConcatanateLetters(word)
	if err != nil {
		log.Panicln(err)
	}
	go func() {
		offset := mx.OffsetStart
		for offset < offsetLimit {
			trimmedMatrix, err := mx.TrimMatrix(matrix, gui.Rows, gui.Columns, offset)
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
