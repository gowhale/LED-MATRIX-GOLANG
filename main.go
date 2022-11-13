// Package main is the entry point for our binary
package main

import (
	"flag"
	"led-matrix/pkg/config"
	"led-matrix/pkg/gui"
	mx "led-matrix/pkg/matrix"
	"log"
	"time"
)

const (
	refreshTime   = time.Millisecond * 100
	offsetLimit   = 1000
	defaultConfig = "eight-by-eight.json"
)

func main() {
	// Use the below command just to run in terminal
	// go run . --debug
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var configName = flag.String("config", defaultConfig, "run in debug mode")
	flag.Parse()

	if *configName == defaultConfig {
		log.Printf("Using default config file %s\n", defaultConfig)
	}

	cfg, err := config.LoadConfig(defaultConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("cols=%d rows=%d", len(cfg.ColPins), len(cfg.RowPins))

	time.Sleep(time.Second)

	screen := gui.NewTerminalGui(cfg)
	if !*debugMode {
		var err error
		screen, err = gui.NewledGUI(cfg)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := screen.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	run(screen, cfg)
}

func run(screen gui.Screen, cfg config.PinConfig) {
	word := "lets go champ"
	matrix, err := mx.ConcatanateLetters(word)
	if err != nil {
		log.Panicln(err)
	}
	for offset := mx.OffsetStart; offset < offsetLimit; offset++ {
		trimmedMatrix, err := mx.TrimMatrix(matrix, cfg, offset)
		if err != nil {
			log.Fatalln(err)
		}
		if err := screen.AllLEDSOff(); err != nil {
			log.Fatalln(err)
		}
		if err := screen.DisplayMatrix(trimmedMatrix, refreshTime); err != nil {
			log.Panicln(err)
		}
	}
	log.Println("fin.")
}
