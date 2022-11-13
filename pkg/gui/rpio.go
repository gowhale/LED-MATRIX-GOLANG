package gui

import "github.com/stianeikeland/go-rpio"

type rpioProc struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name rpioProcessor -inpkg --filename rpio_processor_mock.go
type rpioProcessor interface {
	Open() (err error)
	Close() (err error)
	Pin(p int) pinProcessor
}

func (*rpioProc) Open() (err error) {
	return rpio.Open()
}

func (*rpioProc) Close() (err error) {
	return rpio.Close()
}

func (*rpioProc) Pin(p int) pinProcessor {
	return rpio.Pin(p)
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name pinProcessor -inpkg --filename pin_processor_mock.go
type pinProcessor interface {
	Output()
	Low()
	High()
}

type guiLED struct {
	rowCount, colCount int
	rowPins, colPins   []int
	rpioController     rpioProcessor
}
