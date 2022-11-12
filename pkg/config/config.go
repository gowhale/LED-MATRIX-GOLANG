package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type PinConfig struct {
	RowPins []int `json:"rows"`
	ColPins []int `json:"cols"`
}

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

func (cfg *PinConfig) RowCount() int {
	return len(cfg.RowPins)
}

func (cfg *PinConfig) ColCount() int {
	return len(cfg.ColPins)
}
