// Package config is responsible for loading GPIO configurations
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
func LoadConfig(fileName string) (PinConfig, error) {
	file, err := ioutil.ReadFile(fmt.Sprintf("config-files/%s", fileName))
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
