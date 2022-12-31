// Package main generates the diagram found in the images folder
package main

import (
	"log"
	"sort"

	"github.com/gowhale/go-circuit-diagram/pkg/canvas"
	cc "github.com/gowhale/go-circuit-diagram/pkg/common"
	"github.com/gowhale/go-circuit-diagram/pkg/components"
	"github.com/gowhale/led-matrix-golang/pkg/gui"
)

const (
	padding = 11
)

func main() {
	board := canvas.NewBoard("all-characters", 55, 88, 10)
	availableChars := []string{}
	for l := range gui.LetterMap {
		availableChars = append(availableChars, l)
	}
	sort.Strings(availableChars)

	col := 0
	row := 0
	for _, l := range availableChars {
		if col > 0 && col%5 == 0 {
			row++
			col = 0
		}
		coord := cc.NewCord(5+(col*padding), 5+(row*padding))
		letterLabel := components.NewLabel(coord, l)
		board.AddElement(&letterLabel)
		col++
	}
	err := board.Draw(&cc.OSReal{})
	if err != nil {
		log.Fatalln(err)
	}
}
