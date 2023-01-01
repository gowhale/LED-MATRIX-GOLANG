// Package config is responsible for loading GPIO configurations
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	// DefaultConfig is the filename of defualt config used if no override provided
	DefaultConfig = "eight-by-eight.json"
)

// PinConfig is used to configure GPIO pins to rows / columns
// This is based on the assumption that the Pi is connection to a LED matrix
// Example matrix below
// https://www.jameco.com/Jameco/workshop/learning-center/electronic-fundamentals-working-with-led-dot-matrix-displays-fig3.jpg
type PinConfig struct {
	RowPins []int `json:"rows"`
	ColPins []int `json:"cols"`
}

// LoadConfig will load a json file and return PinConfig
func LoadConfig(filename string) (PinConfig, error) {
	return loadConfigImpl(&readJSON{}, filename)
}

func loadConfigImpl(jReader readerJSON, filename string) (PinConfig, error) {
	file, err := jReader.ReadFile(fmt.Sprintf("config-files/%s", filename))
	if err != nil {
		return PinConfig{}, err
	}

	config := PinConfig{}

	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return PinConfig{}, err
	}

	log.Printf("row pins = %+v", config.RowPins)
	log.Printf("col pins = %+v", config.ColPins)

	return config, nil
}

// RowCount returns the amount of rows specified in a Config file
func (cfg *PinConfig) RowCount() int {
	return len(cfg.RowPins)
}

// ColCount returns the amount of cols specified in a Config file
func (cfg *PinConfig) ColCount() int {
	return len(cfg.ColPins)
}

type readJSON struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name readerJSON -inpkg --filename read_json_mock.go
type readerJSON interface {
	ReadFile(filename string) ([]byte, error)
}

func (*readJSON) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
