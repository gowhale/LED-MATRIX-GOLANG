// Package main contains a main prog which
// goes through each led in a matrix and lights them
// one by one.
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/gowhale/go-circuit-diagram/pkg/common"
	ledmatrix "github.com/gowhale/go-circuit-diagram/pkg/led-matrix"
	"github.com/gowhale/led-matrix-golang/pkg/config"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

func main() {
	var configName = flag.String("config", config.DefaultConfig, "run in debug mode")
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	flag.Parse()
	cfg, err := config.LoadConfig(fmt.Sprintf("%s.json", *configName))
	if err != nil {
		log.Fatal(err)
	}

	if err := ledmatrix.CreateAnodeMatrix(&common.OSReal{}, cfg.RowPins, cfg.ColPins, *configName); err != nil {
		log.Fatalln(err)
	}

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

	if err := lightCoordinates(screen, cfg); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err := screen.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}

func lightOneLED(screen gui.Screen, x, y int) error {
	sleepTime := time.Duration(time.Millisecond * 200)
	log.Println(x, y)
	startTime := time.Now()
	for time.Since(startTime) < sleepTime {
		if err := screen.CordinatesToLED([2]int{x, y}); err != nil {
			return err
		}
	}
	return nil
}

func lightCoordinates(screen gui.Screen, cfg config.PinConfig) error {
	for x := 0; x < 1000; x++ {
		for x := range cfg.ColPins {
			for y := range cfg.RowPins {
				if err := lightOneLED(screen, x, y); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
