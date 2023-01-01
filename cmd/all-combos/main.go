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

	ledmatrix.CreateAnodeMatrix(&common.OSReal{}, cfg.RowPins, cfg.ColPins, *configName)

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

	for x := 0; x < 1000; x++ {
		sleepTime := time.Duration(time.Millisecond * 200)
		for x, _ := range cfg.ColPins {
			for y, _ := range cfg.RowPins {
				log.Println(x, y)
				// time.Sleep(sleepTime)
				startTime := time.Now()
				for time.Since(startTime) < sleepTime {
					screen.CordinatesToLED([2]int{x, y})
				}
			}
		}
	}
	// allPins := append(gui.ColPins, gui.RowsPins...)
	// allPins = append(allPins, gui.R8)
	// // for i, pin1 := range allPins {
	// pin1 := gui.R8
	// i := 0
	// for j, pin2 := range allPins {
	// 	log.Printf("percentage=%.2f pin1=%d pin2=%d \n", (float64(i+j) / float64(len(allPins)) * 100), pin1, pin2)
	// 	// 00
	// 	p1 := rpio.Pin(pin1)
	// 	p2 := rpio.Pin(pin2)
	// 	p1.Low()
	// 	p2.Low()

	// 	// log.Println("00")

	// 	time.Sleep(sleepTime)
	// 	// 01
	// 	p1.Low()
	// 	p2.High()
	// 	// log.Println("01")
	// 	time.Sleep(sleepTime)

	// 	// 10
	// 	p1.High()
	// 	p2.Low()
	// 	// log.Println("10")
	// 	time.Sleep(sleepTime)

	// 	// 11
	// 	p1.High()
	// 	p2.High()
	// 	// log.Println("11")
	// 	time.Sleep(sleepTime)

	// 	// 00
	// 	p1.Low()
	// 	p2.Low()
	// 	// }
	// 	i++
	// }

	defer func() {
		err := screen.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}
