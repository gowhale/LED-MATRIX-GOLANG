// Package main is the entry point for our binary
package main

import (
	"flag"
	"log"
	"time"

	"github.com/gowhale/go-circuit-diagram/pkg/common"
	ledmatrix "github.com/gowhale/go-circuit-diagram/pkg/led-matrix"
	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
	mx "github.com/gowhale/led-matrix-golang/pkg/matrix"
)

const (
	refreshTime = time.Millisecond * 100
)

func main() {
	// Use the below command just to run in terminal
	// go run . --debug
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	var configName = flag.String("config", config.DefaultConfig, "run in debug mode")
	var text = flag.String("text", "abcdefghijklmnopqrstuvwxyz 1234567890", "text to put on the display")
	flag.Parse()

	if *configName == config.DefaultConfig {
		log.Printf("Using default config file %s\n", config.DefaultConfig)
	}

	cfg, err := config.LoadConfig(*configName)
	if err != nil {
		log.Fatal(err)
	}

	if err := ledmatrix.CreateAnodeMatrix(&common.OSReal{}, cfg.RowPins, cfg.ColPins, *configName); err != nil {
		log.Fatalln(err)
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
	run(screen, cfg, *text)
}

func run(screen gui.Screen, cfg config.PinConfig, text string) {
	matrix, err := mx.ConcatanateLetters(text)
	if err != nil {
		log.Panicln(err)
	}
	offsetLimit := len(matrix[0]) + cfg.ColCount()
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
	log.Printf("finished displaying text=%s", text)
}
