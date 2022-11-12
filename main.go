// Package main is the entry point for our binary
package main

import (
	"elf-bar-awareness/pkg/gui"
	mx "elf-bar-awareness/pkg/matrix"
	"flag"
	"log"
	"time"
)

const (
	refreshTime       = time.Millisecond * 50
	offsetLimit       = 1000
	defaultMatrixSize = 8
)

func main() {
	// Use the below command just to run in terminal
	// go run . --debug
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var rowCount = flag.Int("rows", defaultMatrixSize, "run in debug mode")
	var colCount = flag.Int("cols", defaultMatrixSize, "run in debug mode")
	flag.Parse()

	if *rowCount == *colCount && *colCount == defaultMatrixSize {
		log.Printf("Using default row col size of %d\n", defaultMatrixSize)
	}
	log.Printf("cols=%d rows=%d", *colCount, *rowCount)

	time.Sleep(time.Second)

	screen := gui.NewTerminalGui(*rowCount, *colCount)
	if !*debugMode {
		var err error
		screen, err = gui.NewledGUI(*rowCount, *colCount)
		if err != nil {
			log.Fatal(err)
		}
	}
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
		trimmedMatrix, err := mx.TrimMatrix(matrix, *rowCount, *colCount, offset)
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
